import pool from './db'
import { Drug, Category, User, Order, CartItem } from '@/types'

export class DrugService {
  static async getAllDrugs(): Promise<Drug[]> {
    const query = `
      SELECT d.*, c.name as category_name 
      FROM drugs d 
      LEFT JOIN categories c ON d.category_id = c.id 
      ORDER BY d.name
    `
    const result = await pool.query(query)
    return result.rows.map(this.mapRowToDrug)
  }

  static async getDrugById(id: number): Promise<Drug | null> {
    const query = `
      SELECT d.*, c.name as category_name 
      FROM drugs d 
      LEFT JOIN categories c ON d.category_id = c.id 
      WHERE d.id = $1
    `
    const result = await pool.query(query, [id])
    return result.rows.length > 0 ? this.mapRowToDrug(result.rows[0]) : null
  }

  static async searchDrugs(searchTerm: string): Promise<Drug[]> {
    const query = `
      SELECT d.*, c.name as category_name 
      FROM drugs d 
      LEFT JOIN categories c ON d.category_id = c.id 
      WHERE d.name ILIKE $1 OR d.description ILIKE $1 OR d.composition ILIKE $1
      ORDER BY d.name
    `
    const result = await pool.query(query, [`%${searchTerm}%`])
    return result.rows.map(this.mapRowToDrug)
  }

  static async getDrugsByCategory(categoryId: number): Promise<Drug[]> {
    const query = `
      SELECT d.*, c.name as category_name 
      FROM drugs d 
      LEFT JOIN categories c ON d.category_id = c.id 
      WHERE d.category_id = $1 
      ORDER BY d.name
    `
    const result = await pool.query(query, [categoryId])
    return result.rows.map(this.mapRowToDrug)
  }

  private static mapRowToDrug(row: any): Drug {
    return {
      id: row.id,
      name: row.name,
      description: row.description,
      composition: row.composition,
      price: parseFloat(row.price),
      stock: row.stock,
      category: row.category_name || '',
      manufacturer: row.manufacturer,
      dosage: row.dosage,
      sideEffects: row.side_effects || [],
      contraindications: row.contraindications || [],
      imageUrl: row.image_url,
      requiresPrescription: row.requires_prescription
    }
  }
}

export class CategoryService {
  static async getAllCategories(): Promise<Category[]> {
    const query = `
      SELECT c.*, COUNT(d.id) as drug_count 
      FROM categories c 
      LEFT JOIN drugs d ON c.id = d.category_id 
      GROUP BY c.id 
      ORDER BY c.name
    `
    const result = await pool.query(query)
    return result.rows.map(row => ({
      id: row.id,
      name: row.name,
      slug: row.slug,
      description: row.description,
      icon: row.icon,
      count: parseInt(row.drug_count)
    }))
  }

  static async getCategoryBySlug(slug: string): Promise<Category | null> {
    const query = `
      SELECT c.*, COUNT(d.id) as drug_count 
      FROM categories c 
      LEFT JOIN drugs d ON c.id = d.category_id 
      WHERE c.slug = $1 
      GROUP BY c.id
    `
    const result = await pool.query(query, [slug])
    return result.rows.length > 0 ? {
      id: result.rows[0].id,
      name: result.rows[0].name,
      slug: result.rows[0].slug,
      description: result.rows[0].description,
      icon: result.rows[0].icon,
      count: parseInt(result.rows[0].drug_count)
    } : null
  }
}

export class UserService {
  static async createUser(userData: {
    email: string
    passwordHash: string
    name: string
    phone?: string
  }): Promise<User> {
    const query = `
      INSERT INTO users (email, password_hash, name, phone) 
      VALUES ($1, $2, $3, $4) 
      RETURNING id, email, name, phone, created_at
    `
    const result = await pool.query(query, [
      userData.email,
      userData.passwordHash,
      userData.name,
      userData.phone
    ])
    return {
      id: result.rows[0].id,
      email: result.rows[0].email,
      name: result.rows[0].name,
      phone: result.rows[0].phone
    }
  }

  static async getUserByEmail(email: string): Promise<(User & { passwordHash: string }) | null> {
    const query = 'SELECT * FROM users WHERE email = $1'
    const result = await pool.query(query, [email])
    return result.rows.length > 0 ? {
      id: result.rows[0].id,
      email: result.rows[0].email,
      name: result.rows[0].name,
      phone: result.rows[0].phone,
      address: result.rows[0].address,
      passwordHash: result.rows[0].password_hash
    } : null
  }

  static async getUserById(id: string): Promise<User | null> {
    const query = 'SELECT id, email, name, phone, address FROM users WHERE id = $1'
    const result = await pool.query(query, [id])
    return result.rows.length > 0 ? {
      id: result.rows[0].id,
      email: result.rows[0].email,
      name: result.rows[0].name,
      phone: result.rows[0].phone,
      address: result.rows[0].address
    } : null
  }
}

export class CartService {
  static async getCartItems(userId: string): Promise<CartItem[]> {
    const query = `
      SELECT ci.quantity, d.*, c.name as category_name
      FROM cart_items ci
      JOIN drugs d ON ci.drug_id = d.id
      LEFT JOIN categories c ON d.category_id = c.id
      WHERE ci.user_id = $1
    `
    const result = await pool.query(query, [userId])
    return result.rows.map(row => ({
      drug: DrugService['mapRowToDrug'](row),
      quantity: row.quantity
    }))
  }

  static async addToCart(userId: string, drugId: number, quantity: number): Promise<void> {
    const query = `
      INSERT INTO cart_items (user_id, drug_id, quantity) 
      VALUES ($1, $2, $3) 
      ON CONFLICT (user_id, drug_id) 
      DO UPDATE SET quantity = cart_items.quantity + $3, updated_at = CURRENT_TIMESTAMP
    `
    await pool.query(query, [userId, drugId, quantity])
  }

  static async updateCartItem(userId: string, drugId: number, quantity: number): Promise<void> {
    if (quantity === 0) {
      await pool.query('DELETE FROM cart_items WHERE user_id = $1 AND drug_id = $2', [userId, drugId])
    } else {
      const query = `
        UPDATE cart_items 
        SET quantity = $3, updated_at = CURRENT_TIMESTAMP 
        WHERE user_id = $1 AND drug_id = $2
      `
      await pool.query(query, [userId, drugId, quantity])
    }
  }

  static async clearCart(userId: string): Promise<void> {
    await pool.query('DELETE FROM cart_items WHERE user_id = $1', [userId])
  }
}

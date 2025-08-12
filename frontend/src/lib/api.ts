const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'

export interface Drug {
  id: number
  name: string
  description: string
  composition: string
  price: number
  stock: number
  category_id: number
  category_name?: string
  manufacturer: string
  dosage: string
  side_effects: string[]
  contraindications: string[]
  image_url: string
  requires_prescription: boolean
  created_at: string
  updated_at: string
}

export interface Category {
  id: number
  name: string
  slug: string
  description: string
  icon: string
  count?: number
  created_at: string
  updated_at: string
}

export interface APIResponse<T> {
  success: boolean
  message?: string
  data?: T
  error?: string
}

// API functions
export const api = {
  // Drugs
  async getAllDrugs(): Promise<Drug[]> {
    const response = await fetch(`${API_BASE_URL}/drugs`)
    const data: APIResponse<Drug[]> = await response.json()
    return data.data || []
  },

  async getDrugById(id: number): Promise<Drug | null> {
    const response = await fetch(`${API_BASE_URL}/drugs/${id}`)
    if (!response.ok) return null
    const data: APIResponse<Drug> = await response.json()
    return data.data || null
  },

  async searchDrugs(query: string): Promise<Drug[]> {
    const response = await fetch(`${API_BASE_URL}/drugs/search?q=${encodeURIComponent(query)}`)
    const data: APIResponse<Drug[]> = await response.json()
    return data.data || []
  },

  // Categories
  async getAllCategories(): Promise<Category[]> {
    const response = await fetch(`${API_BASE_URL}/categories`)
    const data: APIResponse<Category[]> = await response.json()
    return data.data || []
  },

  async getCategoryBySlug(slug: string): Promise<Category | null> {
    const response = await fetch(`${API_BASE_URL}/categories/${slug}`)
    if (!response.ok) return null
    const data: APIResponse<Category> = await response.json()
    return data.data || null
  },

  async getDrugsByCategory(categoryId: number): Promise<Drug[]> {
    const response = await fetch(`${API_BASE_URL}/categories/${categoryId}/drugs`)
    const data: APIResponse<Drug[]> = await response.json()
    return data.data || []
  },

  // Health check
  async healthCheck(): Promise<{ status: string; service: string }> {
    const response = await fetch(`${API_BASE_URL}/health`)
    return response.json()
  }
}

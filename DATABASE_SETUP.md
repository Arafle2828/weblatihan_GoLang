# PharmaCare Database Setup

## PostgreSQL Installation & Setup

### 1. Install PostgreSQL
Download dan install PostgreSQL dari: https://www.postgresql.org/download/

### 2. Create Database
```sql
-- Login ke PostgreSQL sebagai superuser
psql -U postgres

-- Buat database
CREATE DATABASE pharmacare;

-- Buat user khusus (opsional)
CREATE USER pharmacare_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE pharmacare TO pharmacare_user;
```

### 3. Run Database Schema
```bash
# Dari direktori project
psql -U postgres -d pharmacare -f database/schema.sql
```

### 4. Environment Setup
Copy file `.env.local` dan sesuaikan dengan konfigurasi database Anda:

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=pharmacare
DB_USER=postgres
DB_PASSWORD=your_actual_password
```

### 5. Install Dependencies
Sudah diinstall:
- `pg` - PostgreSQL client
- `@types/pg` - TypeScript types
- `bcryptjs` - Password hashing
- `jsonwebtoken` - JWT authentication

### 6. Test Connection
Coba jalankan aplikasi dan akses:
- http://localhost:3000/drugs - Halaman katalog obat
- http://localhost:3000/api/drugs - API endpoint drugs
- http://localhost:3000/api/categories - API endpoint categories

## Database Structure

### Tables Created:
- `categories` - Kategori obat
- `drugs` - Data obat
- `users` - Data pengguna
- `orders` - Data pesanan
- `order_items` - Item dalam pesanan
- `cart_items` - Keranjang belanja
- `reviews` - Review obat

### Sample Data:
- 6 kategori obat
- 8 obat sample
- 1 admin user (email: admin@pharmacare.com, password: admin123)

## Features Ready:
✅ Database schema dan sample data
✅ Service layer untuk CRUD operations
✅ API routes untuk drugs dan categories
✅ Updated drugs page dengan real data
✅ TypeScript types
✅ Environment configuration

## Next Steps:
1. Setup authentication
2. Shopping cart functionality
3. Order management
4. User profile
5. Admin panel

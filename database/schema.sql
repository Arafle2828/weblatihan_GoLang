-- PharmaCare Database Schema

-- Create database (run this separately)
-- CREATE DATABASE pharmacare;

-- Categories table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Drugs table
CREATE TABLE drugs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    composition TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER DEFAULT 0,
    category_id INTEGER REFERENCES categories(id),
    manufacturer VARCHAR(200),
    dosage VARCHAR(100),
    side_effects TEXT[],
    contraindications TEXT[],
    image_url VARCHAR(500),
    requires_prescription BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(200) NOT NULL,
    phone VARCHAR(20),
    address JSONB,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    shipping_address JSONB,
    payment_method VARCHAR(50),
    payment_status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Order items table
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    drug_id INTEGER REFERENCES drugs(id),
    quantity INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Cart table
CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    drug_id INTEGER REFERENCES drugs(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, drug_id)
);

-- Reviews table
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    drug_id INTEGER REFERENCES drugs(id),
    user_id INTEGER REFERENCES users(id),
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for better performance
CREATE INDEX idx_drugs_category ON drugs(category_id);
CREATE INDEX idx_drugs_name ON drugs(name);
CREATE INDEX idx_orders_user ON orders(user_id);
CREATE INDEX idx_order_items_order ON order_items(order_id);
CREATE INDEX idx_cart_user ON cart_items(user_id);
CREATE INDEX idx_reviews_drug ON reviews(drug_id);

-- Insert sample categories
INSERT INTO categories (name, slug, description, icon) VALUES
('Demam & Flu', 'demam-flu', 'Obat untuk mengatasi demam dan flu', '🤒'),
('Vitamin & Suplemen', 'vitamin-suplemen', 'Vitamin dan suplemen untuk kesehatan', '💊'),
('Obat Pencernaan', 'pencernaan', 'Obat untuk gangguan pencernaan', '🫀'),
('Perawatan Kulit', 'perawatan-kulit', 'Produk perawatan kulit dan antiseptik', '🧴'),
('Obat Jantung', 'jantung', 'Obat untuk kesehatan jantung', '❤️'),
('Diabetes', 'diabetes', 'Obat dan alat untuk diabetes', '🩺');

-- Insert sample drugs
INSERT INTO drugs (name, description, composition, price, stock, category_id, manufacturer, dosage, side_effects, contraindications, requires_prescription) VALUES
('Paracetamol 500mg', 'Obat penurun panas dan pereda nyeri', 'Paracetamol 500mg', 15000, 50, 1, 'Kimia Farma', '3x sehari 1 tablet', ARRAY['Mual', 'Pusing ringan'], ARRAY['Alergi paracetamol', 'Gangguan hati'], FALSE),

('Vitamin C 1000mg', 'Suplemen vitamin C untuk daya tahan tubuh', 'Ascorbic Acid 1000mg', 25000, 30, 2, 'Kalbe Farma', '1x sehari 1 tablet', ARRAY['Gangguan pencernaan ringan'], ARRAY['Batu ginjal'], FALSE),

('Antasida', 'Obat maag dan gangguan pencernaan', 'Aluminum Hydroxide, Magnesium Hydroxide', 12000, 25, 3, 'Dexa Medica', '3x sehari 1 tablet setelah makan', ARRAY['Konstipasi', 'Diare'], ARRAY['Gangguan ginjal'], FALSE),

('Betadine 60ml', 'Antiseptik untuk luka dan infeksi', 'Povidone Iodine 10%', 18000, 40, 4, 'Mundipharma', 'Oleskan pada area yang terluka', ARRAY['Iritasi kulit ringan'], ARRAY['Alergi iodine', 'Hipertiroid'], FALSE),

('Amoxicillin 500mg', 'Antibiotik untuk infeksi bakteri', 'Amoxicillin Trihydrate 500mg', 35000, 15, 1, 'Sanbe Farma', '3x sehari 1 kapsul', ARRAY['Mual', 'Diare', 'Ruam kulit'], ARRAY['Alergi penisilin'], TRUE),

('Omeprazole 20mg', 'Obat asam lambung dan GERD', 'Omeprazole 20mg', 45000, 20, 3, 'Dexa Medica', '1x sehari 1 kapsul sebelum makan', ARRAY['Sakit kepala', 'Mual'], ARRAY['Alergi omeprazole'], TRUE),

('Multivitamin', 'Suplemen vitamin dan mineral lengkap', 'Vitamin A, B, C, D, E + Mineral', 65000, 40, 2, 'Blackmores', '1x sehari 1 tablet setelah makan', ARRAY['Gangguan pencernaan ringan'], ARRAY['Hipervitaminosis'], FALSE),

('Ibuprofen 400mg', 'Anti inflamasi dan pereda nyeri', 'Ibuprofen 400mg', 22000, 35, 1, 'Tempo Scan', '3x sehari 1 tablet', ARRAY['Mual', 'Nyeri perut'], ARRAY['Tukak lambung', 'Gangguan ginjal'], FALSE);

-- Create a sample admin user (password: admin123)
INSERT INTO users (email, password_hash, name, phone, is_verified) VALUES
('admin@pharmacare.com', '$2b$10$8K1p/a0dqailSRMGiSHkTOZOaKXnhOONWkz8MHkV2kY9mNZ.v9XrK', 'Admin PharmaCare', '08123456789', TRUE);

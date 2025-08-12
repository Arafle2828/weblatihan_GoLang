export interface Drug {
  id: number
  name: string
  description: string
  composition: string
  price: number
  stock: number
  category: string
  manufacturer: string
  dosage: string
  sideEffects: string[]
  contraindications: string[]
  imageUrl?: string
  requiresPrescription: boolean
}

export interface Category {
  id: number
  name: string
  slug: string
  description: string
  icon: string
  count: number
}

export interface CartItem {
  drug: Drug
  quantity: number
}

export interface Order {
  id: string
  items: CartItem[]
  total: number
  status: 'pending' | 'processing' | 'shipped' | 'delivered' | 'cancelled'
  createdAt: Date
  shippingAddress: Address
}

export interface Address {
  street: string
  city: string
  province: string
  postalCode: string
  phone: string
}

export interface User {
  id: string
  name: string
  email: string
  phone: string
  address?: Address
}

'use client'

import Link from 'next/link'
import { useState } from 'react'
import { Menu, X, ShoppingCart, User, Search } from 'lucide-react'

export default function Header() {
  const [isMenuOpen, setIsMenuOpen] = useState(false)

  return (
    <header className="bg-white shadow-md sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link href="/" className="flex items-center space-x-2">
            <div className="bg-blue-600 text-white w-8 h-8 rounded-full flex items-center justify-center font-bold">
              P
            </div>
            <span className="font-bold text-xl text-gray-900">PharmaCare</span>
          </Link>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <Link href="/drugs" className="text-gray-700 hover:text-blue-600 transition-colors">
              Katalog Obat
            </Link>
            <Link href="/consultation" className="text-gray-700 hover:text-blue-600 transition-colors">
              Konsultasi
            </Link>
            <Link href="/about" className="text-gray-700 hover:text-blue-600 transition-colors">
              Tentang Kami
            </Link>
            <Link href="/contact" className="text-gray-700 hover:text-blue-600 transition-colors">
              Kontak
            </Link>
          </nav>

          {/* Actions */}
          <div className="flex items-center space-x-4">
            <button className="text-gray-700 hover:text-blue-600 transition-colors">
              <Search className="h-5 w-5" />
            </button>
            <Link href="/cart" className="text-gray-700 hover:text-blue-600 transition-colors relative">
              <ShoppingCart className="h-5 w-5" />
              <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">
                3
              </span>
            </Link>
            <Link href="/profile" className="text-gray-700 hover:text-blue-600 transition-colors">
              <User className="h-5 w-5" />
            </Link>

            {/* Mobile menu button */}
            <button
              className="md:hidden text-gray-700"
              onClick={() => setIsMenuOpen(!isMenuOpen)}
            >
              {isMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
            </button>
          </div>
        </div>

        {/* Mobile Navigation */}
        {isMenuOpen && (
          <div className="md:hidden py-4 border-t border-gray-200">
            <nav className="flex flex-col space-y-2">
              <Link href="/drugs" className="px-2 py-2 text-gray-700 hover:text-blue-600 transition-colors">
                Katalog Obat
              </Link>
              <Link href="/consultation" className="px-2 py-2 text-gray-700 hover:text-blue-600 transition-colors">
                Konsultasi
              </Link>
              <Link href="/about" className="px-2 py-2 text-gray-700 hover:text-blue-600 transition-colors">
                Tentang Kami
              </Link>
              <Link href="/contact" className="px-2 py-2 text-gray-700 hover:text-blue-600 transition-colors">
                Kontak
              </Link>
            </nav>
          </div>
        )}
      </div>
    </header>
  )
}

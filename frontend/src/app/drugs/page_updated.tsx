'use client'

import { useState, useEffect } from 'react'
import { Search, Filter, ShoppingCart } from 'lucide-react'
import Link from 'next/link'
import { api, Drug, Category } from '@/lib/api'

export default function DrugsPage() {
  const [drugs, setDrugs] = useState<Drug[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [filteredDrugs, setFilteredDrugs] = useState<Drug[]>([])
  const [searchQuery, setSearchQuery] = useState('')
  const [selectedCategory, setSelectedCategory] = useState<string>('all')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [drugsData, categoriesData] = await Promise.all([
          api.getAllDrugs(),
          api.getAllCategories()
        ])
        setDrugs(drugsData)
        setCategories(categoriesData)
        setFilteredDrugs(drugsData)
      } catch (err) {
        setError('Failed to load data. Please check if the backend server is running.')
        console.error('Error fetching data:', err)
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [])

  useEffect(() => {
    let filtered = drugs

    // Filter by search query
    if (searchQuery) {
      filtered = filtered.filter(drug =>
        drug.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        drug.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
        drug.composition.toLowerCase().includes(searchQuery.toLowerCase())
      )
    }

    // Filter by category
    if (selectedCategory !== 'all') {
      const category = categories.find(cat => cat.slug === selectedCategory)
      if (category) {
        filtered = filtered.filter(drug => drug.category_id === category.id)
      }
    }

    setFilteredDrugs(filtered)
  }, [searchQuery, selectedCategory, drugs, categories])

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading drugs...</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <p className="text-red-600 mb-4">{error}</p>
          <button 
            onClick={() => window.location.reload()} 
            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
          >
            Try Again
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 py-8">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-4">Katalog Obat</h1>
          <p className="text-gray-600">Temukan obat yang Anda butuhkan dari koleksi lengkap kami</p>
        </div>

        {/* Search and Filter */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-8">
          <div className="flex flex-col md:flex-row gap-4">
            <div className="flex-1 relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5" />
              <input
                type="text"
                placeholder="Cari obat berdasarkan nama atau kandungan..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            <select
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
              className="px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="all">Semua Kategori</option>
              {categories.map((category) => (
                <option key={category.id} value={category.slug}>
                  {category.name}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Categories */}
        <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-8">
          {categories.map((category) => (
            <button
              key={category.id}
              onClick={() => setSelectedCategory(category.slug)}
              className={`bg-white rounded-lg p-4 text-center hover:shadow-md transition-shadow border ${
                selectedCategory === category.slug ? 'border-blue-500 bg-blue-50' : ''
              }`}
            >
              <div className="text-2xl mb-2">{category.icon}</div>
              <h3 className="font-medium text-sm">{category.name}</h3>
            </button>
          ))}
        </div>

        {/* Results Count */}
        <div className="mb-6">
          <p className="text-gray-600">
            Menampilkan {filteredDrugs.length} dari {drugs.length} obat
            {searchQuery && ` untuk "${searchQuery}"`}
            {selectedCategory !== 'all' && ` dalam kategori "${categories.find(c => c.slug === selectedCategory)?.name}"`}
          </p>
        </div>

        {/* Products Grid */}
        {filteredDrugs.length === 0 ? (
          <div className="text-center py-12">
            <div className="text-6xl mb-4">🔍</div>
            <h3 className="text-xl font-semibold text-gray-900 mb-2">Tidak ada obat ditemukan</h3>
            <p className="text-gray-600 mb-4">
              Coba ubah kata kunci pencarian atau pilih kategori yang berbeda
            </p>
            <button
              onClick={() => {
                setSearchQuery('')
                setSelectedCategory('all')
              }}
              className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
            >
              Reset Filter
            </button>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {filteredDrugs.map((drug) => (
              <div key={drug.id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
                <div className="aspect-square bg-gray-100 flex items-center justify-center">
                  <div className="text-4xl">💊</div>
                </div>
                
                <div className="p-4">
                  <h3 className="font-semibold text-lg text-gray-900 mb-2">{drug.name}</h3>
                  <p className="text-sm text-gray-600 mb-2">{drug.description}</p>
                  <p className="text-xs text-gray-500 mb-3">{drug.composition}</p>
                  
                  <div className="flex items-center justify-between mb-3">
                    <span className="text-lg font-bold text-blue-600">Rp {drug.price.toLocaleString('id-ID')}</span>
                    <span className={`px-2 py-1 rounded-full text-xs ${
                      drug.stock > 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                    }`}>
                      {drug.stock > 0 ? `${drug.stock} tersedia` : 'Habis'}
                    </span>
                  </div>

                  {drug.requires_prescription && (
                    <div className="mb-3">
                      <span className="inline-block bg-red-100 text-red-800 text-xs px-2 py-1 rounded-full">
                        Resep Dokter
                      </span>
                    </div>
                  )}
                  
                  <div className="flex space-x-2">
                    <Link
                      href={`/drugs/${drug.id}`}
                      className="flex-1 bg-blue-600 text-white text-center py-2 rounded-lg hover:bg-blue-700 transition-colors"
                    >
                      Detail
                    </Link>
                    <button 
                      className="flex items-center justify-center w-10 h-10 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
                      disabled={drug.stock === 0}
                    >
                      <ShoppingCart className="h-4 w-4" />
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

import { Search, Filter, ShoppingCart } from 'lucide-react'
import Link from 'next/link'
import { DrugService, CategoryService } from '@/lib/services'

export default async function DrugsPage() {
  const drugs = await DrugService.getAllDrugs()
  const categories = await CategoryService.getAllCategories()

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
                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            <button className="flex items-center space-x-2 px-4 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
              <Filter className="h-5 w-5" />
              <span>Filter</span>
            </button>
          </div>
        </div>

        {/* Categories */}
        <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-8">
          {categories.map((category) => (
            <Link
              key={category.id}
              href={`/drugs/category/${category.slug}`}
              className="bg-white rounded-lg p-4 text-center hover:shadow-md transition-shadow border"
            >
              <div className="text-2xl mb-2">{category.icon}</div>
              <h3 className="font-medium text-sm text-gray-900">{category.name}</h3>
              <p className="text-xs text-gray-500">{category.count} produk</p>
            </Link>
          ))}
        </div>

        {/* Products Grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {drugs.map((drug) => (
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

                {drug.requiresPrescription && (
                  <div className="mb-3">
                    <span className="inline-block bg-red-100 text-red-800 text-xs px-2 py-1 rounded-full">
                      Resep Dokter
                    </span>
                  </div>
                )}
                
                <div className="flex space-x-2">
                  <Link
                    href={`/drugs/${drug.id}`}
                    className="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg text-center text-sm hover:bg-blue-700 transition-colors"
                  >
                    Detail
                  </Link>
                  <button
                    disabled={drug.stock === 0}
                    className="flex items-center justify-center bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-colors disabled:bg-gray-400 disabled:cursor-not-allowed"
                  >
                    <ShoppingCart className="h-4 w-4" />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Load More */}
        <div className="text-center mt-8">
          <button className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
            Muat Lebih Banyak
          </button>
        </div>
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
                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>
            <button className="flex items-center space-x-2 px-4 py-3 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
              <Filter className="h-5 w-5" />
              <span>Filter</span>
            </button>
          </div>
        </div>

        {/* Categories */}
        <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4 mb-8">
          {categories.map((category, index) => (
            <Link
              key={index}
              href={`/drugs/category/${category.slug}`}
              className="bg-white rounded-lg p-4 text-center hover:shadow-md transition-shadow border"
            >
              <div className="text-2xl mb-2">{category.icon}</div>
              <h3 className="font-medium text-sm text-gray-900">{category.name}</h3>
            </Link>
          ))}
        </div>

        {/* Products Grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {drugs.map((drug, index) => (
            <div key={index} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
              <div className="aspect-square bg-gray-100 flex items-center justify-center">
                <div className="text-4xl">{drug.icon}</div>
              </div>
              
              <div className="p-4">
                <h3 className="font-semibold text-lg text-gray-900 mb-2">{drug.name}</h3>
                <p className="text-sm text-gray-600 mb-2">{drug.description}</p>
                <p className="text-xs text-gray-500 mb-3">{drug.composition}</p>
                
                <div className="flex items-center justify-between mb-3">
                  <span className="text-lg font-bold text-blue-600">Rp {drug.price.toLocaleString()}</span>
                  <span className={`px-2 py-1 rounded-full text-xs ${
                    drug.stock > 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  }`}>
                    {drug.stock > 0 ? 'Tersedia' : 'Habis'}
                  </span>
                </div>
                
                <div className="flex space-x-2">
                  <Link
                    href={`/drugs/${drug.id}`}
                    className="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg text-center text-sm hover:bg-blue-700 transition-colors"
                  >
                    Detail
                  </Link>
                  <button
                    disabled={drug.stock === 0}
                    className="flex items-center justify-center bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-colors disabled:bg-gray-400 disabled:cursor-not-allowed"
                  >
                    <ShoppingCart className="h-4 w-4" />
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Load More */}
        <div className="text-center mt-8">
          <button className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
            Muat Lebih Banyak
          </button>
        </div>
      </div>
    </div>
  )
}

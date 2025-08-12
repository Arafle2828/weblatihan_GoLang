import Link from 'next/link'
import { Search, Package, Clock, Shield } from 'lucide-react'

export default function HomePage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-green-50">
      {/* Hero Section */}
      <section className="relative py-20 px-4">
        <div className="max-w-6xl mx-auto text-center">
          <h1 className="text-5xl font-bold text-gray-900 mb-6">
            PharmaCare
            <span className="text-blue-600"> Online</span>
          </h1>
          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
            Temukan informasi obat lengkap dan pesan obat dengan mudah. 
            Layanan farmasi online terpercaya untuk kesehatan Anda.
          </p>
          
          {/* Search Bar */}
          <div className="max-w-md mx-auto mb-8">
            <div className="relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5" />
              <input
                type="text"
                placeholder="Cari obat, vitamin, atau suplemen..."
                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
              <button className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-blue-600 text-white px-4 py-1.5 rounded-md hover:bg-blue-700 transition-colors">
                Cari
              </button>
            </div>
          </div>

          <div className="flex flex-wrap justify-center gap-4">
            <Link href="/drugs" className="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
              Lihat Katalog Obat
            </Link>
            <Link href="/consultation" className="bg-white text-blue-600 border-2 border-blue-600 px-6 py-3 rounded-lg hover:bg-blue-50 transition-colors">
              Konsultasi Online
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 px-4 bg-white">
        <div className="max-w-6xl mx-auto">
          <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
            Mengapa Memilih PharmaCare?
          </h2>
          
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div className="text-center">
              <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <Package className="h-8 w-8 text-blue-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Stok Lengkap</h3>
              <p className="text-gray-600">Ribuan jenis obat dan suplemen tersedia</p>
            </div>

            <div className="text-center">
              <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <Clock className="h-8 w-8 text-green-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Pengiriman Cepat</h3>
              <p className="text-gray-600">Dikirim dalam 24 jam ke seluruh Indonesia</p>
            </div>

            <div className="text-center">
              <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <Shield className="h-8 w-8 text-purple-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Terjamin Asli</h3>
              <p className="text-gray-600">Semua produk berlisensi resmi BPOM</p>
            </div>

            <div className="text-center">
              <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
                <Search className="h-8 w-8 text-orange-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">Mudah Dicari</h3>
              <p className="text-gray-600">Pencarian obat berdasarkan gejala atau nama</p>
            </div>
          </div>
        </div>
      </section>

      {/* Popular Categories */}
      <section className="py-16 px-4 bg-gray-50">
        <div className="max-w-6xl mx-auto">
          <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
            Kategori Populer
          </h2>
          
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {categories.map((category, index) => (
              <Link
                key={index}
                href={`/drugs/category/${category.slug}`}
                className="bg-white rounded-lg p-6 shadow-md hover:shadow-lg transition-shadow border-l-4 border-blue-500"
              >
                <h3 className="text-xl font-semibold text-gray-900 mb-2">
                  {category.name}
                </h3>
                <p className="text-gray-600 mb-3">{category.description}</p>
                <span className="text-blue-600 font-medium">
                  {category.count} produk →
                </span>
              </Link>
            ))}
          </div>
        </div>
      </section>
    </div>
  )
}

const categories = [
  {
    name: "Obat Demam & Flu",
    description: "Paracetamol, Ibuprofen, Obat batuk",
    count: 45,
    slug: "demam-flu"
  },
  {
    name: "Vitamin & Suplemen",
    description: "Vitamin C, D3, Multivitamin",
    count: 120,
    slug: "vitamin-suplemen"
  },
  {
    name: "Obat Pencernaan",
    description: "Antasida, Probiotik, Anti diare",
    count: 38,
    slug: "pencernaan"
  },
  {
    name: "Perawatan Kulit",
    description: "Antiseptik, Krim, Salep",
    count: 67,
    slug: "perawatan-kulit"
  },
  {
    name: "Obat Jantung",
    description: "Obat hipertensi, Kolesterol",
    count: 29,
    slug: "jantung"
  },
  {
    name: "Diabetes",
    description: "Insulin, Glukometer, Strip",
    count: 34,
    slug: "diabetes"
  }
]

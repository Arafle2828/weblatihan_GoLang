import { NextRequest, NextResponse } from 'next/server'
import { DrugService } from '@/lib/services'

export async function GET(request: NextRequest) {
  try {
    const { searchParams } = new URL(request.url)
    const search = searchParams.get('search')
    const category = searchParams.get('category')

    let drugs

    if (search) {
      drugs = await DrugService.searchDrugs(search)
    } else if (category) {
      drugs = await DrugService.getDrugsByCategory(parseInt(category))
    } else {
      drugs = await DrugService.getAllDrugs()
    }

    return NextResponse.json({ drugs })
  } catch (error) {
    console.error('Error fetching drugs:', error)
    return NextResponse.json(
      { error: 'Failed to fetch drugs' },
      { status: 500 }
    )
  }
}

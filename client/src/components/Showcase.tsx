import { Smile } from 'lucide-react'
import ProductCard from './ProductCard'

export type Product = {
  title: string
  price: number
  permalink: string
  thumbnail: string
}

type ShowcaseProps = {
  products: Product[]
}

const Showcase = ({ products }: ShowcaseProps) => {
  return (
    <section className="flex flex-col gap-4 justify-center mt-6 items-center w-full h-8/12 overflow-y-auto">
      {products.length === 0 ? (
        <h1 className="flex gap-2 items-center mt-8 text-2xl font-bold text-gray-700 sm:text-4xl">
          Faça sua pesquisa
          <span>
            <Smile size={32} />
          </span>
        </h1>
      ) : (
        products.map((product) => <ProductCard key={product.title} {...product} />)
      )}
    </section>
  )
}

export default Showcase

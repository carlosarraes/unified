import type { Product } from './Showcase'
import Placeholder from '../assets/placeholder.png'
import { formatCurrency } from '../utils/formatCurrency'

const ProductCard = ({ thumbnail, title, price, permalink }: Product) => {
  return (
    <div className="flex w-10/12 bg-white gap-4 p-4 rounded-md shadow-md">
      {thumbnail.startsWith('data:') ? (
        <img src={Placeholder} alt={title} className="ml-3 w-32" />
      ) : (
        <img src={thumbnail} alt={title} className="ml-3 w-32" />
      )}
      <div className="flex flex-col items-stretch justify-evenly w-9/12">
        <h1 className="text-sm overflow-hidden whitespace-nowrap overflow-ellipsis h-12">
          {title}
        </h1>
        <div className="flex justify-between mt-8 gap-4">
          <p className="mt-4 text-right text-xl">{formatCurrency(price)}</p>
          <a href={permalink} className="text-right">
            <button className="btn btn-info bg-sky-900 text-white">Ir a Web</button>
          </a>
        </div>
      </div>
    </div>
  )
}

export default ProductCard

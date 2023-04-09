import { useState } from 'react'
import Header from './components/Header'
import Showcase from './components/Showcase'
import Spinner from './components/Spinner'

const API_URL = 'http://localhost:8080/search'

function App() {
  const [formControls, setFormControls] = useState({
    web: '',
    categorias: '',
    search: '',
  })
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(false)

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target
    setFormControls({ ...formControls, [name]: value })
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setLoading(true)
    const { web, categorias } = formControls

    const response = await fetch(`${API_URL}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ web, category: categorias }),
    })
    const data = await response.json()

    console.log(data)
    setProducts(data)
    setLoading(false)
  }

  return (
    <main className="flex flex-col justify-center items-center w-8/12">
      <Header
        formControls={formControls}
        handleChange={handleInputChange}
        handleSubmit={handleSubmit}
      />
      {loading ? <Spinner size={48} /> : <Showcase products={products} />}
    </main>
  )
}

export default App

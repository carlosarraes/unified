import { useState } from 'react'
import Header from './components/Header'

function App() {
  const [formControls, setFormControls] = useState({
    web: '',
    categorias: '',
    search: '',
  })

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = event.target
    setFormControls({ ...formControls, [name]: value })
  }

  return (
    <main className="flex flex-col justify-center items-center w-screen">
      <Header formControls={formControls} handleChange={handleInputChange} />
    </main>
  )
}

export default App

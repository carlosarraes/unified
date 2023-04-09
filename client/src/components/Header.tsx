import Input from './Input'
import Select from './Select'

type HeaderProps = {
  formControls: {
    web: string
    categorias: string
    search: string
  }
  handleChange: (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => void
  handleSubmit: (e: React.FormEvent<HTMLFormElement>) => void
}

const Header = ({ formControls, handleChange, handleSubmit }: HeaderProps) => {
  const webOptions = {
    disabledOption: 'Web',
    options: [
      { value: 'meli', label: 'MercadoLivre' },
      { value: 'busca', label: 'BuscaPé' },
    ],
  }

  const categoriesOptions = {
    disabledOption: 'Categorias',
    options: [
      { value: 'geladeira', label: 'Geladeira' },
      { value: 'tv', label: 'TV' },
      { value: 'celular', label: 'Celular' },
    ],
  }

  return (
    <header className="flex justify-center w-full p-4 mt-2 shadow-sm">
      <form className="flex justify-center items-center gap-4 w-full" onSubmit={handleSubmit}>
        <Select
          name={webOptions.disabledOption.toLowerCase()}
          disabledOption={webOptions.disabledOption}
          options={webOptions.options}
          control={formControls.web}
          handleChange={handleChange}
        />
        <Select
          name={categoriesOptions.disabledOption.toLowerCase()}
          disabledOption={categoriesOptions.disabledOption}
          options={categoriesOptions.options}
          control={formControls.categorias}
          handleChange={handleChange}
        />
        <Input control={formControls.search} handleChange={handleChange} />
        <button className="btn btn-info bg-sky-600 text-white" type="submit">
          Search
        </button>
      </form>
    </header>
  )
}

export default Header

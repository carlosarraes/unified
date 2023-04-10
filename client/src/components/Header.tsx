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
  validateBtn: () => boolean
}

const Header = ({ formControls, handleChange, handleSubmit, validateBtn }: HeaderProps) => {
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
    <header className="flex justify-center w-full p-2 mt-2 shadow-sm">
      <form
        className="flex flex-col justify-center items-center gap-4 w-full md:flex-row"
        onSubmit={handleSubmit}
      >
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
        <button
          className="btn btn-info bg-sky-600 text-white w-full md:w-1/5 max-w-xs"
          type="submit"
          disabled={validateBtn()}
        >
          Search
        </button>
      </form>
    </header>
  )
}

export default Header

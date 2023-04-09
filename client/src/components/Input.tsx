type InputProps = {
  control: string
  handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void
}

const Input = ({ control, handleChange }: InputProps) => {
  return (
    <input
      name="search"
      type="text"
      placeholder="Buscar Produto"
      value={control}
      onChange={handleChange}
      className="input input-bordered input-success w-full max-w-md border-sky-900"
    />
  )
}

export default Input

type option = {
  value: string
  label: string
}

type SelectProps = {
  name: string
  disabledOption: string
  options: option[]
  control: string
  handleChange: (event: React.ChangeEvent<HTMLSelectElement>) => void
}

const Select = ({ name, disabledOption, options, control, handleChange }: SelectProps) => {
  return (
    <select
      className="select w-1/4 max-w-xs text-white bg-sky-900"
      aria-label={name}
      name={name}
      value={control}
      onChange={handleChange}
    >
      <option disabled value="">
        {disabledOption}
      </option>
      {options.map((option) => (
        <option key={option.value} value={option.value}>
          {option.label}
        </option>
      ))}
    </select>
  )
}

export default Select

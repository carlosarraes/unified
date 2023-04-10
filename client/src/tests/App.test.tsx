import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import App from '../App'

describe('App tests', () => {
  it('renders App component and find starting elements', () => {
    render(<App />)
    const selectWebElement = screen.getByRole('combobox', { name: /web/i })
    const selectCategoriesElement = screen.getByRole('combobox', { name: /categorias/i })
    const inputSearchElement = screen.getByRole('textbox', { name: /search/i })
    const buttonSearchElement = screen.getByRole('button', { name: /search/i })
    const h1Element = screen.getByRole('heading', { name: /sua pesquisa/i })

    expect(selectWebElement).toBeInTheDocument()
    expect(selectCategoriesElement).toBeInTheDocument()
    expect(inputSearchElement).toBeInTheDocument()
    expect(buttonSearchElement).toBeInTheDocument()
    expect(h1Element).toBeInTheDocument()
  })

  it('button starts disabled, then enabled after selecting web and category', async () => {
    render(<App />)
    const buttonSearchElement = screen.getByRole('button', { name: /search/i })

    expect(buttonSearchElement).toBeDisabled()

    const selectWebElement = screen.getByRole('combobox', { name: /web/i })
    const selectCategoriesElement = screen.getByRole('combobox', { name: /categorias/i })

    await userEvent.selectOptions(selectWebElement, 'meli')
    await userEvent.selectOptions(selectCategoriesElement, 'geladeira')

    expect(buttonSearchElement).toBeEnabled()
  })

  it('Renders one item after search', async () => {
    render(<App />)
    const selectWebElement = screen.getByRole('combobox', { name: /web/i })
    const selectCategoriesElement = screen.getByRole('combobox', { name: /categorias/i })
    const buttonSearchElement = screen.getByRole('button', { name: /search/i })

    await userEvent.selectOptions(selectWebElement, 'meli')
    await userEvent.selectOptions(selectCategoriesElement, 'geladeira')
    await userEvent.click(buttonSearchElement)

    const itemElement = await screen.findByRole('heading', { name: /geladeira/i })

    expect(itemElement).toBeInTheDocument()
  })

  it('Renders one item after search with buscape', async () => {
    render(<App />)
    const selectWebElement = screen.getByRole('combobox', { name: /web/i })
    const selectCategoriesElement = screen.getByRole('combobox', { name: /categorias/i })
    const buttonSearchElement = screen.getByRole('button', { name: /search/i })

    await userEvent.selectOptions(selectWebElement, 'busca')
    await userEvent.selectOptions(selectCategoriesElement, 'geladeira')
    await userEvent.click(buttonSearchElement)

    const itemElement = await screen.findByRole('heading', { name: /geladeira/i })
    const placeholderElement = await screen.findByTestId('placeholder')

    expect(itemElement).toBeInTheDocument()
    expect(placeholderElement).toBeInTheDocument()
  })
})

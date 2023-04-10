import { rest } from 'msw'

const url = 'https://gouni-4wgfen3n5q-rj.a.run.app/'

const meliMockData = [
  {
    title: 'Geladeira',
    price: 1000,
    permalink: 'http://geladeira',
    thumbnail: 'http://geladeira.jpg',
  },
]

const buscaMockData = [
  {
    title: 'Geladeira',
    price: 1000,
    permalink: 'buscapelink',
    thumbnail: 'data:link',
  },
]

export const handlers = [
  rest.post(url + 'search', async (req, res, ctx) => {
    const body = await req.json()
    if (typeof body === 'object' && body !== null && body.web === 'meli') {
      return res(ctx.status(200), ctx.json(meliMockData))
    }
    return res(ctx.status(200), ctx.json(buscaMockData))
  }),
]

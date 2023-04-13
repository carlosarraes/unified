# Unified Store

# Errata

Added MercadoLivre (Scrape) option too.

## Frontend

The frontend is built with React and TypeScript, and is deployed on Vercel. It utilizes Tailwind CSS and DaisyUI for styling, and Lucide React for icons. The design is completely modular, making it easy to add more APIs and categories in the future.

Please note that it may take a few seconds to retrieve data from the backend due to cold start delays.

### Technologies

- React (with TypeScript)
- Tailwind CSS
- DaisyUI
- Lucide React

### Modular Design

The frontend's modular design ensures that adding new APIs or categories is a straightforward process. This flexible architecture enables the application to grow and adapt to new requirements with ease.

**Note:** Due to cold start delays, it might take a few seconds for the frontend to fetch data from the backend.

## Backend

This backend is written in Go, utilizing the Go-chi and Gorm libraries. It is designed as a stateless API with a single powerful endpoint (POST) at `/search`. The backend follows the MSC (Model, Service, Controller) architecture.

### Endpoint

The search POST endpoint expects a request body containing the data source (MercadoLivre or Buscape) and the category to search in.

When a user chooses MercadoLivre and it's their first search, the API queries [MercadoLivre's API](https://developers.mercadolivre.com.br/pt_br/api-docs-pt-br), filters results by category, saves the data to the database, and sends it to the frontend. If the data already exists in the database, it retrieves it from there and sends it to the frontend. If the user chooses Buscape, the same process occurs, but instead of using the deprecated Buscape API, a web scraper is employed to obtain the data.

Handling search terms or searching by product name is delegated to the frontend.

### Database and Dockerfile

The stateless API connects to PlanetScale to maintain persistent data. The database schema is managed and auto-migrated based on GORM-defined structs. We will use a Dockerfile with a Go image to host on Google Cloud.

<table align="center">
<thead>
<tr>
<th>Database</th>
<th>Dockerfile</th>
</tr>
</thead>
<tbody>
<tr>
<td>

```sql
CREATE TABLE search_histories (
	id bigint unsigned NOT NULL AUTO_INCREMENT,
	created_at datetime(3),
	updated_at datetime(3),
	deleted_at datetime(3),
	web longtext,
	category longtext,
	search_results longblob,
	PRIMARY KEY (id),
	KEY idx_search_histories_deleted_at (deleted_at)
);
```

</td>
<td>

```Dockerfile
FROM golang:1.20.1-alpine3.17

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ./out/dist .

EXPOSE 8080

CMD ./out/dist
```

</td>
</tr>
</tbody>
</table>

### Web Scraper

The web scraper is built using the GoColly library. It retrieves data based on the specified category and fetches the first six images. Additional images are lazy-loaded, so the frontend should display a placeholder image as needed.

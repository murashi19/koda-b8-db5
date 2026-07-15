# Contact List CLI

Simple contact management application built with Go and PostgreSQL. This project was created as part of a bootcamp assignment to practice database integration, CRUD operations, Docker, and GitHub Actions.

## Features

* List all contacts
* Add a new contact
* Update contact information
* Delete a contact

## Tech Stack

* Go
* PostgreSQL
* pgx
* Docker
* GitHub Actions
* GitHub Container Registry (GHCR)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/<username>/<repository>.git

cd <repository>
```

### 2. Run PostgreSQL

```bash
docker run -d \
  --name postgres-db2 \
  -e POSTGRES_USER=postgres2 \
  -e POSTGRES_PASSWORD=1 \
  -e POSTGRES_DB=postgres3 \
  -p 5432:5432 \
  postgres:alpine
```

### 3. Import database schema

```bash
docker exec -i postgres-db2 \
psql -U postgres2 -d postgres3 < init.sql
```

### 4. Configure environment

Create a `.env` file in the project root.

```env
DATABASE_URL=postgres://postgres2:1@localhost:5432/postgres3?sslmode=disable
```

### 5. Run the application

```bash
go run main.go
```

Or build the binary first.

```bash
go build -o contact-app

./contact-app
```

## Running with Docker

Build the image.

```bash
docker build -t contact-app .
```

Create a Docker network.

```bash
docker network create contact-network
```

Connect the PostgreSQL container to the network.

```bash
docker network connect contact-network postgres-db2
```

Update the host in your `.env` file:

```env
DATABASE_URL=postgres://postgres2:1@postgres-db2:5432/postgres3?sslmode=disable
```

Run the container.

```bash
docker run -it --rm \
  --network contact-network \
  --env-file .env \
  contact-app
```

## CI/CD

This project uses GitHub Actions to automate the release process. When a tag starting with `v` is pushed, the workflow will:

* Build the Go application
* Upload the build artifact
* Create a GitHub Release
* Build the Docker image
* Push the image to GitHub Container Registry (GHCR)

Create a new release by pushing a version tag:

```bash
git tag v1.0.0
git push origin v1.0.0
```

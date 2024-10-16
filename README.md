# TicketTracker

![Alt text](/ticket-tracker-in-action.gif "GIF video of TicketTracker")

## Why?

Most incident management tools I tried were:

- _molasses_ slow
- **clunky**
- **impossible** to navigate

I needed something:

- **quicker** (loads everything in less than 250 ms)
- **smoother**
- way more **intuitive**

That's why I built **TicketTracker**.

Now, tracking and resolving incidents is as lightning fast as it should be.

**TicketTracker** is a IT ticketing system designed to simplify the management of incidents.

## Features

- **Incident Management:** Log, track, and resolve incidents with ease.
- **Fast:** Blazing quick **[Go](https://golang.org/)** backend
- **Intuitive:** Smooth, modern **[React](https://reactjs.org/)** frontend.

## Dependencies

This program uses the following packages among many others. Please refer to the `go.mod` and `package.json` files for a complete list of dependencies.

### Go Modules

The following Go modules are required for this project:

- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - Version 5.2.1
- [google/uuid](https://github.com/google/uuid) - Version 1.6.0
- [joho/godotenv](https://github.com/joho/godotenv) - Version 1.5.1
- [x/crypto](https://pkg.go.dev/golang.org/x/crypto) - Version 0.27.0

### NPM Packages

The frontend uses the following npm packages, among others:

- [@headlessui/react](https://github.com/tailwindlabs/headlessui) - Version 2.1.9
- [@heroicons/react](https://github.com/tailwindlabs/heroicons) - Version 2.1.5
- [@tailwindcss/forms](https://github.com/tailwindlabs/tailwindcss-forms) - Version 0.5.9

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/barturba/ticket-tracker.git
   ```

2. **Install dependencies:**

   ```bash
   cd ticket-tracker
   go mod tidy
   ```

3. **Configure the application:**

- Create a .env file in the root directory of your project with the following format:

  ```bash .env
  # Application Environment
  ENV="production"

  # Server Configuration
  SERVER_HOST="0.0.0.0"
  SERVER_PORT="8080"

  # Security Settings
  JWT_SECRET="GENERATE-SECRET-HERE"    # Use a secure method to generate this secret

  # Database Configuration
  DB_HOST="db"
  DB_PORT="5432"
  DB_USER="tickettrackeradmin"
  DB_PASSWORD="GENERATE-PASSWORD-HERE" # Use a secure method to generate this password
  DB_NAME="tickettrackerdb"

  # Database URLs
  DATABASE_URL_DEV="postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"
  DATABASE_URL_PROD="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

  ```

- **ENV**: Defines the application environment. Use "development" for local testing or "production" for deployment.
- **SERVER_HOST**: The server's host address. Use "0.0.0.0" to allow external access.
- **SERVER_PORT**: Specifies the port the server will listen on.
- **JWT_SECRET**: Secret key for signing JSON Web Tokens. Use a strong, randomly generated string here.
- **DB_HOST**: The hostname for the database. Set to "db" if using Docker Compose.
- **DB_PORT**: The port the database is running on (typically 5432 for PostgreSQL).
- **DB_USER**: Database username for authentication.
- **DB_PASSWORD**: Password for the database user. Use a strong password.
- **DB_NAME**: Name of the database to connect to.
- **DATABASE_URL_DEV**: Connection string for the development database.
- **DATABASE_URL_PROD**: Connection string for the production database, pointing to the host specified in DB_HOST.

Use the following command to generate a secure JWT secret:

```bash
openssl rand -base64 60
```

4. **Build and run the application**
   Once you have (Docker)[https://www.docker.com/] and (Docker Compose)[https://docs.docker.com/compose/] installed, setting up TicketTracker is as easy as running:

   ```bash
   docker-compose up --build
   ```

This command will build and start all the necessary services. You can then access the application by navigating to http://localhost:3000 in your web browser.

5. **Access the application:** Open a web browser and go to `http://localhost:3000`.

## Support

For support, please open an issue on the [GitHub repository](https://github.com/barturba/ticket-tracker/issues).

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/barturba/ticket-tracker@latest
cd ticket-tracker
```

### Build the project

```bash
make build/ticket-tracker
```

### Run the project

```bash
make
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

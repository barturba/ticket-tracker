# TicketTracker Readme

![Alt text](/ticket-tracker-screenshot.png?raw=true "Screenshot of TicketTracker")

## Why?

Most incident management tools I tried were molasses slow, clunky, and to impossible navigate. I needed something quicker, smoother, and more intuitive —- so I built **TicketTracker**. Now, tracking and resolving incidents is as lightning fast as it should be.

**TicketTracker** is a IT ticketing system designed to simplify the management of incidents. It provides basic user support and essential features for IT service management.

## Features

- **Incident Management:** Log, track, and resolve incidents with ease.
- **Configuration Item Management:** Maintain an inventory of configuration items (CIs).

## Dependencies

This program uses the following packages among many others. Please refer to the `go.mod` and `package.json` files for a complete list of dependencies.

### Go Modules

- The following Go modules are required for this project:

- github.com/golang-jwt/jwt/v5 - Version 5.2.1
- github.com/google/uuid - Version 1.6.0
- github.com/joho/godotenv - Version 1.5.1
- golang.org/x/crypto - Version 0.27.0

### NPM Packages

The frontend uses the following npm packages:

- @headlessui/react - Version 2.1.9
- @heroicons/react - Version 2.1.5
- @tailwindcss/forms - Version 0.5.9
- @types/node - Version 20
- @types/react - Version 18
- @types/react-dom - Version 18
- clsx - Version 2.1.1
- eslint - Version 8
- eslint-config-next - Version 14.2.14
- framer-motion - Version 11.11.8

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
  ENV="development"
  HOST="localhost"
  PORT="8080"
  DATABASE_URL="postgres://YOURNAME:@localhost:5432/tickettracker?sslmode=disable"
  JWT_SECRET="GENERATEME"
  ADMIN_PASSWORD="pass123"
  ```

- **ENV:** Set the environment. For development, use "development".
- **HOST:** Specify the host for the server.
- **PORT:** Set the port the application will run on.
- **DATABASE_URL:** Provide your PostgreSQL connection string, replacing YOURNAME with your database username.
- **JWT_SECRET:** Generate a secret for JSON Web Token encryption. Use the following command to generate a secure key:

  ```bash
  openssl rand -base64 60
  ```

- **ADMIN_PASSWORD:** Set the password for the admin user.

4. **Build and run the backend application:**

   ```bash
   make
   ```

5. **Build and run the frontend:**

- Open a new teriminal in the roow of the project and run the following commands:

  ```bash
  cd frontend
  pnpm install
  pnpm dev
  ```

6. **Access the application:** Open a web browser and go to `http://localhost:3000`.

## Disclaimer

This is a basic implementation of a ticket tracker. It is not intended for production use. Please use at your own risk.

## Support

For support, please open an issue on the GitHub repository.

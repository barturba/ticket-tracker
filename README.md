# TicketTracker

[![Go Report Card](https://goreportcard.com/badge/github.com/barturba/ticket-tracker)](https://goreportcard.com/report/github.com/barturba/ticket-tracker)

![TicketTracker in action](/ticket-tracker-in-action.gif "GIF video of TicketTracker")

[Live Demo](https://frontend-patient-voice-578.fly.dev)

## Overview

TicketTracker is a lightning-fast, intuitive IT ticketing system designed to simplify incident management. Built with a Go backend and React frontend, it offers a smooth, modern experience for tracking and resolving incidents.

### Why TicketTracker?

- **Speed**: Loads everything in less than 250 ms
- **Simplicity**: Intuitive interface for easy navigation
- **Efficiency**: Streamlined incident logging, tracking, and resolution

## Features

- Rapid incident management
- Blazing-fast Go backend
- Smooth, modern React frontend

## Tech Stack

### Backend

- Go
- Key packages: golang-jwt/jwt, google/uuid, joho/godotenv, x/crypto

### Frontend

- React
- Key packages: @headlessui/react, @heroicons/react, @tailwindcss/forms

## Getting Started

1. **Clone the repository**

   ```bash
   git clone https://github.com/barturba/ticket-tracker.git
   cd ticket-tracker
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure the application**
   Create a `.env` file in the root directory with the following structure:

   ```env
   ENV="production"
   SERVER_HOST="0.0.0.0"
   SERVER_PORT="8080"
   JWT_SECRET="GENERATE-SECRET-HERE"
   DB_HOST="db"
   DB_PORT="5432"
   DB_USER="tickettrackeradmin"
   DB_PASSWORD="GENERATE-PASSWORD-HERE"
   DB_NAME="tickettrackerdb"
   DATABASE_URL_DEV="postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"
   DATABASE_URL_PROD="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
   ```

   Generate a secure JWT secret:

   ```bash
   openssl rand -base64 60
   ```

4. **Build and run the application**

   ```bash
   docker-compose up --build
   ```

5. **Access the application**
   Open a web browser and navigate to `http://localhost:3000`.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## Support

For support, please [open an issue](https://github.com/barturba/ticket-tracker/issues) on the GitHub repository.

## License

[MIT License](LICENSE)

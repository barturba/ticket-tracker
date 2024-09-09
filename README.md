# TicketTracker Readme

**TicketTracker** is a IT ticketing system designed to simplify the management of incidents, change requests, and configuration items. It provides basic user support and essential features for IT service management.

## Features

- **Incident Management:** Log, track, and resolve incidents with ease. Assign priorities, set due dates, and collaborate with team members to ensure timely resolution.
- **Change Request Management:** Initiate, review, approve, and implement change requests following established processes. Minimize risks and maintain system stability.
- **Configuration Item Management:** Maintain an inventory of configuration items (CIs). Track relationships between CIs, monitor changes, and support IT asset management.
- **Basic User Support:** Provide users with a self-service portal to submit tickets, track progress, and view knowledge base articles for common issues.

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/barturba/ticket-tracker.git
   ```

2. **Install dependencies:**

   ```bash
   cd tickettracker
   go mod tidy
   ```

3. **Configure the application:**

- Update .env with your database connection details.
- Set up user authentication (if required).

4. **Build and run the application:**

   ```bash
   go build .
   ./ticket-tracker
   ```

5. **Access the application:** Open a web browser and go to `http://localhost:8080`.

## Disclaimer

This is a basic implementation of a ticket tracker. It is not intended for production use. Please use at your own risk.

## Support

For support, please open an issue on the GitHub repository.

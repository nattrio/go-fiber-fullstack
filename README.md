# Go Fiber Fullstack Trivia App

This is a simple trivia web application built using the Go programming language and the [Fiber web framework](https://github.com/gofiber/fiber). The application allows users to view a list of trivia facts, add new facts, view single facts, edit existing facts, and delete facts.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/nattrio/go-fiber-fullstack.git
   ```

2. Install the required dependencies:

   ```bash
   go mod download
   ```

3. Set up the environment variables:

   ```bash
   cp .env.example .env
   ```

4. Update the `.env` file with your local database credentials.

5. Start the application:

   ```bash
   docker compose up
   ```

## Usage

After starting the application, navigate to `http://localhost:3000` in your web browser to use the application.

## Routes

The following routes are available in the application:

- `GET /`: List all trivia facts
- `GET /fact`: Show the form for adding a new trivia fact
- `POST /fact`: Add a new trivia fact
- `GET /fact/:id`: Show a single trivia fact
- `GET /fact/:id/edit`: Show the form for editing an existing trivia fact
- `PATCH /fact/:id`: Update an existing trivia fact
- `DELETE /fact/:id`: Delete an existing trivia fact

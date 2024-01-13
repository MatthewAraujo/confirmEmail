
# Go Confirm Email Application

This is a Go application that serves as a simple HTTP server providing an endpoint to handle user data, specifically for confirming email addresses. The application exposes a `POST` endpoint (`/postUser`), expecting JSON data with fields `firstName`, `lastName`, and `email`. Upon receiving user data, it processes the information and sends a confirmation email to the provided address.

## Features

- **User Data Handling:** Accepts `POST` requests with JSON payload containing user information (`firstName`, `lastName`, `email`).
- **Email Confirmation:** Sends a confirmation email to the provided address welcoming the user to the club.

## Prerequisites

- Go installed on your machine.

## How to Run

1. Clone this repository:

    ```bash
    git clone https://github.com/matthewAraujo/confirmEmail.git
    cd confirmEmail
    ```

2. Create a `.env` file in the project directory and set the following environment variables:

    ```env
    SMTP_PASSWORD=your_smtp_password
    SMTP_HOST=your_smtp_host
    SMTP_PORT=your_smtp_port
    SMTP_AUTHOR=your_email_address
    ```

3. Run the Go application:

    ```bash
    go run main.go
    ```

4. The server will start on `http://localhost:8080`.

## Usage

### POST /postUser

Send a `POST` request to `http://localhost:8080/postUser` with a JSON payload containing `firstname`, `lastname`, and `email` fields. For example:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"firstname":"John","lastname":"Doe","email":"john.doe@example.com"}' http://localhost:8080/postUser
```

Replace the example data with your own user information.

## Customization

Modify the `handlePostUser` and `sendMail` functions in `main.go` to customize the processing logic for user data and email confirmation.

## Dependencies

This project uses the `github.com/joho/godotenv` package for environment variable loading and has no other external dependencies.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

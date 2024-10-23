# Customer Feedback API

This is a simple API for managing customer feedback. The API allows users to submit feedback, retrieve all feedback, and validate email addresses.

## Prerequisites

- Go version 1.16 or higher
- A database (SQLite is used in this example)

## Installation

1. Clone the repository:

```bash
git clone git@github.com:Venturing-Intellect/go-codium.git
```

2. Change into the project directory:

```
cd customer-feedback-api
```

3. Run the API:

```
go run main.go
```


## Usage

The API has two endpoints:

- `POST /feedback/create`: Creates a new feedback record. The request body should be a JSON object with the following fields:
- `Name` (string): The name of the person submitting the feedback.
- `Email` (string): The email address of the person submitting the feedback.
- `Feedback` (string): The feedback message.

- `GET /feedback`: Retrieves all feedback records.

## Testing

The API includes unit tests for the service and integration tests for the API endpoints. To run the tests, execute the following command:

```
go test ./...

```


## Email Validation

The API includes a function for validating email addresses. The function uses the following regular expression:

```
^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-]+$
```


This regular expression matches email addresses with the following format:
- The local part (before the `@`) can contain alphanumeric characters, underscores, periods, and plus signs.
- The domain part (after the `@`) can contain alphanumeric characters and hyphens.
- The top-level domain (after the last dot) can contain alphanumeric characters and hyphens.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue. If you'd like to contribute code, please submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
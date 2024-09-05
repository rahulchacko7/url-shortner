# URL Shortener using Gin

A simple URL shortener service built with the Gin framework in Go. This application allows you to shorten long URLs and retrieve the original URL using the shortened version.

## Features

- **Shorten URL**: Convert a long URL into a shorter version.
- **Retrieve URL**: Redirect from a shortened URL to the original long URL.
- **In-Memory Storage**: Stores URL mappings in memory for quick access.
- **Concurrency Safe**: Utilizes synchronization to handle concurrent access.

## Requirements

- Go 1.20 or later
- Docker (for containerization)

## Installation

### Without Docker

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener
   ```

2. **Build the Go application:**

   ```bash
   go build -o url-shortener
   ```

3. **Run the application:**

   ```bash
   ./url-shortener
   ```

   The application will run on `http://localhost:8080`.

### With Docker

1. **Build the Docker image:**

   ```bash
   sudo docker build -t url-shortener .
   ```

2. **Run the Docker container:**

   ```bash
   sudo docker run -p 8080:8080 url-shortener
   ```

   The application will be accessible at `http://localhost:8080`.

## API Endpoints

### Shorten URL

- **Endpoint**: `POST /shorten`
- **Request Body**:
  ```json
  {
    "url": "<long_url>"
  }
  ```
- **Response**:
  ```json
  {
    "short_url": "<short_url>"
  }
  ```

### Retrieve URL

- **Endpoint**: `GET /:short_url`
- **Response**: Redirects to the original URL or returns a 404 error if the short URL is not found.

## Example

**Shorten URL**

```bash
curl -X POST http://localhost:8080/shorten -d '{"url": "https://example.com"}' -H "Content-Type: application/json"
```

**Retrieve URL**

```bash
curl -X GET http://localhost:8080/abc123
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

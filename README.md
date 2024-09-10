# URL Shortener Service

A simple URL shortener service written in Go. This service allows you to shorten long URLs and track their usage.

## Features

- **Create Short URLs**: Generate a short code for a long URL with an optional expiration date.
- **Redirect**: Redirect users to the original URL using the short code.
- **Track Usage**: Track the number of times a short URL has been accessed.
- **Handle Expiration**: Automatically handle URL expiration.

## Architecture

- **Service Layer**: Handles business logic for creating and retrieving short URLs.
- **DAO Layer**: Manages database operations for short URLs.
- **Handlers**: Defines HTTP endpoints for creating short URLs and redirection.

## API Endpoints

### Create Short URL

- **Endpoint**: `POST /shorten`
- **Request Body**:

    ```json
    {
        "original_url": "https://example.com",
        "expires_at": "2024-12-31T23:59:59Z"
    }
    ```

- **Response**:

    ```json
    {
        "short_code": "abc123",
        "expires_at": "2024-12-31T23:59:59Z"
    }
    ```

### Redirect to Original URL

- **Endpoint**: `GET /:shortCode`
- **Response**: Redirects to the original URL.

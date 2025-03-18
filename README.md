# FRC-REEF-METRICS Data API

This is a Go-based API designed to retrieve data collected by our robotics team (1156). It is built using the Gin web framework and provides a Swagger UI for easy documentation and testing. The API runs in a Docker container and uses PostgreSQL as its database.

## Features
- RESTful API built with [Gin Gonic](https://github.com/gin-gonic/gin)
- API documentation using [Swagger](https://swagger.io/)
- Containerized with Docker
- PostgreSQL as the database

## Current Status
The API is functional but currently operates with fictitious data, as the real dataset remains private due to the next competition.

## Requirements
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL image

## Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/ljb6/frc-reef-metrics
   cd frc-reef-metrics
   ```
2. Build the API image:
   ```sh
   docker build -t frc-reef-metrics-api .
   ```
3. Start the database and API using Docker:
   ```sh
   docker-compose up -d
   ```
4. The API will be available at `http://localhost:8080`.
5. Access the Swagger UI at `http://localhost:8000/swagger/index.html`.

## Endpoints (Example)
- Endpoint: `"/api/reef-metrics/v1"`
- `GET /all-matches` - Retrieve all stored data
- `GET /matches/{team}` - Fetch specific team matches by team number
- `GET /match/{match}` - Fetch specific data from a match by match number
- `GET /match/{match}/{team}` - Fetch specific match of a team bt match and team number

## Contributing
Contributions are welcome! Feel free to submit issues or pull requests.


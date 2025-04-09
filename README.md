# Weather API - Go + Gin

This is a simple and clean weather forecast API written in **Go**, following idiomatic best practices and a modular architecture. It allows you to query weather data using a **Brazilian ZIP code (CEP)**.

---

## Cloud Run URL
To use the public URL hosted on Cloud Run - Service URL: https://weather-api-574418562183.us-central1.run.app

For successful testing:
```bash
curl https://weather-api-574418562183.us-central1.run.app/weather/31340110
```
Response:
```json
{"celsius":28.1,"fahrenheit":82.6,"kelvin":301.1}
```

For a zip code that does not exist:
```bash
curl https://weather-api-574418562183.us-central1.run.app/weather/00000000
```
Response:
```json
{"message":"can not find zipcode"}
```

For invalid ZIP codes
```bash
curl https://weather-api-574418562183.us-central1.run.app/weather/0
```
Response:
```json
{"message":"invalid zipcode"}
```

## Architecture Overview

The project is structured with a **clean separation of concerns**, inspired by principles of **hexagonal/clean architecture**:

```
cmd/
└── server/             -> Entry point of the app (main.go)
internal/
├── adapters/           -> API clients (BrasilAPI and WeatherAPI)
├── domain/             -> Business models, interfaces, and domain errors
├── service/            -> Business logic layer (WeatherService)
├── handler/            -> HTTP handlers (Gin)
├── router/             -> HTTP route definitions
├── server/             -> HTTP server setup
```

---

## Technologies Used

- **Go 1.21+**
- **Gin** – for HTTP server and routing
- **Docker & Docker Compose** – for containerized environment
- **BrasilAPI** – to resolve Brazilian CEP to address and coordinates
- **WeatherAPI** – to get weather data using latitude and longitude

---

## Features

- `GET /weather/:cep`: Returns current weather forecast by CEP (ZIP code)
- Modular and testable architecture
- Mocks and test coverage
- Environment variables support via `.env`

---

## Requirements

To run the app locally, make sure you have:

- [Go](https://go.dev/dl/) 1.21 or higher
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Installation & Usage

### 1. Clone the repository

```bash
git clone https://github.com/fmantinossi/weather-app.git
cd weather-app
```

---

### 2. Create a `.env` file

Create a file named `.env` in the root of the project:

```env
WEATHER_API_KEY=your_weatherapi_key_here
PORT=8080
```

> You can get your free API key from: https://www.weatherapi.com/

---

### 3. Run using Docker Compose

```bash
docker-compose up --build
```

The app will be available at: [http://localhost:8080](http://localhost:8080)

---

### 4. Example Request

```bash
curl http://localhost:8080/weather/01001000
```

Response:
```json
{
  "celsius": 25.0,
  "fahrenheit": 77.0,
  "kelvin": 298.0
}
```

---

### 5. Run tests

To run all tests:

```bash
go test ./... -v
```

To run inside the container:

```bash
docker exec -it weather-api sh
go test ./... -v
```

---

## Development Logic

This API was designed with **testability, maintainability, and clarity** in mind:

- **Adapters**: External APIs are abstracted via interfaces (`AddressProvider`, `WeatherProvider`)
- **Domain**: Central business definitions and errors
- **Service layer**: Orchestrates adapters and processes business rules
- **Handler**: Handles HTTP requests/responses and maps domain errors to status codes
- **Router**: Defines clean routing using Gin
- **Server**: Central point that wires everything and launches the app

---

## 👤 Author

**Fernando Mantinossi**  
[GitHub Profile](https://github.com/fmantinossi)

---
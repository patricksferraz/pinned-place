# Pinned Place 🎯

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/pinned-place)](https://goreportcard.com/report/github.com/patricksferraz/pinned-place)
[![GoDoc](https://godoc.org/github.com/patricksferraz/pinned-place?status.svg)](https://godoc.org/github.com/patricksferraz/pinned-place)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, scalable location-based service built with Go, designed to handle location data efficiently and reliably.

## 🌟 Features

- RESTful API for location management
- PostgreSQL database for reliable data storage
- Kafka integration for event streaming
- Docker and Kubernetes support for easy deployment
- Hot-reload development environment with Air
- Comprehensive database migrations
- Admin interface for database management

## 🚀 Quick Start

### Prerequisites

- Go 1.18+
- Docker and Docker Compose
- Make (optional, but recommended)

### Environment Setup

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/pinned-place.git
cd pinned-place
```

2. Copy the example environment file and configure it:
```bash
cp .env.example .env
```

3. Start the services using Docker Compose:
```bash
docker-compose up -d
```

### Development

The project uses Air for hot-reloading during development. To start the development server:

```bash
make dev
```

## 🏗️ Project Structure

```
.
├── app/          # Application layer
├── cmd/          # Command-line interface
├── domain/       # Domain models and business logic
├── infra/        # Infrastructure implementations
├── k8s/          # Kubernetes configurations
└── utils/        # Utility functions and helpers
```

## 🛠️ Technology Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Message Broker**: Kafka
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Development**: Air (hot-reload)
- **Database Admin**: Adminer

## 📚 API Documentation

The API documentation is available at `/swagger` when running the application.

## 🔧 Configuration

The application can be configured through environment variables:

- `REST_PORT`: Port for the REST API
- `POSTGRES_*`: PostgreSQL configuration
- `KAFKA_*`: Kafka configuration
- `DSN_*`: Data Source Name configuration

See `.env.example` for all available configuration options.

## 🧪 Testing

Run the test suite:

```bash
make test
```

## 📦 Deployment

### Docker

Build and run using Docker:

```bash
make docker-build
make docker-run
```

### Kubernetes

Deploy to Kubernetes:

```bash
kubectl apply -f k8s/
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Ferraz** - *Initial work*

## 🙏 Acknowledgments

- Thanks to all contributors who have helped shape this project
- Inspired by modern microservices architecture
- Built with best practices in mind

---

⭐ Star this repository if you find it useful!

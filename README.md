## Getting Started

### Prerequisites

- Go 1.21 or higher


### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd hng1
```

2. Install dependencies:
```bash
go mod tidy
```

5. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /` - Check if the API is running

### Authentication
- `GET /me` - Cat Fact

You can view this also on `https://hng1.up.railway.app`

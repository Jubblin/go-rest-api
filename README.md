# Activity Monitoring API

A Go-based REST API for monitoring device activities and usage statistics with built-in metrics collection using Prometheus and visualization using Grafana.

## Features

- 📊 Device activity tracking with grid-based organization
- 📈 Usage statistics collection
- 🔍 HTTP header capture for POST/PUT requests
- 📝 Swagger documentation
- 📊 Prometheus metrics
- 📈 Grafana dashboards
- 🗄️ ObjectBox persistence
- 🐳 Docker deployment

## Quick Start

### Using Just

```bash
# Install and setup
just install-all
just dev-setup

# Run the application
just run
```

### Using Docker

```bash
just deploy
```

## API Endpoints

### Activities

#### Create Activity
```bash
curl -X POST http://localhost:8080/api/v1/activities \
  -H "Content-Type: application/json" \
  -H "X-Device-ID: device123" \
  -d '{
    "source_ip": "192.168.1.100",
    "device_name": "device-alpha",
    "grid_name": "grid-east",
    "action": "login"
  }'
```

Example Response:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "unique_id": "123e4567-e89b-12d3-a456-426614174000",
  "source_ip": "192.168.1.100",
  "device_name": "device-alpha",
  "grid_name": "grid-east",
  "action": "login",
  "headers": {
    "Content-Type": "application/json",
    "X-Device-ID": "device123"
  },
  "timestamp": "2024-01-20T15:04:05Z"
}
```

Other Endpoints:
- `GET /api/v1/activities` - List all activities
- `GET /api/v1/activities/device/{device}` - Get activities by device
- `GET /api/v1/activities/grid/{grid}` - Get activities by grid
- `DELETE /api/v1/activities/{id}` - Delete an activity

### Usage Statistics

- `POST /api/v1/stats` - Record usage statistics
- `GET /api/v1/stats` - Get all statistics
- `GET /api/v1/stats/endpoints/{endpoint}` - Get statistics by endpoint
- `DELETE /api/v1/stats/{id}` - Delete statistics

## Development

### Available Commands

```bash
just                    # Show all available commands
just run               # Run locally
just test              # Run tests
just test-coverage     # Run tests with coverage
just lint              # Run linter
just fmt               # Format code
just swagger           # Update Swagger docs
```

### Docker Commands

```bash
just docker-build      # Build containers
just docker-up         # Start services
just docker-down       # Stop services
just docker-logs       # View logs
just docker-restart    # Restart services
```

## Monitoring

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (admin/admin)
- Metrics endpoint: http://localhost:8080/metrics

### Available Metrics

- `http_requests_total` - Total HTTP requests
- `http_request_duration_seconds` - Request duration
- `activity_operations_total` - Activity operations by type
- `activity_count` - Current activity count by grid/device
- `objectbox_operations_total` - Database operations

## Project Structure

```
.
├── controllers/       # Request handlers
├── models/           # Data models
├── repositories/     # Data access layer
├── metrics/          # Prometheus metrics
├── middleware/       # HTTP middleware
├── utils/           # Utility functions
├── docs/            # Swagger documentation
└── db/              # Database configuration
```

## Documentation

- API Documentation: http://localhost:8080/swagger/index.html
- Metrics Documentation: http://localhost:8080/metrics

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## AsTeRICS Grid Integration

The API includes a local copy of AsTeRICS Grid for development and testing.

### Accessing the Grid

- Grid UI: http://localhost:8082
- Default credentials: none required in local mode

### Grid Configuration

The Grid is configured to:
- Run in local mode
- Persist data across restarts
- Communicate with the API endpoints

### Example Grid Setup

1. Access the Grid UI at http://localhost:8082
2. Create a new grid
3. Add action buttons that trigger API calls:

```javascript
// Example Grid action to log activity
fetch('http://localhost:8080/api/v1/activities', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'X-Grid-ID': 'my-grid'
  },
  body: JSON.stringify({
    source_ip: window.location.hostname,
    device_name: 'grid-device',
    grid_name: 'my-grid',
    action: 'button_click'
  })
})
```

## Dependency Management

This project uses GitHub Dependabot to keep dependencies up to date. Dependabot will automatically create pull requests for:

- Go module updates (weekly)
- Docker base image updates (weekly)
- GitHub Actions updates (weekly)

### Configuration

The Dependabot configuration can be found in `.github/dependabot.yml`. It includes:

- Automated dependency version checks
- Weekly update schedule
- Automatic pull request creation
- Dependency-type specific labels
- Scoped commit messages

### Update Process

1. Dependabot creates a PR when updates are available
2. CI runs on the PR to verify changes
3. PR can be reviewed and merged if tests pass

## Container Security & Building

This project uses GitHub Actions for secure container building and publishing:

### Security Features

- Hadolint - Dockerfile linting
- Trivy - Vulnerability scanning
- Checkov - Security best practices
- Copacetic - SCA scanning
- Cosign - Container signing
- SBOM - Software Bill of Materials

### Verification

```bash
# Verify container signature
cosign verify ghcr.io/[owner]/[repo]:main

# Download and verify SBOMs
cosign download sbom ghcr.io/[owner]/[repo]:main

# Verify SPDX SBOM
cosign verify-blob --signature code-sbom.spdx.sig code-sbom.spdx.json

# Verify CycloneDX SBOMs
cosign verify-blob --signature code-sbom.cdx.sig code-sbom.cdx.json
cosign verify-blob --signature container-sbom.cdx.sig container-sbom.cdx.json
```

### SBOM Information

The project generates and signs three types of SBOMs:
- Code SBOM (SPDX): Covers all source code dependencies
- Code SBOM (CycloneDX): Detailed dependency graph with licenses
- Container SBOM (CycloneDX): Covers container image contents

All SBOMs are:
- Generated in both SPDX and CycloneDX formats
- Signed using Cosign
- Stored with the container image
- Available as build artifacts

### SBOM Analysis

The CycloneDX format provides:
- Detailed dependency information
- License compliance data
- Component relationships
- Vulnerability correlation

### Build Features

- Multi-architecture support (AMD64/ARM64)
- GitHub Container Registry publishing
- Automated builds on push/PR
- Layer caching for faster builds
- Metadata tagging

### Container Tags

Images are tagged with:
- Branch name for branch builds
- PR number for pull requests
- Semantic version for releases
- Git SHA for precise tracking

### Usage

```bash
# Pull the latest image
docker pull ghcr.io/[owner]/[repo]:main

# Pull for specific architecture
docker pull ghcr.io/[owner]/[repo]:main --platform linux/amd64
docker pull ghcr.io/[owner]/[repo]:main --platform linux/arm64
```

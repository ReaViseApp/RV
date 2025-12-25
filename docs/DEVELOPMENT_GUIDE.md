# ReaVise Development Guide

This guide provides detailed instructions for developers working on the ReaVise platform.

## Table of Contents

1. [Development Environment Setup](#development-environment-setup)
2. [Project Structure](#project-structure)
3. [Development Workflow](#development-workflow)
4. [Code Style Guidelines](#code-style-guidelines)
5. [Testing](#testing)
6. [API Development](#api-development)
7. [Frontend Development](#frontend-development)
8. [Database Management](#database-management)
9. [Deployment](#deployment)
10. [Troubleshooting](#troubleshooting)

## Development Environment Setup

### Prerequisites

Ensure you have the following installed:
- **Node.js**: v18.0.0 or higher
- **Go**: v1.21 or higher
- **MongoDB**: v7.0 or higher
- **Docker** (optional but recommended)
- **Git**: Latest version

### Initial Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ReaViseApp/RV.git
   cd RV
   ```

2. **Start MongoDB:**
   ```bash
   docker-compose up -d
   ```

3. **Set up backend:**
   ```bash
   cd backend
   cp .env.example .env
   # Edit .env with your configuration
   go mod download
   go run cmd/main.go
   ```

4. **Set up frontend:**
   ```bash
   cd frontend
   cp .env.example .env
   # Edit .env with your configuration
   npm install
   npm run dev
   ```

### Environment Variables

#### Backend (.env)
```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=reavise
JWT_SECRET=your-super-secret-jwt-key-change-this
STRIPE_SECRET_KEY=sk_test_your_key
STRIPE_WEBHOOK_SECRET=whsec_your_secret
PAYPAL_CLIENT_ID=your_paypal_client_id
PAYPAL_CLIENT_SECRET=your_paypal_secret
ENVIRONMENT=development
```

#### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
VITE_STRIPE_PUBLIC_KEY=pk_test_your_key
VITE_PAYPAL_CLIENT_ID=your_paypal_client_id
```

## Project Structure

### Backend Structure

```
backend/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── database/
│   │   └── database.go         # MongoDB connection and helpers
│   ├── handlers/
│   │   ├── auth.go            # Authentication endpoints
│   │   ├── posts.go           # Post management
│   │   ├── users.go           # User management
│   │   ├── messages.go        # Messaging system
│   │   ├── comments.go        # Comment system
│   │   ├── transactions.go    # E-commerce transactions
│   │   └── nft.go             # NFT marketplace
│   ├── middleware/
│   │   ├── auth.go            # JWT authentication
│   │   └── error.go           # Error handling and logging
│   ├── models/
│   │   └── models.go          # Data structures
│   └── services/
│       ├── auth.go            # Authentication logic
│       ├── payment.go         # Payment processing
│       └── recommendation.go   # Content recommendation
├── go.mod
├── go.sum
└── .env.example
```

### Frontend Structure

```
frontend/
├── src/
│   ├── components/            # Reusable components
│   │   ├── Navigation.svelte
│   │   ├── PostCard.svelte
│   │   ├── UserCard.svelte
│   │   ├── RubiksCylinder.svelte
│   │   └── RubiksCube.svelte
│   ├── pages/                 # Page components
│   │   ├── MyRV.svelte
│   │   ├── Profile.svelte
│   │   ├── Cubeativity.svelte
│   │   ├── Upload.svelte
│   │   ├── Curated.svelte
│   │   ├── Likes.svelte
│   │   ├── Exchange.svelte
│   │   ├── Settings.svelte
│   │   ├── Checkout.svelte
│   │   └── PostDetails.svelte
│   ├── services/              # API and external services
│   │   ├── api.ts
│   │   └── auth.ts
│   ├── stores/                # Svelte stores
│   │   └── index.ts
│   ├── styles/                # Global styles
│   │   └── global.css
│   ├── types/                 # TypeScript type definitions
│   │   └── index.ts
│   ├── App.svelte            # Root component
│   └── main.ts               # Entry point
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── .env.example
```

## Development Workflow

### Git Workflow

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes and commit:**
   ```bash
   git add .
   git commit -m "Add feature: description"
   ```

3. **Push to remote:**
   ```bash
   git push origin feature/your-feature-name
   ```

4. **Create a Pull Request** on GitHub

### Branch Naming Convention

- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Adding tests
- `chore/` - Maintenance tasks

## Code Style Guidelines

### Go (Backend)

- Follow official Go conventions
- Use `gofmt` for formatting
- Follow effective Go guidelines
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions small and focused

Example:
```go
// GetUser retrieves a user by ID from the database
func (h *UserHandler) GetUser(c *gin.Context) {
    userID := c.Param("id")
    // Implementation...
}
```

### TypeScript/Svelte (Frontend)

- Use TypeScript for type safety
- Follow Svelte best practices
- Use meaningful component and variable names
- Keep components small and focused
- Use Svelte stores for global state
- Follow reactive programming patterns

Example:
```typescript
interface User {
    id: string;
    username: string;
    email: string;
}

export let user: User;
```

## Testing

### Backend Testing

```bash
cd backend
go test ./...
```

Create tests in `*_test.go` files:

```go
func TestGetUser(t *testing.T) {
    // Test implementation
}
```

### Frontend Testing

```bash
cd frontend
npm run check  # TypeScript type checking
```

## API Development

### Adding New Endpoints

1. **Define the model** in `internal/models/models.go`
2. **Create handler** in `internal/handlers/`
3. **Add routes** in `cmd/main.go`
4. **Test the endpoint** using curl or Postman

Example:
```go
// In handlers/example.go
func (h *ExampleHandler) GetExample(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "example"})
}

// In cmd/main.go
api.GET("/example", exampleHandler.GetExample)
```

### API Response Format

Success response:
```json
{
    "data": { /* response data */ }
}
```

Error response:
```json
{
    "error": "Error message"
}
```

## Frontend Development

### Creating New Components

1. Create file in `src/components/`
2. Define props and types
3. Implement component logic
4. Export the component

Example:
```svelte
<script lang="ts">
    export let title: string;
    export let count: number = 0;
</script>

<div class="component">
    <h2>{title}</h2>
    <p>Count: {count}</p>
</div>

<style>
    .component {
        padding: 20px;
    }
</style>
```

### Using Babylon.js

For 3D components:
```typescript
import { Engine, Scene, ArcRotateCamera } from '@babylonjs/core';

let engine: Engine;
let scene: Scene;

onMount(() => {
    engine = new Engine(canvas, true);
    scene = new Scene(engine);
    // Setup scene...
});
```

## Database Management

### MongoDB Collections

- **users**: User accounts
- **posts**: User posts
- **comments**: Post comments
- **messages**: Direct messages
- **transactions**: E-commerce transactions
- **nft_listings**: NFT auctions
- **likes**: Post likes
- **follows**: Follow relationships

### Accessing MongoDB

**Via Mongo Express** (GUI):
```
http://localhost:8081
```

**Via Mongo Shell**:
```bash
docker exec -it reavise-mongodb mongosh
use reavise
db.users.find()
```

### Creating Indexes

Add indexes in database initialization:
```go
_, err := db.Users().Indexes().CreateOne(ctx, mongo.IndexModel{
    Keys: bson.D{{Key: "email", Value: 1}},
    Options: options.Index().SetUnique(true),
})
```

## Deployment

### Production Build

**Backend:**
```bash
cd backend
go build -o reavise-backend cmd/main.go
./reavise-backend
```

**Frontend:**
```bash
cd frontend
npm run build
# Output in dist/ directory
```

### Environment Configuration

- Use production database URI
- Enable HTTPS
- Set secure JWT secret
- Configure CORS for production domain
- Set up proper logging
- Enable rate limiting

### Deployment Checklist

- [ ] Set all environment variables
- [ ] Build frontend production bundle
- [ ] Build backend binary
- [ ] Set up reverse proxy (nginx)
- [ ] Configure SSL certificates
- [ ] Set up database backups
- [ ] Configure monitoring
- [ ] Set up error tracking
- [ ] Test all critical paths
- [ ] Document deployment process

## Troubleshooting

### Common Issues

**Port already in use:**
```bash
# Find process using port
lsof -i :8080
# Kill process
kill -9 <PID>
```

**MongoDB connection failed:**
```bash
# Check if MongoDB is running
docker ps
# Restart MongoDB
docker-compose restart mongodb
```

**Frontend build fails:**
```bash
# Clear node_modules and reinstall
rm -rf node_modules package-lock.json
npm install
```

**Backend won't start:**
```bash
# Check Go modules
go mod tidy
go mod download
```

### Debug Mode

**Backend:**
```bash
# Run with verbose logging
GIN_MODE=debug go run cmd/main.go
```

**Frontend:**
```bash
# Run dev server with verbose output
npm run dev -- --debug
```

### Logs

**Backend logs:**
- Check console output
- Logs include request/response info
- Error stack traces included

**Frontend logs:**
- Browser console (F12)
- Network tab for API calls
- Vue devtools (if installed)

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

### Pull Request Guidelines

- Clear description of changes
- Link related issues
- Add tests if applicable
- Update documentation
- Follow code style guidelines
- Ensure all tests pass

## Support

For questions or issues:
- Email: dev@reavise.com
- GitHub Issues: [Create an issue](https://github.com/ReaViseApp/RV/issues)

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Svelte Documentation](https://svelte.dev/docs)
- [Babylon.js Documentation](https://doc.babylonjs.com/)
- [MongoDB Documentation](https://docs.mongodb.com/)
- [Gin Framework](https://gin-gonic.com/docs/)

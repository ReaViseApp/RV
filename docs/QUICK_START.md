# ReaVise Quick Start Guide

Get the ReaVise platform up and running in under 5 minutes!

## Prerequisites Checklist

Before starting, ensure you have:
- [ ] Node.js v18+ installed (`node --version`)
- [ ] Go v1.21+ installed (`go version`)
- [ ] Docker and Docker Compose installed (`docker --version`)
- [ ] Git installed (`git --version`)

## Quick Start Steps

### 1. Clone and Navigate

```bash
git clone https://github.com/ReaViseApp/RV.git
cd RV
```

### 2. Start MongoDB

```bash
docker-compose up -d
```

This starts:
- MongoDB on `localhost:27017`
- Mongo Express (database GUI) on `localhost:8081`

**Verify:** Visit http://localhost:8081 to see the database interface.

### 3. Start Backend

```bash
cd backend

# Copy environment file
cp .env.example .env

# Download dependencies
go mod download

# Start the server
go run cmd/main.go
```

**Verify:** You should see:
```
Connected to MongoDB!
Server starting on port 8080
```

Visit http://localhost:8080/health - you should see `{"status":"ok"}`

### 4. Start Frontend (in a new terminal)

```bash
cd frontend

# Copy environment file
cp .env.example .env

# Install dependencies (first time only)
npm install

# Start development server
npm run dev
```

**Verify:** You should see:
```
VITE v5.x.x  ready in xxx ms

➜  Local:   http://localhost:5173/
➜  Network: use --host to expose
```

### 5. Open Application

Visit: **http://localhost:5173**

You should see the ReaVise home page with the 3D Rubik's Cylinder!

## First Steps in the Application

### Create Your Account

1. Click on any protected feature (or navigate to `/upload`)
2. You'll need to register first
3. Use the API directly or implement a registration UI

**Using curl to register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

Save the token from the response!

### Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Create a Post

```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "media": [{
      "url": "https://via.placeholder.com/600",
      "type": "image",
      "category": "design"
    }],
    "description": "My first ReaVise post!",
    "category": "design",
    "hashtags": ["test", "firstpost"]
  }'
```

## Explore the Platform

### Navigation (Left Sidebar)

1. **MyRV** 🏠 - Home feed with 3D cylinder
2. **Curated** 🛒 - Shopping cart
3. **Cubeativity** 🔍 - Search with 3D cube
4. **Upload** ➕ - Create new posts
5. **Likes** ❤️ - Your liked content
6. **Exchange** 💬 - Messages
7. **Profile** 👤 - Your profile
8. **Settings** ⚙️ - Account settings

### 3D Navigation Controls

**MyRV & Profile (Rubik's Cylinder):**
- **Rotate entire cylinder**: Scroll or two-finger swipe
- **Rotate single row**: Click on row + drag
- **Camera orbit**: Click and drag on background

**Cubeativity (Rubik's Cube):**
- **Rotate entire cube**: Scroll or two-finger swipe
- **Rotate single column**: Click on column + drag up/down
- **Zoom**: Cmd+Up (enlarge) / Cmd+Down (shrink)
- **Camera orbit**: Click and drag on background

### Content Categories

- **The Lot** (Pink): Materials needing customization
- **Design** (Blue): Creative ideas and proposals
- **ReaVise** (Green): Finished products with reviews

## Common Tasks

### View Database

1. Open http://localhost:8081
2. Click on "reavise" database
3. Browse collections (users, posts, etc.)

### Stop Services

```bash
# Stop frontend
Ctrl+C in frontend terminal

# Stop backend
Ctrl+C in backend terminal

# Stop MongoDB
docker-compose down
```

### Restart Services

```bash
# Restart MongoDB
docker-compose up -d

# Restart backend
cd backend && go run cmd/main.go

# Restart frontend
cd frontend && npm run dev
```

## Troubleshooting

### Port Already in Use

**Backend (8080):**
```bash
lsof -i :8080
kill -9 <PID>
```

**Frontend (5173):**
```bash
lsof -i :5173
kill -9 <PID>
```

### MongoDB Won't Start

```bash
# Check Docker
docker ps

# View logs
docker-compose logs mongodb

# Restart
docker-compose restart mongodb
```

### Backend Dependencies Issue

```bash
cd backend
go mod tidy
go mod download
```

### Frontend Dependencies Issue

```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
```

### "Cannot connect to database"

1. Ensure MongoDB is running: `docker ps`
2. Check MongoDB URI in `backend/.env`
3. Try: `docker-compose restart mongodb`

## Default Configuration

### Ports
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- MongoDB: localhost:27017
- Mongo Express: http://localhost:8081

### Database
- Name: `reavise`
- Collections: users, posts, comments, messages, transactions, nft_listings, likes, follows

### API Base URL
- Local: http://localhost:8080/api
- Health Check: http://localhost:8080/health

## Next Steps

1. **Read the docs:**
   - [Development Guide](./docs/DEVELOPMENT_GUIDE.md)
   - [API Reference](./docs/API_REFERENCE.md)

2. **Explore features:**
   - Create posts in different categories
   - Test the 3D navigation
   - Try the messaging system
   - Create NFT listings

3. **Customize:**
   - Update environment variables
   - Modify styles in `frontend/src/styles/global.css`
   - Add your own components

4. **Deploy:**
   - Build for production
   - Set up hosting
   - Configure domain

## Getting Help

- **Documentation**: Check `/docs` folder
- **Issues**: GitHub Issues
- **Email**: support@reavise.com

## Quick Commands Reference

```bash
# Start everything
docker-compose up -d                    # MongoDB
cd backend && go run cmd/main.go       # Backend (new terminal)
cd frontend && npm run dev             # Frontend (new terminal)

# Build
cd backend && go build -o reavise-backend cmd/main.go
cd frontend && npm run build

# Test
cd backend && go test ./...
cd frontend && npm run check

# Stop
docker-compose down                     # MongoDB
Ctrl+C                                 # Backend & Frontend
```

## Success Indicators

✅ MongoDB running: `docker ps` shows `reavise-mongodb`
✅ Backend running: http://localhost:8080/health returns `{"status":"ok"}`
✅ Frontend running: http://localhost:5173 loads the application
✅ Database accessible: http://localhost:8081 shows Mongo Express

---

**Enjoy building with ReaVise!** 🚀

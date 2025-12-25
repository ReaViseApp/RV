# ReaVise Platform

![ReaVise Logo](https://via.placeholder.com/150x50?text=ReaVise)

**ReaVise** is a comprehensive creative content sharing platform that combines social media, e-commerce, and NFT auction functionality. It enables users to search and match their creative demands with design proposals for deeply personalized transactions.

## 🎯 Overview

ReaVise connects creators, designers, and consumers through three content categories:
- **The Lot**: User-collected materials and used items needing customization
- **Design**: Creative ideas, illustrations, AI-generated proposals, and NFT auction items
- **ReaVise**: Finished products and verified reviews showcasing completed work

## ✨ Key Features

### Interactive 3D Navigation
- **MyRV (Home)**: 3D Rubik's Cylinder with three rotating rows (The Lot, Design, ReaVise)
- **Cubeativity (Search)**: Interactive 3×3 Rubik's Cube for exploring categorized content
- Mouse and touch gesture controls for intuitive navigation

### Social Features
- User profiles with followers/following
- Post creation with images and videos
- Comments and likes
- Direct messaging (Exchange)
- Hashtag support

### E-Commerce System
- Shopping cart (Curated page)
- Secure checkout with Stripe/PayPal
- Transaction management
- Category-specific purchasing rules
- Review system

### NFT Marketplace
- NFT auctions for digital art
- Smart contract integration
- Bid management
- Blockchain verification

### Business Features
- Verified business accounts
- Multi-category content labeling
- Enhanced visibility
- Professional tools

## 🏗️ Architecture

### Frontend
- **Framework**: Svelte with TypeScript
- **Build Tool**: Vite
- **3D Rendering**: Babylon.js
- **Routing**: svelte-routing
- **HTTP Client**: Axios

### Backend
- **Language**: Go (Golang)
- **Framework**: Gin
- **Authentication**: JWT
- **Database Driver**: MongoDB driver

### Database
- **Database**: MongoDB
- **Collections**: users, posts, comments, messages, transactions, nft_listings, likes, follows

### Payment & Blockchain
- **Payment Processing**: Stripe, PayPal
- **NFT Support**: Smart contract integration (extensible)

## 📋 Prerequisites

- Node.js (v18 or higher)
- Go (v1.21 or higher)
- MongoDB (v7.0 or higher)
- Docker and Docker Compose (optional, for containerized MongoDB)

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ReaViseApp/RV.git
cd RV
```

### 2. Set Up MongoDB

#### Option A: Using Docker Compose (Recommended)

```bash
docker-compose up -d
```

This starts MongoDB on `localhost:27017` and Mongo Express (GUI) on `localhost:8081`.

#### Option B: Local MongoDB Installation

Install MongoDB locally and ensure it's running on `localhost:27017`.

### 3. Set Up Backend

```bash
cd backend

# Copy environment file
cp .env.example .env

# Edit .env with your configuration
nano .env

# Install dependencies
go mod download

# Run the backend
go run cmd/main.go
```

The backend will start on `http://localhost:8080`.

### 4. Set Up Frontend

```bash
cd frontend

# Copy environment file
cp .env.example .env

# Edit .env with your configuration
nano .env

# Install dependencies
npm install

# Run the development server
npm run dev
```

The frontend will start on `http://localhost:5173`.

## 🔧 Configuration

### Backend Environment Variables

```env
PORT=8080
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=reavise
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
STRIPE_SECRET_KEY=your_stripe_secret_key_here
STRIPE_WEBHOOK_SECRET=your_stripe_webhook_secret_here
PAYPAL_CLIENT_ID=your_paypal_client_id_here
PAYPAL_CLIENT_SECRET=your_paypal_client_secret_here
ENVIRONMENT=development
```

### Frontend Environment Variables

```env
VITE_API_URL=http://localhost:8080/api
VITE_STRIPE_PUBLIC_KEY=your_stripe_public_key_here
VITE_PAYPAL_CLIENT_ID=your_paypal_client_id_here
```

## 📱 Features by Page

### MyRV (Home Feed)
- Binary filter: "Following" or "RVed for You"
- 3D Rubik's Cylinder with three content rows
- Surface rotation: Mouse + two-finger swipe or scroll
- Row rotation: Mouse on row + one-finger drag
- Post interactions: like, comment, curate

### Profile
- User information display
- Follow/unfollow functionality
- Direct messaging
- User's content organized by category
- 3D cylinder showing user's posts

### Cubeativity (Search)
- Search bar for accounts and hashtags
- Interactive 3×3 Rubik's Cube
- Size adjustment: Cmd+Up/Down or gestures
- Column rotation for category browsing
- Surface rotation for different views

### Upload
- Multi-file upload support
- Per-file category labeling
- Description and hashtag input
- Business account features

### Curated (Shopping Cart)
- View all curated items
- Remove items from cart
- Proceed to checkout

### Checkout
- Contact information form
- Shipping address
- Payment method selection (Stripe/PayPal)
- Legal agreement acceptance
- Order summary

### Exchange (Messaging)
- Conversation list
- Real-time messaging
- Message history
- Unread indicators

### Likes
- View all liked posts
- Quick access to favorites

### Settings
- Profile editing
- Privacy settings
- Data export
- Account deletion
- Legal document access

## 🎨 3D Navigation Controls

### Rubik's Cylinder (MyRV, Profile)
- **Full Rotation**: Mouse on cylinder + two-finger swipe left/right OR scroll
- **Single Row**: Mouse on specific row + one-finger drag left/right
- **Camera**: Click and drag to orbit view

### Rubik's Cube (Cubeativity)
- **Size**: Cmd+Up to enlarge, Cmd+Down to shrink
- **Surface Rotation**: Mouse on cube + two-finger swipe OR scroll
- **Column Rotation**: Mouse on column + one-finger drag up/down
- **Camera**: Click and drag to orbit view

## 🔒 Security & Legal

### Security Features
- JWT-based authentication
- Password hashing with bcrypt
- HTTPS enforcement (production)
- CORS configuration
- Input validation
- SQL injection prevention
- XSS protection

### Legal Compliance
- **Privacy Policy**: GDPR compliant
- **Terms of Service**: Comprehensive user agreement
- **Data Processing Agreement**: Data handling terms
- **Cookie Policy**: Cookie usage disclosure
- **E-Commerce Terms**: Purchase terms and conditions
- **NFT Terms**: NFT-specific legal framework
- **Copyright Framework**: IP protection and DMCA compliance

All legal documents are available in the `/legal` directory and linked in the application.

## 🛠️ Development

### Backend Structure

```
backend/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── database/
│   │   └── database.go      # MongoDB connection
│   ├── handlers/
│   │   ├── auth.go          # Authentication endpoints
│   │   └── posts.go         # Post management endpoints
│   ├── middleware/
│   │   ├── auth.go          # JWT middleware
│   │   └── error.go         # Error handling
│   ├── models/
│   │   └── models.go        # Data structures
│   └── services/
│       └── auth.go          # Authentication service
├── go.mod
└── .env.example
```

### Frontend Structure

```
frontend/
├── src/
│   ├── components/
│   │   ├── Navigation.svelte       # Sidebar navigation
│   │   ├── PostCard.svelte        # Post display
│   │   ├── UserCard.svelte        # User info card
│   │   ├── RubiksCylinder.svelte  # 3D cylinder
│   │   └── RubiksCube.svelte      # 3D cube
│   ├── pages/
│   │   ├── MyRV.svelte            # Home feed
│   │   ├── Profile.svelte         # User profile
│   │   ├── Cubeativity.svelte     # Search page
│   │   ├── Upload.svelte          # Content upload
│   │   ├── Curated.svelte         # Shopping cart
│   │   ├── Likes.svelte           # Liked posts
│   │   ├── Exchange.svelte        # Messaging
│   │   ├── Settings.svelte        # User settings
│   │   ├── Checkout.svelte        # Payment page
│   │   └── PostDetails.svelte     # Post view
│   ├── services/
│   │   ├── api.ts                 # API client
│   │   └── auth.ts                # Auth service
│   ├── stores/
│   │   └── index.ts               # Svelte stores
│   ├── styles/
│   │   └── global.css             # Global styles
│   ├── types/
│   │   └── index.ts               # TypeScript types
│   ├── App.svelte                 # Root component
│   └── main.ts                    # Entry point
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── .env.example
```

### Building for Production

#### Backend

```bash
cd backend
go build -o reavise-backend cmd/main.go
./reavise-backend
```

#### Frontend

```bash
cd frontend
npm run build
# Output in dist/ directory
```

## 🧪 Testing

```bash
# Frontend
cd frontend
npm run check  # TypeScript checking

# Backend
cd backend
go test ./...
```

## 📊 API Documentation

### Authentication

#### POST /api/auth/register
Register a new user.

**Request:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "user": { /* user object */ },
  "token": "jwt_token_here"
}
```

#### POST /api/auth/login
Login existing user.

#### GET /api/auth/me
Get current user (requires authentication).

### Posts

#### GET /api/posts
Get posts (with optional filters).

**Query Parameters:**
- `category`: Filter by category (lot, design, reavise)
- `userId`: Filter by user ID

#### POST /api/posts
Create a new post (requires authentication).

#### GET /api/posts/:id
Get specific post.

#### POST /api/posts/:id/like
Like a post (requires authentication).

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is proprietary software. All rights reserved.

## 🔗 Links

- **Website**: [https://reavise.com](https://reavise.com) (Coming soon)
- **Documentation**: [https://docs.reavise.com](https://docs.reavise.com) (Coming soon)
- **Support**: support@reavise.com

## 📞 Contact

- **General Inquiries**: info@reavise.com
- **Support**: support@reavise.com
- **Legal**: legal@reavise.com
- **Privacy**: privacy@reavise.com
- **Copyright/DMCA**: copyright@reavise.com

## 🙏 Acknowledgments

- Babylon.js for 3D rendering capabilities
- Svelte team for the amazing framework
- Gin framework for Go
- MongoDB for flexible data storage
- Stripe and PayPal for payment processing

---

**Built with ❤️ by the ReaVise Team**

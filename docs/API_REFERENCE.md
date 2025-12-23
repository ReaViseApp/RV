# ReaVise Platform - API Reference

## Base URL
```
Development: http://localhost:8080/api
Production: https://api.reavise.com
```

## Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

---

## Authentication Endpoints

### POST /auth/register
Register a new user account.

**Request:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:** `201 Created`
```json
{
  "user": {
    "id": "...",
    "username": "john_doe",
    "email": "john@example.com",
    "followersCount": 0,
    "followingCount": 0,
    "isBusinessAccount": false,
    "isVerified": false,
    "createdAt": "2024-12-23T..."
  },
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### POST /auth/login
Login with existing credentials.

**Request:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:** `200 OK`
```json
{
  "user": { /* user object */ },
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### GET /auth/me
Get current authenticated user. **[Protected]**

**Response:** `200 OK`
```json
{
  "id": "...",
  "username": "john_doe",
  "email": "john@example.com",
  ...
}
```

---

## User Endpoints

### GET /users/:id
Get user profile by ID.

**Response:** `200 OK`
```json
{
  "id": "...",
  "username": "john_doe",
  "profilePhoto": "https://...",
  "bio": "Creative designer",
  "followersCount": 150,
  "followingCount": 200,
  ...
}
```

### PUT /users/:id
Update user profile. **[Protected]** (own profile only)

**Request:**
```json
{
  "username": "john_doe_updated",
  "bio": "Updated bio",
  "website": "https://johndoe.com",
  "location": "New York",
  "profilePhoto": "https://..."
}
```

**Response:** `200 OK`

### POST /users/:id/follow
Follow a user. **[Protected]**

**Response:** `200 OK`

### DELETE /users/:id/follow
Unfollow a user. **[Protected]**

**Response:** `200 OK`

### DELETE /users/:id
Delete user account. **[Protected]** (own account only)

**Response:** `200 OK`

---

## Post Endpoints

### GET /posts
Get posts with optional filters.

**Query Parameters:**
- `category` (optional): Filter by category (lot, design, reavise)
- `userId` (optional): Filter by user ID

**Response:** `200 OK`
```json
[
  {
    "id": "...",
    "userId": "...",
    "username": "john_doe",
    "media": [
      {
        "url": "https://...",
        "type": "image",
        "category": "design"
      }
    ],
    "description": "My latest creation",
    "category": "design",
    "hashtags": ["art", "design"],
    "likesCount": 42,
    "commentsCount": 10,
    "createdAt": "2024-12-23T..."
  }
]
```

### POST /posts
Create a new post. **[Protected]**

**Request:**
```json
{
  "media": [
    {
      "url": "https://...",
      "type": "image",
      "category": "design"
    }
  ],
  "description": "My latest creation",
  "category": "design",
  "hashtags": ["art", "design"]
}
```

**Response:** `201 Created`

### GET /posts/:id
Get specific post by ID.

**Response:** `200 OK`

### POST /posts/:id/like
Like a post. **[Protected]**

**Response:** `200 OK`

---

## Comment Endpoints

### GET /comments/post/:postId
Get comments for a post.

**Response:** `200 OK`
```json
[
  {
    "id": "...",
    "postId": "...",
    "userId": "...",
    "username": "jane_doe",
    "userAvatar": "https://...",
    "text": "Great work!",
    "createdAt": "2024-12-23T..."
  }
]
```

### POST /comments/post/:postId
Create a comment. **[Protected]**

**Request:**
```json
{
  "text": "Great work!"
}
```

**Response:** `201 Created`

### DELETE /comments/:id
Delete a comment. **[Protected]** (own comment only)

**Response:** `200 OK`

---

## Message Endpoints

### POST /messages
Send a message. **[Protected]**

**Request:**
```json
{
  "receiverId": "...",
  "text": "Hello!"
}
```

**Response:** `201 Created`

### GET /messages/conversations
Get all conversations. **[Protected]**

**Response:** `200 OK`

### GET /messages/:userId
Get messages with a specific user. **[Protected]**

**Response:** `200 OK`
```json
[
  {
    "id": "...",
    "senderId": "...",
    "receiverId": "...",
    "text": "Hello!",
    "isRead": true,
    "createdAt": "2024-12-23T..."
  }
]
```

---

## Transaction Endpoints

### POST /transactions
Create a transaction. **[Protected]**

**Request:**
```json
{
  "postId": "...",
  "amount": 99.99,
  "paymentMethod": "stripe"
}
```

**Response:** `201 Created`

### GET /transactions
Get user's transactions. **[Protected]**

**Response:** `200 OK`

### GET /transactions/:id
Get specific transaction. **[Protected]**

**Response:** `200 OK`

### PUT /transactions/:id/status
Update transaction status. **[Protected]** (seller only)

**Request:**
```json
{
  "status": "completed"
}
```

**Response:** `200 OK`

---

## NFT Endpoints

### GET /nft
Get NFT listings.

**Query Parameters:**
- `status` (optional): Filter by status (active, sold, expired)

**Response:** `200 OK`

### POST /nft
Create NFT listing. **[Protected]**

**Request:**
```json
{
  "postId": "...",
  "startingBid": 0.5,
  "auctionEndDate": "2024-12-31T23:59:59Z"
}
```

**Response:** `201 Created`

### GET /nft/:id
Get specific NFT listing.

**Response:** `200 OK`

### POST /nft/:id/bid
Place bid on NFT. **[Protected]**

**Request:**
```json
{
  "bidAmount": 1.0
}
```

**Response:** `200 OK`

### POST /nft/:id/complete
Complete NFT auction. **[Protected]** (owner only)

**Response:** `200 OK`

---

## Recommendation Endpoints

### GET /recommendations/foryou
Get recommended posts. **[Protected]**

**Response:** `200 OK`
```json
[/* array of posts */]
```

### GET /recommendations/following
Get posts from followed users. **[Protected]**

**Response:** `200 OK`
```json
[/* array of posts */]
```

---

## Payment Endpoints

### POST /payment/create-intent
Create Stripe payment intent. **[Protected]**

**Request:**
```json
{
  "amount": 99.99,
  "currency": "usd"
}
```

**Response:** `200 OK`
```json
{
  "id": "pi_...",
  "amount": 9999,
  "currency": "usd",
  "status": "requires_payment_method",
  "clientSecret": "pi_...secret..."
}
```

---

## Error Responses

All error responses follow this format:

**Response:** `4xx` or `5xx`
```json
{
  "error": "Error message description"
}
```

### Common Error Codes

- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource already exists
- `500 Internal Server Error` - Server error

---

## Rate Limiting

*To be implemented*

- Rate limits will be enforced per endpoint
- Headers will indicate current usage and limits
- Exceeding limits results in `429 Too Many Requests`

---

## Pagination

*To be implemented for list endpoints*

Standard pagination parameters:
- `page`: Page number (default: 1)
- `perPage`: Items per page (default: 20, max: 100)

Response includes:
```json
{
  "data": [/* items */],
  "pagination": {
    "page": 1,
    "perPage": 20,
    "total": 100,
    "totalPages": 5
  }
}
```

---

## WebSocket Events

*To be implemented*

Real-time features will use WebSocket connections for:
- New messages
- Notifications
- Live auction updates
- Post interactions

---

## Versioning

API version is included in the base URL path:
- Current: `/api/v1/...`
- Future versions: `/api/v2/...`

---

For questions or support, contact: api@reavise.com

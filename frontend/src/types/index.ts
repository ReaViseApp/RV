export interface User {
  id: string;
  username: string;
  email: string;
  profilePhoto?: string;
  bio?: string;
  website?: string;
  location?: string;
  followersCount: number;
  followingCount: number;
  isBusinessAccount: boolean;
  isVerified: boolean;
  createdAt: string;
}

export interface Post {
  id: string;
  userId: string;
  username: string;
  userAvatar?: string;
  userLocation?: string;
  media: MediaItem[];
  description: string;
  category: 'lot' | 'design' | 'reavise';
  hashtags: string[];
  likesCount: number;
  commentsCount: number;
  isLiked: boolean;
  createdAt: string;
}

export interface MediaItem {
  url: string;
  type: 'image' | 'video';
  category?: 'lot' | 'design' | 'reavise';
}

export interface Comment {
  id: string;
  postId: string;
  userId: string;
  username: string;
  userAvatar?: string;
  text: string;
  createdAt: string;
}

export interface Message {
  id: string;
  senderId: string;
  receiverId: string;
  text: string;
  isRead: boolean;
  createdAt: string;
}

export interface Transaction {
  id: string;
  buyerId: string;
  sellerId: string;
  postId: string;
  amount: number;
  status: 'pending' | 'completed' | 'cancelled';
  paymentMethod: 'stripe' | 'paypal';
  createdAt: string;
}

export interface NFTListing {
  id: string;
  postId: string;
  ownerId: string;
  startingBid: number;
  currentBid: number;
  auctionEndDate: string;
  status: 'active' | 'sold' | 'expired';
  createdAt: string;
}

export interface CartItem {
  postId: string;
  post: Post;
  addedAt: string;
}

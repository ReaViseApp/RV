import { writable } from 'svelte/store';
import type { User, Post, CartItem } from '../types';

// Auth store
export const currentUser = writable<User | null>(null);
export const isAuthenticated = writable<boolean>(false);
export const authToken = writable<string | null>(null);

// Feed mode store
export const feedMode = writable<'following' | 'foryou'>('foryou');

// Cart store
export const cart = writable<CartItem[]>([]);

// Posts store
export const posts = writable<Post[]>([]);

// Notifications store
export const notifications = writable<any[]>([]);

// Loading state
export const isLoading = writable<boolean>(false);

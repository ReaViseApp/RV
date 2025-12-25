<script lang="ts">
  import { onMount } from 'svelte';
  import { cart } from '../stores';
  import PostCard from '../components/PostCard.svelte';
  import type { CartItem } from '../types';

  let cartItems: CartItem[] = [];

  onMount(() => {
    cart.subscribe(items => {
      cartItems = items;
    });
  });

  function removeFromCart(postId: string) {
    cart.update(items => items.filter(item => item.postId !== postId));
  }

  function proceedToCheckout() {
    if (cartItems.length === 0) {
      alert('Your cart is empty');
      return;
    }
    window.location.href = '/checkout';
  }

  $: totalItems = cartItems.length;
</script>

<div class="curated-page">
  <div class="header">
    <h1>Curated</h1>
    <p class="subtitle">Your shopping cart</p>
  </div>

  {#if cartItems.length > 0}
    <div class="cart-items">
      {#each cartItems as item}
        <div class="cart-item">
          <PostCard post={item.post} />
          <button class="remove-btn" on:click={() => removeFromCart(item.postId)}>
            Remove from Cart
          </button>
        </div>
      {/each}
    </div>

    <div class="cart-summary">
      <div class="summary-card">
        <h2>Order Summary</h2>
        <div class="summary-row">
          <span>Total Items:</span>
          <strong>{totalItems}</strong>
        </div>
        <p class="note">Prices will be confirmed during checkout</p>
        <button class="btn btn-primary checkout-btn" on:click={proceedToCheckout}>
          Proceed to Checkout
        </button>
      </div>
    </div>
  {:else}
    <div class="empty-cart">
      <p class="empty-icon">🛒</p>
      <h2>Your cart is empty</h2>
      <p>Browse posts and click "Curate" to add items to your cart</p>
      <a href="/" class="btn btn-primary">Browse Posts</a>
    </div>
  {/if}
</div>

<style>
  .curated-page {
    max-width: 1200px;
    margin: 0 auto;
  }

  .header {
    margin-bottom: 30px;
    padding: 20px 0;
    border-bottom: 1px solid var(--border-color);
  }

  h1 {
    font-size: 32px;
    font-weight: 700;
    margin-bottom: 5px;
  }

  .subtitle {
    color: var(--text-secondary);
    font-size: 16px;
  }

  .cart-items {
    display: grid;
    gap: 20px;
    margin-bottom: 30px;
  }

  .cart-item {
    position: relative;
  }

  .remove-btn {
    margin: 10px 0 0 0;
    padding: 8px 16px;
    background-color: var(--error-color);
    color: var(--white);
    border-radius: 6px;
    font-weight: 600;
    transition: opacity 0.2s;
  }

  .remove-btn:hover {
    opacity: 0.9;
  }

  .cart-summary {
    position: sticky;
    top: 20px;
  }

  .summary-card {
    background-color: var(--white);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    padding: 24px;
  }

  .summary-card h2 {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .summary-row {
    display: flex;
    justify-content: space-between;
    margin-bottom: 15px;
    padding-bottom: 15px;
    border-bottom: 1px solid var(--border-color);
  }

  .note {
    font-size: 14px;
    color: var(--text-secondary);
    margin: 15px 0;
  }

  .checkout-btn {
    width: 100%;
    padding: 14px;
    font-size: 16px;
  }

  .empty-cart {
    text-align: center;
    padding: 80px 20px;
  }

  .empty-icon {
    font-size: 80px;
    margin-bottom: 20px;
  }

  .empty-cart h2 {
    font-size: 24px;
    margin-bottom: 10px;
  }

  .empty-cart p {
    color: var(--text-secondary);
    margin-bottom: 30px;
  }
</style>

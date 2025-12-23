<script lang="ts">
  import { cart } from '../stores';
  import { onMount } from 'svelte';
  import type { CartItem } from '../types';

  let cartItems: CartItem[] = [];
  let contactInfo = {
    fullName: '',
    email: '',
    phone: '',
    address: '',
    city: '',
    zipCode: '',
    country: ''
  };
  let paymentMethod: 'stripe' | 'paypal' = 'stripe';
  let agreedToTerms: boolean = false;
  let agreedToPrivacy: boolean = false;
  let agreedToEcommerce: boolean = false;

  onMount(() => {
    cart.subscribe(items => {
      cartItems = items;
    });
  });

  async function handleCheckout() {
    if (!agreedToTerms || !agreedToPrivacy || !agreedToEcommerce) {
      alert('Please agree to all terms and policies');
      return;
    }

    if (!contactInfo.fullName || !contactInfo.email) {
      alert('Please fill in all required fields');
      return;
    }

    // TODO: Implement payment processing
    console.log('Processing payment...', { contactInfo, paymentMethod, cartItems });
  }
</script>

<div class="checkout-page">
  <div class="header">
    <h1>Checkout</h1>
  </div>

  <div class="checkout-container">
    <div class="checkout-form">
      <section class="form-section">
        <h2>Contact Information</h2>
        <div class="form-row">
          <div class="form-group">
            <label>Full Name *</label>
            <input type="text" bind:value={contactInfo.fullName} required />
          </div>
          <div class="form-group">
            <label>Email *</label>
            <input type="email" bind:value={contactInfo.email} required />
          </div>
        </div>
        <div class="form-group">
          <label>Phone</label>
          <input type="tel" bind:value={contactInfo.phone} />
        </div>
      </section>

      <section class="form-section">
        <h2>Shipping Address</h2>
        <div class="form-group">
          <label>Address</label>
          <input type="text" bind:value={contactInfo.address} />
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>City</label>
            <input type="text" bind:value={contactInfo.city} />
          </div>
          <div class="form-group">
            <label>ZIP Code</label>
            <input type="text" bind:value={contactInfo.zipCode} />
          </div>
        </div>
        <div class="form-group">
          <label>Country</label>
          <input type="text" bind:value={contactInfo.country} />
        </div>
      </section>

      <section class="form-section">
        <h2>Payment Method</h2>
        <div class="payment-options">
          <label class="payment-option">
            <input type="radio" bind:group={paymentMethod} value="stripe" />
            <span>💳 Credit Card (Stripe)</span>
          </label>
          <label class="payment-option">
            <input type="radio" bind:group={paymentMethod} value="paypal" />
            <span>🅿️ PayPal</span>
          </label>
        </div>
      </section>

      <section class="form-section legal-section">
        <h2>Legal Agreements</h2>
        <div class="checkbox-group">
          <label>
            <input type="checkbox" bind:checked={agreedToTerms} />
            <span>I agree to the <a href="/legal/terms-of-service" target="_blank">Terms of Service</a></span>
          </label>
          <label>
            <input type="checkbox" bind:checked={agreedToPrivacy} />
            <span>I agree to the <a href="/legal/privacy-policy" target="_blank">Privacy Policy</a></span>
          </label>
          <label>
            <input type="checkbox" bind:checked={agreedToEcommerce} />
            <span>I agree to the <a href="/legal/ecommerce-terms" target="_blank">E-Commerce Purchase Terms</a></span>
          </label>
        </div>
      </section>

      <button class="btn btn-primary complete-btn" on:click={handleCheckout}>
        Complete Purchase
      </button>
    </div>

    <div class="order-summary">
      <div class="summary-card">
        <h2>Order Summary</h2>
        {#if cartItems.length > 0}
          <div class="summary-items">
            {#each cartItems as item}
              <div class="summary-item">
                <img src={item.post.media[0]?.url} alt={item.post.description} />
                <div class="item-info">
                  <p class="item-title">{item.post.description.substring(0, 50)}...</p>
                  <p class="item-category">{item.post.category.toUpperCase()}</p>
                </div>
              </div>
            {/each}
          </div>
          <div class="summary-total">
            <span>Total Items:</span>
            <strong>{cartItems.length}</strong>
          </div>
        {:else}
          <p class="empty-summary">No items in cart</p>
        {/if}
      </div>
    </div>
  </div>
</div>

<style>
  .checkout-page {
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
  }

  .checkout-container {
    display: grid;
    grid-template-columns: 1fr 400px;
    gap: 30px;
  }

  .checkout-form {
    display: flex;
    flex-direction: column;
    gap: 30px;
  }

  .form-section {
    background-color: var(--white);
    padding: 24px;
    border-radius: 12px;
    border: 1px solid var(--border-color);
  }

  .form-section h2 {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 15px;
  }

  .form-group {
    margin-bottom: 15px;
  }

  .form-group label {
    display: block;
    font-weight: 600;
    margin-bottom: 8px;
  }

  input[type="text"],
  input[type="email"],
  input[type="tel"] {
    width: 100%;
    padding: 12px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-size: 16px;
  }

  .payment-options {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .payment-option {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 15px;
    border: 2px solid var(--border-color);
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.2s;
  }

  .payment-option:hover {
    border-color: var(--primary-color);
  }

  .payment-option input[type="radio"] {
    width: 20px;
    height: 20px;
  }

  .checkbox-group {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .checkbox-group label {
    display: flex;
    gap: 10px;
    cursor: pointer;
  }

  .checkbox-group input[type="checkbox"] {
    width: 20px;
    height: 20px;
    flex-shrink: 0;
  }

  .checkbox-group a {
    color: var(--primary-color);
    text-decoration: underline;
  }

  .complete-btn {
    width: 100%;
    padding: 16px;
    font-size: 18px;
    font-weight: 700;
  }

  .order-summary {
    position: sticky;
    top: 20px;
    height: fit-content;
  }

  .summary-card {
    background-color: var(--white);
    padding: 24px;
    border-radius: 12px;
    border: 1px solid var(--border-color);
  }

  .summary-card h2 {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .summary-items {
    display: flex;
    flex-direction: column;
    gap: 15px;
    margin-bottom: 20px;
    max-height: 400px;
    overflow-y: auto;
  }

  .summary-item {
    display: flex;
    gap: 12px;
  }

  .summary-item img {
    width: 60px;
    height: 60px;
    object-fit: cover;
    border-radius: 6px;
  }

  .item-info {
    flex: 1;
  }

  .item-title {
    font-size: 14px;
    margin-bottom: 4px;
  }

  .item-category {
    font-size: 12px;
    color: var(--text-secondary);
  }

  .summary-total {
    display: flex;
    justify-content: space-between;
    padding-top: 15px;
    border-top: 1px solid var(--border-color);
    font-size: 16px;
  }

  .empty-summary {
    text-align: center;
    color: var(--text-secondary);
    padding: 20px;
  }

  @media (max-width: 768px) {
    .checkout-container {
      grid-template-columns: 1fr;
    }

    .form-row {
      grid-template-columns: 1fr;
    }
  }
</style>

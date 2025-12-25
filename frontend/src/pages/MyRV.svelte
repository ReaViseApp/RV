<script lang="ts">
  import { onMount } from 'svelte';
  import RubiksCylinder from '../components/RubiksCylinder.svelte';
  import { feedMode } from '../stores';
  import type { Post } from '../types';
  
  let posts: { lot: Post[], design: Post[], reavise: Post[] } = {
    lot: [],
    design: [],
    reavise: []
  };

  onMount(() => {
    // Fetch posts from API
    loadPosts();
  });

  async function loadPosts() {
    // TODO: Implement API call
    // For now, using mock data
    console.log('Loading posts for feed mode:', $feedMode);
  }

  function toggleFeedMode() {
    feedMode.update(mode => mode === 'following' ? 'foryou' : 'following');
  }
</script>

<div class="myrv-page">
  <div class="header">
    <h1>MyRV</h1>
    <div class="feed-toggle">
      <button 
        class="toggle-btn" 
        class:active={$feedMode === 'following'}
        on:click={toggleFeedMode}
      >
        Following
      </button>
      <button 
        class="toggle-btn" 
        class:active={$feedMode === 'foryou'}
        on:click={toggleFeedMode}
      >
        RVed for You
      </button>
    </div>
  </div>

  <div class="cylinder-wrapper">
    <RubiksCylinder {posts} />
  </div>

  <div class="feed-info">
    <div class="row-label lot">The Lot</div>
    <div class="row-label design">Design</div>
    <div class="row-label reavise">ReaVise</div>
  </div>
</div>

<style>
  .myrv-page {
    max-width: 1400px;
    margin: 0 auto;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 30px;
    padding: 20px 0;
    border-bottom: 1px solid var(--border-color);
  }

  h1 {
    font-size: 32px;
    font-weight: 700;
  }

  .feed-toggle {
    display: flex;
    gap: 10px;
  }

  .toggle-btn {
    padding: 10px 20px;
    background-color: var(--white);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-weight: 600;
    transition: all 0.2s;
  }

  .toggle-btn.active {
    background-color: var(--primary-color);
    color: var(--white);
    border-color: var(--primary-color);
  }

  .cylinder-wrapper {
    margin: 30px 0;
  }

  .feed-info {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
  }

  .row-label {
    padding: 8px 16px;
    border-radius: 20px;
    font-weight: 600;
    color: var(--white);
    font-size: 14px;
  }

  .row-label.lot {
    background-color: #e91e63;
  }

  .row-label.design {
    background-color: #2196f3;
  }

  .row-label.reavise {
    background-color: #4caf50;
  }
</style>

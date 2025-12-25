<script lang="ts">
  import { onMount } from 'svelte';
  import RubiksCube from '../components/RubiksCube.svelte';
  import UserCard from '../components/UserCard.svelte';
  import type { Post, User } from '../types';
  
  let searchQuery: string = '';
  let searchResults: { users: User[], posts: { lot: Post[], design: Post[], reavise: Post[] } } = {
    users: [],
    posts: { lot: [], design: [], reavise: [] }
  };

  onMount(() => {
    // Initialize search
  });

  async function handleSearch() {
    if (!searchQuery.trim()) return;
    
    // TODO: Implement API call
    console.log('Searching for:', searchQuery);
  }

  function handleKeyPress(event: KeyboardEvent) {
    if (event.key === 'Enter') {
      handleSearch();
    }
  }
</script>

<div class="cubeativity-page">
  <div class="header">
    <h1>Cubeativity</h1>
    <div class="search-bar">
      <input 
        type="text" 
        placeholder="Search accounts and hashtags..." 
        bind:value={searchQuery}
        on:keypress={handleKeyPress}
      />
      <button class="search-btn" on:click={handleSearch}>
        🔍 Search
      </button>
    </div>
  </div>

  <div class="cube-wrapper">
    <RubiksCube posts={searchResults.posts} />
  </div>

  {#if searchResults.users.length > 0}
    <div class="search-results">
      <h2>Users</h2>
      <div class="users-grid">
        {#each searchResults.users as user}
          <UserCard {user} />
        {/each}
      </div>
    </div>
  {/if}

  <div class="cube-info">
    <div class="column-label lot">The Lot</div>
    <div class="column-label design">Design</div>
    <div class="column-label reavise">ReaVise</div>
  </div>
</div>

<style>
  .cubeativity-page {
    max-width: 1400px;
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
    margin-bottom: 20px;
  }

  .search-bar {
    display: flex;
    gap: 10px;
    max-width: 600px;
  }

  input {
    flex: 1;
    padding: 12px 16px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-size: 16px;
  }

  .search-btn {
    padding: 12px 24px;
    background-color: var(--primary-color);
    color: var(--white);
    border-radius: 8px;
    font-weight: 600;
    transition: background-color 0.2s;
  }

  .search-btn:hover {
    background-color: var(--secondary-color);
  }

  .cube-wrapper {
    margin: 30px 0;
  }

  .search-results {
    margin: 30px 0;
  }

  h2 {
    font-size: 24px;
    margin-bottom: 20px;
  }

  .users-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 20px;
  }

  .cube-info {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
  }

  .column-label {
    padding: 8px 16px;
    border-radius: 20px;
    font-weight: 600;
    color: var(--white);
    font-size: 14px;
  }

  .column-label.lot {
    background-color: #e91e63;
  }

  .column-label.design {
    background-color: #2196f3;
  }

  .column-label.reavise {
    background-color: #4caf50;
  }
</style>

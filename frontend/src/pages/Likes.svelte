<script lang="ts">
  import { onMount } from 'svelte';
  import PostCard from '../components/PostCard.svelte';
  import type { Post } from '../types';

  let likedPosts: Post[] = [];

  onMount(() => {
    loadLikedPosts();
  });

  async function loadLikedPosts() {
    // TODO: Implement API call
    console.log('Loading liked posts...');
  }
</script>

<div class="likes-page">
  <div class="header">
    <h1>Likes</h1>
    <p class="subtitle">Posts you've liked</p>
  </div>

  {#if likedPosts.length > 0}
    <div class="posts-grid">
      {#each likedPosts as post}
        <PostCard {post} />
      {/each}
    </div>
  {:else}
    <div class="empty-state">
      <p class="empty-icon">❤️</p>
      <h2>No liked posts yet</h2>
      <p>Start exploring and like posts you love</p>
      <a href="/" class="btn btn-primary">Explore Posts</a>
    </div>
  {/if}
</div>

<style>
  .likes-page {
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

  .posts-grid {
    display: grid;
    gap: 20px;
  }

  .empty-state {
    text-align: center;
    padding: 80px 20px;
  }

  .empty-icon {
    font-size: 80px;
    margin-bottom: 20px;
  }

  .empty-state h2 {
    font-size: 24px;
    margin-bottom: 10px;
  }

  .empty-state p {
    color: var(--text-secondary);
    margin-bottom: 30px;
  }
</style>

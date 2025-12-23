<script lang="ts">
  import type { Post } from '../types';
  
  export let post: Post;
</script>

<div class="post-card">
  <div class="post-header">
    <img src={post.userAvatar || '/default-avatar.png'} alt={post.username} class="avatar" />
    <div class="user-info">
      <span class="username">{post.username}</span>
      {#if post.userLocation}
        <span class="location">{post.userLocation}</span>
      {/if}
    </div>
    <span class="post-time">{new Date(post.createdAt).toLocaleDateString()}</span>
  </div>
  
  <div class="post-media">
    {#if post.media.length > 0}
      {#if post.media[0].type === 'image'}
        <img src={post.media[0].url} alt={post.description} />
      {:else}
        <video src={post.media[0].url} controls />
      {/if}
    {/if}
  </div>
  
  <div class="post-actions">
    <button class="action-btn" class:liked={post.isLiked}>
      ❤️ {post.likesCount}
    </button>
    <button class="action-btn">
      💬 {post.commentsCount}
    </button>
    <button class="action-btn curate-btn">
      Curate
    </button>
  </div>
  
  <div class="post-description">
    <span class="username">{post.username}</span>
    {post.description}
  </div>
  
  {#if post.hashtags.length > 0}
    <div class="hashtags">
      {#each post.hashtags as tag}
        <span class="hashtag">#{tag}</span>
      {/each}
    </div>
  {/if}
  
  <div class="category-badge" class:lot={post.category === 'lot'} 
       class:design={post.category === 'design'} 
       class:reavise={post.category === 'reavise'}>
    {post.category.toUpperCase()}
  </div>
</div>

<style>
  .post-card {
    background-color: var(--white);
    border: 1px solid var(--border-color);
    border-radius: 8px;
    margin-bottom: 20px;
    position: relative;
  }

  .post-header {
    display: flex;
    align-items: center;
    padding: 12px 16px;
    gap: 10px;
  }

  .avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    object-fit: cover;
  }

  .user-info {
    flex: 1;
    display: flex;
    flex-direction: column;
  }

  .username {
    font-weight: 600;
    font-size: 14px;
  }

  .location {
    font-size: 12px;
    color: var(--text-secondary);
  }

  .post-time {
    font-size: 12px;
    color: var(--text-secondary);
  }

  .post-media {
    width: 100%;
    max-height: 600px;
    overflow: hidden;
  }

  .post-media img,
  .post-media video {
    width: 100%;
    height: auto;
    object-fit: cover;
  }

  .post-actions {
    display: flex;
    gap: 15px;
    padding: 12px 16px;
  }

  .action-btn {
    background: none;
    border: none;
    font-size: 14px;
    display: flex;
    align-items: center;
    gap: 5px;
    cursor: pointer;
    transition: transform 0.2s;
  }

  .action-btn:hover {
    transform: scale(1.1);
  }

  .action-btn.liked {
    color: var(--error-color);
  }

  .curate-btn {
    margin-left: auto;
    background-color: var(--primary-color);
    color: var(--white);
    padding: 6px 12px;
    border-radius: 4px;
  }

  .post-description {
    padding: 0 16px 8px;
    font-size: 14px;
  }

  .hashtags {
    padding: 0 16px 12px;
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
  }

  .hashtag {
    color: var(--primary-color);
    font-size: 14px;
  }

  .category-badge {
    position: absolute;
    top: 60px;
    right: 10px;
    padding: 4px 8px;
    border-radius: 4px;
    font-size: 11px;
    font-weight: 600;
    color: var(--white);
  }

  .category-badge.lot {
    background-color: #e91e63;
  }

  .category-badge.design {
    background-color: #2196f3;
  }

  .category-badge.reavise {
    background-color: #4caf50;
  }
</style>

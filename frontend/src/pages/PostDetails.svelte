<script lang="ts">
  import { onMount } from 'svelte';
  import type { Post, Comment } from '../types';
  
  export let postId: string;
  
  let post: Post | null = null;
  let comments: Comment[] = [];
  let newComment: string = '';

  onMount(() => {
    loadPost();
    loadComments();
  });

  async function loadPost() {
    // TODO: Implement API call
    console.log('Loading post:', postId);
  }

  async function loadComments() {
    // TODO: Implement API call
    console.log('Loading comments for post:', postId);
  }

  async function submitComment() {
    if (!newComment.trim()) return;
    
    // TODO: Implement API call
    console.log('Submitting comment:', newComment);
    newComment = '';
  }

  function handleLike() {
    if (!post) return;
    post.isLiked = !post.isLiked;
    post.likesCount += post.isLiked ? 1 : -1;
  }

  function handleCurate() {
    // TODO: Add to cart
    console.log('Adding to cart:', postId);
  }
</script>

<div class="post-details-page">
  {#if post}
    <div class="post-container">
      <div class="media-section">
        {#if post.media.length > 0}
          {#if post.media[0].type === 'image'}
            <img src={post.media[0].url} alt={post.description} class="main-media" />
          {:else}
            <video src={post.media[0].url} controls class="main-media" />
          {/if}
        {/if}
      </div>

      <div class="details-section">
        <div class="post-header">
          <img src={post.userAvatar || '/default-avatar.png'} alt={post.username} class="avatar" />
          <div class="user-info">
            <span class="username">{post.username}</span>
            {#if post.userLocation}
              <span class="location">{post.userLocation}</span>
            {/if}
          </div>
          <span class="category-badge" class:lot={post.category === 'lot'} 
                class:design={post.category === 'design'} 
                class:reavise={post.category === 'reavise'}>
            {post.category.toUpperCase()}
          </span>
        </div>

        <div class="post-content">
          <p class="description">
            <strong>{post.username}</strong> {post.description}
          </p>
          {#if post.hashtags.length > 0}
            <div class="hashtags">
              {#each post.hashtags as tag}
                <span class="hashtag">#{tag}</span>
              {/each}
            </div>
          {/if}
          <p class="post-date">{new Date(post.createdAt).toLocaleString()}</p>
        </div>

        <div class="post-actions">
          <button class="action-btn" class:liked={post.isLiked} on:click={handleLike}>
            ❤️ {post.likesCount}
          </button>
          <button class="action-btn">
            💬 {post.commentsCount}
          </button>
          <button class="curate-btn" on:click={handleCurate}>
            Curate
          </button>
        </div>

        <div class="comments-section">
          <h3>Comments</h3>
          <div class="comments-list">
            {#each comments as comment}
              <div class="comment">
                <img src={comment.userAvatar || '/default-avatar.png'} alt={comment.username} class="comment-avatar" />
                <div class="comment-content">
                  <p>
                    <strong>{comment.username}</strong> {comment.text}
                  </p>
                  <span class="comment-date">{new Date(comment.createdAt).toLocaleString()}</span>
                </div>
              </div>
            {/each}
          </div>

          <div class="comment-input">
            <input 
              type="text" 
              bind:value={newComment} 
              placeholder="Add a comment..."
              on:keypress={(e) => e.key === 'Enter' && submitComment()}
            />
            <button on:click={submitComment}>Post</button>
          </div>
        </div>
      </div>
    </div>
  {:else}
    <p>Loading post...</p>
  {/if}
</div>

<style>
  .post-details-page {
    max-width: 1200px;
    margin: 0 auto;
  }

  .post-container {
    display: grid;
    grid-template-columns: 1fr 500px;
    gap: 0;
    background-color: var(--white);
    border: 1px solid var(--border-color);
    border-radius: 12px;
    overflow: hidden;
    min-height: 600px;
  }

  .media-section {
    background-color: #000;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .main-media {
    max-width: 100%;
    max-height: 800px;
    object-fit: contain;
  }

  .details-section {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .post-header {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 16px;
    border-bottom: 1px solid var(--border-color);
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

  .category-badge {
    padding: 4px 10px;
    border-radius: 12px;
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

  .post-content {
    padding: 16px;
    border-bottom: 1px solid var(--border-color);
  }

  .description {
    margin-bottom: 10px;
    line-height: 1.6;
  }

  .hashtags {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
    margin-bottom: 10px;
  }

  .hashtag {
    color: var(--primary-color);
    font-size: 14px;
  }

  .post-date {
    font-size: 12px;
    color: var(--text-secondary);
  }

  .post-actions {
    display: flex;
    gap: 15px;
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color);
  }

  .action-btn {
    background: none;
    border: none;
    font-size: 14px;
    display: flex;
    align-items: center;
    gap: 5px;
    cursor: pointer;
  }

  .action-btn.liked {
    color: var(--error-color);
  }

  .curate-btn {
    margin-left: auto;
    background-color: var(--primary-color);
    color: var(--white);
    padding: 6px 16px;
    border-radius: 4px;
    border: none;
    font-weight: 600;
    cursor: pointer;
  }

  .comments-section {
    flex: 1;
    display: flex;
    flex-direction: column;
    padding: 16px;
  }

  .comments-section h3 {
    font-size: 16px;
    margin-bottom: 15px;
  }

  .comments-list {
    flex: 1;
    overflow-y: auto;
    margin-bottom: 15px;
  }

  .comment {
    display: flex;
    gap: 12px;
    margin-bottom: 15px;
  }

  .comment-avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    object-fit: cover;
  }

  .comment-content {
    flex: 1;
  }

  .comment-content p {
    margin-bottom: 4px;
    font-size: 14px;
  }

  .comment-date {
    font-size: 11px;
    color: var(--text-secondary);
  }

  .comment-input {
    display: flex;
    gap: 10px;
    border-top: 1px solid var(--border-color);
    padding-top: 15px;
  }

  .comment-input input {
    flex: 1;
    padding: 8px 12px;
    border: 1px solid var(--border-color);
    border-radius: 20px;
  }

  .comment-input button {
    padding: 8px 20px;
    background-color: var(--primary-color);
    color: var(--white);
    border-radius: 20px;
    font-weight: 600;
    border: none;
    cursor: pointer;
  }

  @media (max-width: 768px) {
    .post-container {
      grid-template-columns: 1fr;
    }
  }
</style>

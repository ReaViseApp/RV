<script lang="ts">
  import { onMount } from 'svelte';
  import RubiksCylinder from '../components/RubiksCylinder.svelte';
  import type { User, Post } from '../types';
  
  export let userId: string;
  
  let user: User | null = null;
  let posts: { lot: Post[], design: Post[], reavise: Post[] } = {
    lot: [],
    design: [],
    reavise: []
  };
  let isFollowing: boolean = false;

  onMount(() => {
    loadUserProfile();
  });

  async function loadUserProfile() {
    // TODO: Implement API call
    console.log('Loading profile for user:', userId);
  }

  function handleFollow() {
    isFollowing = !isFollowing;
  }

  function handleMessage() {
    // Navigate to Exchange with this user
    console.log('Message user:', userId);
  }
</script>

<div class="profile-page">
  {#if user}
    <div class="profile-header">
      <img src={user.profilePhoto || '/default-avatar.png'} alt={user.username} class="profile-photo" />
      
      <div class="profile-info">
        <div class="username-row">
          <h1>{user.username}</h1>
          {#if user.isVerified}
            <span class="verified-badge">✓</span>
          {/if}
          {#if user.isBusinessAccount}
            <span class="business-badge">Business</span>
          {/if}
        </div>
        
        <div class="stats">
          <div class="stat">
            <strong>{user.followersCount}</strong>
            <span>followers</span>
          </div>
          <div class="stat">
            <strong>{user.followingCount}</strong>
            <span>following</span>
          </div>
        </div>
        
        <div class="actions">
          <button class="btn btn-primary" on:click={handleFollow}>
            {isFollowing ? 'Unfollow' : 'Follow'}
          </button>
          <button class="btn btn-secondary" on:click={handleMessage}>
            Message
          </button>
        </div>
        
        {#if user.bio}
          <p class="bio">{user.bio}</p>
        {/if}
        
        {#if user.website}
          <a href={user.website} target="_blank" class="website">{user.website}</a>
        {/if}
        
        {#if user.email}
          <p class="email">{user.email}</p>
        {/if}
      </div>
    </div>

    <div class="content-section">
      <h2>User Content</h2>
      <RubiksCylinder {posts} />
    </div>
  {:else}
    <p>Loading profile...</p>
  {/if}
</div>

<style>
  .profile-page {
    max-width: 1400px;
    margin: 0 auto;
  }

  .profile-header {
    display: flex;
    gap: 40px;
    padding: 40px 0;
    border-bottom: 1px solid var(--border-color);
  }

  .profile-photo {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    object-fit: cover;
  }

  .profile-info {
    flex: 1;
  }

  .username-row {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 20px;
  }

  h1 {
    font-size: 28px;
    font-weight: 400;
  }

  .verified-badge {
    background-color: var(--primary-color);
    color: var(--white);
    width: 24px;
    height: 24px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
  }

  .business-badge {
    background-color: var(--accent-color);
    color: var(--white);
    padding: 4px 12px;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 600;
  }

  .stats {
    display: flex;
    gap: 40px;
    margin-bottom: 20px;
  }

  .stat {
    display: flex;
    gap: 5px;
  }

  .stat strong {
    font-weight: 600;
  }

  .stat span {
    color: var(--text-secondary);
  }

  .actions {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }

  .bio {
    margin-bottom: 10px;
    line-height: 1.6;
  }

  .website {
    display: block;
    color: var(--primary-color);
    margin-bottom: 5px;
  }

  .email {
    color: var(--text-secondary);
  }

  .content-section {
    margin-top: 40px;
  }

  h2 {
    font-size: 24px;
    margin-bottom: 20px;
  }
</style>

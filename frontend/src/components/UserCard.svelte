<script lang="ts">
  import type { User } from '../types';
  
  export let user: User;
  export let showFollowButton: boolean = true;
</script>

<div class="user-card">
  <img src={user.profilePhoto || '/default-avatar.png'} alt={user.username} class="profile-photo" />
  
  <div class="user-details">
    <div class="username-row">
      <span class="username">{user.username}</span>
      {#if user.isVerified}
        <span class="verified-badge">✓</span>
      {/if}
    </div>
    
    {#if user.bio}
      <p class="bio">{user.bio}</p>
    {/if}
    
    <div class="stats">
      <span class="stat">
        <strong>{user.followersCount}</strong> followers
      </span>
      <span class="stat">
        <strong>{user.followingCount}</strong> following
      </span>
    </div>
  </div>
  
  {#if showFollowButton}
    <button class="btn btn-primary follow-btn">Follow</button>
  {/if}
</div>

<style>
  .user-card {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 16px;
    background-color: var(--white);
    border: 1px solid var(--border-color);
    border-radius: 8px;
  }

  .profile-photo {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    object-fit: cover;
  }

  .user-details {
    flex: 1;
  }

  .username-row {
    display: flex;
    align-items: center;
    gap: 5px;
    margin-bottom: 4px;
  }

  .username {
    font-weight: 600;
    font-size: 16px;
  }

  .verified-badge {
    background-color: var(--primary-color);
    color: var(--white);
    width: 18px;
    height: 18px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
  }

  .bio {
    font-size: 14px;
    color: var(--text-secondary);
    margin-bottom: 8px;
  }

  .stats {
    display: flex;
    gap: 15px;
    font-size: 14px;
  }

  .stat {
    color: var(--text-secondary);
  }

  .follow-btn {
    padding: 8px 24px;
  }
</style>

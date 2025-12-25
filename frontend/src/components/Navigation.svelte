<script lang="ts">
  import { Link } from 'svelte-routing';
  import { currentUser } from '../stores';
</script>

<nav class="sidebar">
  <div class="logo">
    <h1>ReaVise</h1>
  </div>
  
  <div class="nav-items">
    <Link to="/" class="nav-item">
      <span class="icon">🏠</span>
      <span class="label">MyRV</span>
    </Link>
    
    <Link to="/curated" class="nav-item">
      <span class="icon">🛒</span>
      <span class="label">Curated</span>
    </Link>
    
    <Link to="/search" class="nav-item">
      <span class="icon">🔍</span>
      <span class="label">Cubeativity</span>
    </Link>
    
    <Link to="/upload" class="nav-item">
      <span class="icon">➕</span>
      <span class="label">Upload</span>
    </Link>
    
    <Link to="/likes" class="nav-item">
      <span class="icon">❤️</span>
      <span class="label">Likes</span>
    </Link>
    
    <Link to="/exchange" class="nav-item">
      <span class="icon">💬</span>
      <span class="label">Exchange</span>
    </Link>
    
    <Link to="/profile/{$currentUser?.id || 'me'}" class="nav-item">
      <span class="icon">👤</span>
      <span class="label">Profile</span>
    </Link>
    
    <Link to="/settings" class="nav-item">
      <span class="icon">⚙️</span>
      <span class="label">Settings</span>
    </Link>
  </div>
</nav>

<style>
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    width: 250px;
    height: 100vh;
    background-color: var(--white);
    border-right: 1px solid var(--border-color);
    padding: 20px;
    overflow-y: auto;
  }

  .logo {
    margin-bottom: 40px;
  }

  .logo h1 {
    font-size: 28px;
    font-weight: 700;
    color: var(--primary-color);
  }

  .nav-items {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .nav-item {
    display: flex;
    align-items: center;
    gap: 15px;
    padding: 12px 16px;
    border-radius: 8px;
    color: var(--text-color);
    transition: background-color 0.2s;
    text-decoration: none;
  }

  .nav-item:hover {
    background-color: var(--background-color);
  }

  .icon {
    font-size: 24px;
  }

  .label {
    font-size: 16px;
    font-weight: 500;
  }

  @media (max-width: 768px) {
    .sidebar {
      width: 100%;
      height: auto;
      position: relative;
      border-right: none;
      border-bottom: 1px solid var(--border-color);
    }

    .nav-items {
      flex-direction: row;
      justify-content: space-around;
      overflow-x: auto;
    }

    .label {
      display: none;
    }
  }
</style>

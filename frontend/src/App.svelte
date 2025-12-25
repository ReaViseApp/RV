<script lang="ts">
  import { Router, Route } from 'svelte-routing';
  import Navigation from './components/Navigation.svelte';
  import MyRV from './pages/MyRV.svelte';
  import Profile from './pages/Profile.svelte';
  import Cubeativity from './pages/Cubeativity.svelte';
  import Upload from './pages/Upload.svelte';
  import Curated from './pages/Curated.svelte';
  import Likes from './pages/Likes.svelte';
  import Exchange from './pages/Exchange.svelte';
  import Settings from './pages/Settings.svelte';
  import Checkout from './pages/Checkout.svelte';
  import PostDetails from './pages/PostDetails.svelte';
</script>

<Router>
  <div class="app">
    <Navigation />
    <main class="main-content">
      <Route path="/" component={MyRV} />
      <Route path="/profile/:userId" component={Profile} />
      <Route path="/search" component={Cubeativity} />
      <Route path="/upload" component={Upload} />
      <Route path="/curated" component={Curated} />
      <Route path="/likes" component={Likes} />
      <Route path="/exchange" component={Exchange} />
      <Route path="/settings" component={Settings} />
      <Route path="/checkout" component={Checkout} />
      <Route path="/post/:postId" component={PostDetails} />
    </main>
  </div>
</Router>

<style>
  .app {
    display: flex;
    min-height: 100vh;
    background-color: #fafafa;
  }

  .main-content {
    flex: 1;
    margin-left: 250px;
    padding: 20px;
  }

  @media (max-width: 768px) {
    .main-content {
      margin-left: 0;
      padding: 10px;
    }
  }
</style>

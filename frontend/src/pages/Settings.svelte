<script lang="ts">
  import { currentUser } from '../stores';
  
  let editMode: boolean = false;
  let username: string = '';
  let email: string = '';
  let bio: string = '';
  let website: string = '';

  $: if ($currentUser) {
    username = $currentUser.username;
    email = $currentUser.email;
    bio = $currentUser.bio || '';
    website = $currentUser.website || '';
  }

  function handleSave() {
    // TODO: Implement API call
    console.log('Saving settings...', { username, email, bio, website });
    editMode = false;
  }

  function handleLogout() {
    // TODO: Implement logout
    console.log('Logging out...');
  }

  function handleExportData() {
    // TODO: Implement data export
    console.log('Exporting user data...');
  }

  function handleDeleteAccount() {
    if (confirm('Are you sure you want to delete your account? This action cannot be undone.')) {
      // TODO: Implement account deletion
      console.log('Deleting account...');
    }
  }
</script>

<div class="settings-page">
  <div class="header">
    <h1>Settings</h1>
  </div>

  <div class="settings-container">
    <section class="settings-section">
      <h2>Profile Settings</h2>
      <div class="form-group">
        <label>Username</label>
        <input type="text" bind:value={username} disabled={!editMode} />
      </div>
      <div class="form-group">
        <label>Email</label>
        <input type="email" bind:value={email} disabled={!editMode} />
      </div>
      <div class="form-group">
        <label>Bio</label>
        <textarea bind:value={bio} disabled={!editMode} rows="4" />
      </div>
      <div class="form-group">
        <label>Website</label>
        <input type="url" bind:value={website} disabled={!editMode} />
      </div>
      <div class="form-actions">
        {#if !editMode}
          <button class="btn btn-primary" on:click={() => editMode = true}>
            Edit Profile
          </button>
        {:else}
          <button class="btn btn-secondary" on:click={() => editMode = false}>
            Cancel
          </button>
          <button class="btn btn-primary" on:click={handleSave}>
            Save Changes
          </button>
        {/if}
      </div>
    </section>

    <section class="settings-section">
      <h2>Privacy & Data</h2>
      <div class="info-text">
        <p>Your privacy is important to us. You can export your data or request account deletion at any time.</p>
      </div>
      <div class="action-buttons">
        <button class="btn btn-secondary" on:click={handleExportData}>
          Export My Data
        </button>
        <a href="/legal/privacy-policy" class="link">View Privacy Policy</a>
        <a href="/legal/terms-of-service" class="link">View Terms of Service</a>
      </div>
    </section>

    <section class="settings-section">
      <h2>Account Actions</h2>
      <div class="action-buttons">
        <button class="btn btn-secondary" on:click={handleLogout}>
          Logout
        </button>
        <button class="btn btn-danger" on:click={handleDeleteAccount}>
          Delete Account
        </button>
      </div>
    </section>
  </div>
</div>

<style>
  .settings-page {
    max-width: 800px;
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

  .settings-container {
    display: flex;
    flex-direction: column;
    gap: 30px;
  }

  .settings-section {
    background-color: var(--white);
    padding: 30px;
    border-radius: 12px;
    border: 1px solid var(--border-color);
  }

  .settings-section h2 {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .form-group {
    margin-bottom: 20px;
  }

  .form-group label {
    display: block;
    font-weight: 600;
    margin-bottom: 8px;
  }

  input, textarea {
    width: 100%;
    padding: 12px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-size: 16px;
  }

  input:disabled, textarea:disabled {
    background-color: var(--background-color);
    cursor: not-allowed;
  }

  .form-actions {
    display: flex;
    gap: 10px;
    justify-content: flex-end;
  }

  .info-text {
    margin-bottom: 20px;
  }

  .info-text p {
    color: var(--text-secondary);
    line-height: 1.6;
  }

  .action-buttons {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .link {
    padding: 8px 0;
    color: var(--primary-color);
  }

  .btn-danger {
    background-color: var(--error-color);
    color: var(--white);
    padding: 12px 24px;
    border-radius: 8px;
    font-weight: 600;
  }
</style>

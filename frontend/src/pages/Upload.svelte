<script lang="ts">
  let files: FileList | null = null;
  let description: string = '';
  let category: 'lot' | 'design' | 'reavise' = 'design';
  let hashtags: string = '';
  let isBusinessAccount: boolean = false;
  let multiCategoryLabels: Record<number, 'lot' | 'design' | 'reavise'> = {};

  async function handleSubmit() {
    if (!files || files.length === 0) {
      alert('Please select at least one file');
      return;
    }

    const formData = new FormData();
    
    for (let i = 0; i < files.length; i++) {
      formData.append('media', files[i]);
      if (files.length > 1) {
        formData.append(`category_${i}`, multiCategoryLabels[i] || category);
      }
    }
    
    formData.append('description', description);
    formData.append('category', category);
    formData.append('hashtags', hashtags);

    // TODO: Implement API call
    console.log('Uploading post...', { description, category, hashtags });
  }

  function handleFileChange(event: Event) {
    const target = event.target as HTMLInputElement;
    files = target.files;
    
    if (files && files.length > 1) {
      // Initialize category labels for multiple files
      for (let i = 0; i < files.length; i++) {
        multiCategoryLabels[i] = category;
      }
    }
  }
</script>

<div class="upload-page">
  <div class="header">
    <h1>Create New Post</h1>
  </div>

  <div class="upload-form">
    <div class="file-upload">
      <label for="media-files" class="upload-label">
        <div class="upload-box">
          {#if files && files.length > 0}
            <p>✓ {files.length} file(s) selected</p>
          {:else}
            <p>➕ Click to upload images or videos</p>
          {/if}
        </div>
      </label>
      <input 
        id="media-files" 
        type="file" 
        multiple 
        accept="image/*,video/*" 
        on:change={handleFileChange}
        style="display: none;"
      />
    </div>

    {#if files && files.length > 1}
      <div class="multi-category">
        <h3>Category for each file:</h3>
        {#each Array.from(files) as file, index}
          <div class="file-category">
            <span class="file-name">{file.name}</span>
            <select bind:value={multiCategoryLabels[index]}>
              <option value="lot">The Lot</option>
              <option value="design">Design</option>
              <option value="reavise">ReaVise</option>
            </select>
          </div>
        {/each}
      </div>
    {/if}

    <div class="form-group">
      <label for="description">Description</label>
      <textarea 
        id="description" 
        bind:value={description} 
        placeholder="Write a description..."
        rows="4"
      />
    </div>

    <div class="form-group">
      <label for="category">Main Category</label>
      <select id="category" bind:value={category}>
        <option value="lot">The Lot</option>
        <option value="design">Design</option>
        <option value="reavise">ReaVise</option>
      </select>
    </div>

    <div class="form-group">
      <label for="hashtags">Hashtags</label>
      <input 
        id="hashtags" 
        type="text" 
        bind:value={hashtags} 
        placeholder="art design creative (space-separated)"
      />
    </div>

    {#if isBusinessAccount}
      <div class="business-info">
        <p>✓ As a verified business account, you can label posts as both Design and ReaVise</p>
      </div>
    {/if}

    <div class="form-actions">
      <button type="button" class="btn btn-secondary">Cancel</button>
      <button type="button" class="btn btn-primary" on:click={handleSubmit}>
        Publish
      </button>
    </div>
  </div>
</div>

<style>
  .upload-page {
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

  .upload-form {
    background-color: var(--white);
    padding: 30px;
    border-radius: 12px;
    border: 1px solid var(--border-color);
  }

  .upload-label {
    cursor: pointer;
  }

  .upload-box {
    border: 2px dashed var(--border-color);
    border-radius: 12px;
    padding: 60px;
    text-align: center;
    transition: all 0.2s;
  }

  .upload-box:hover {
    border-color: var(--primary-color);
    background-color: var(--background-color);
  }

  .upload-box p {
    font-size: 18px;
    color: var(--text-secondary);
  }

  .multi-category {
    margin: 20px 0;
    padding: 20px;
    background-color: var(--background-color);
    border-radius: 8px;
  }

  .multi-category h3 {
    font-size: 16px;
    margin-bottom: 15px;
  }

  .file-category {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    margin-bottom: 10px;
    background-color: var(--white);
    border-radius: 6px;
  }

  .file-name {
    flex: 1;
    font-size: 14px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .form-group {
    margin: 20px 0;
  }

  label {
    display: block;
    font-weight: 600;
    margin-bottom: 8px;
  }

  input, textarea, select {
    width: 100%;
    padding: 12px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    font-size: 16px;
  }

  textarea {
    resize: vertical;
  }

  .business-info {
    padding: 15px;
    background-color: #e3f2fd;
    border-radius: 8px;
    margin: 20px 0;
  }

  .business-info p {
    color: var(--primary-color);
    margin: 0;
  }

  .form-actions {
    display: flex;
    gap: 10px;
    justify-content: flex-end;
    margin-top: 30px;
  }

  .btn {
    padding: 12px 30px;
  }
</style>

<script lang="ts">
  import { onMount } from 'svelte';
  import type { Message, User } from '../types';

  let conversations: { user: User, lastMessage: Message }[] = [];
  let selectedConversation: { user: User, messages: Message[] } | null = null;
  let newMessage: string = '';

  onMount(() => {
    loadConversations();
  });

  async function loadConversations() {
    // TODO: Implement API call
    console.log('Loading conversations...');
  }

  async function selectConversation(user: User) {
    // TODO: Load messages for this user
    console.log('Loading messages with:', user.username);
  }

  async function sendMessage() {
    if (!newMessage.trim() || !selectedConversation) return;

    // TODO: Implement API call
    console.log('Sending message:', newMessage);
    newMessage = '';
  }

  function handleKeyPress(event: KeyboardEvent) {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault();
      sendMessage();
    }
  }
</script>

<div class="exchange-page">
  <div class="header">
    <h1>Exchange</h1>
    <p class="subtitle">Your messages</p>
  </div>

  <div class="messaging-container">
    <div class="conversations-list">
      <h2>Messages</h2>
      {#if conversations.length > 0}
        {#each conversations as conv}
          <div 
            class="conversation-item" 
            class:active={selectedConversation?.user.id === conv.user.id}
            on:click={() => selectConversation(conv.user)}
          >
            <img src={conv.user.profilePhoto || '/default-avatar.png'} alt={conv.user.username} class="avatar" />
            <div class="conv-info">
              <span class="username">{conv.user.username}</span>
              <span class="last-message">{conv.lastMessage.text}</span>
            </div>
            {#if !conv.lastMessage.isRead}
              <span class="unread-badge"></span>
            {/if}
          </div>
        {/each}
      {:else}
        <p class="empty-text">No conversations yet</p>
      {/if}
    </div>

    <div class="chat-area">
      {#if selectedConversation}
        <div class="chat-header">
          <img src={selectedConversation.user.profilePhoto || '/default-avatar.png'} alt={selectedConversation.user.username} class="avatar" />
          <span class="username">{selectedConversation.user.username}</span>
        </div>

        <div class="messages-container">
          {#each selectedConversation.messages as message}
            <div class="message" class:own={message.senderId === 'currentUserId'}>
              <p>{message.text}</p>
              <span class="timestamp">{new Date(message.createdAt).toLocaleTimeString()}</span>
            </div>
          {/each}
        </div>

        <div class="message-input">
          <textarea 
            bind:value={newMessage}
            placeholder="Type a message..."
            on:keypress={handleKeyPress}
          />
          <button class="send-btn" on:click={sendMessage}>
            Send
          </button>
        </div>
      {:else}
        <div class="no-selection">
          <p class="empty-icon">💬</p>
          <p>Select a conversation to start messaging</p>
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  .exchange-page {
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
    margin-bottom: 5px;
  }

  .subtitle {
    color: var(--text-secondary);
    font-size: 16px;
  }

  .messaging-container {
    display: grid;
    grid-template-columns: 350px 1fr;
    gap: 0;
    height: 600px;
    border: 1px solid var(--border-color);
    border-radius: 12px;
    overflow: hidden;
    background-color: var(--white);
  }

  .conversations-list {
    border-right: 1px solid var(--border-color);
    overflow-y: auto;
  }

  .conversations-list h2 {
    padding: 20px;
    font-size: 20px;
    border-bottom: 1px solid var(--border-color);
  }

  .conversation-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 15px 20px;
    border-bottom: 1px solid var(--border-color);
    cursor: pointer;
    position: relative;
    transition: background-color 0.2s;
  }

  .conversation-item:hover {
    background-color: var(--background-color);
  }

  .conversation-item.active {
    background-color: var(--background-color);
  }

  .avatar {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    object-fit: cover;
  }

  .conv-info {
    flex: 1;
    overflow: hidden;
  }

  .username {
    font-weight: 600;
    display: block;
    margin-bottom: 4px;
  }

  .last-message {
    font-size: 14px;
    color: var(--text-secondary);
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .unread-badge {
    width: 10px;
    height: 10px;
    background-color: var(--primary-color);
    border-radius: 50%;
  }

  .empty-text {
    padding: 20px;
    text-align: center;
    color: var(--text-secondary);
  }

  .chat-area {
    display: flex;
    flex-direction: column;
  }

  .chat-header {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 15px 20px;
    border-bottom: 1px solid var(--border-color);
  }

  .chat-header .avatar {
    width: 40px;
    height: 40px;
  }

  .messages-container {
    flex: 1;
    padding: 20px;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .message {
    max-width: 60%;
    padding: 12px 16px;
    border-radius: 18px;
    background-color: var(--background-color);
  }

  .message.own {
    align-self: flex-end;
    background-color: var(--primary-color);
    color: var(--white);
  }

  .message p {
    margin-bottom: 4px;
  }

  .timestamp {
    font-size: 11px;
    opacity: 0.7;
  }

  .message-input {
    display: flex;
    gap: 10px;
    padding: 15px 20px;
    border-top: 1px solid var(--border-color);
  }

  .message-input textarea {
    flex: 1;
    padding: 10px;
    border: 1px solid var(--border-color);
    border-radius: 20px;
    resize: none;
    max-height: 100px;
  }

  .send-btn {
    padding: 10px 24px;
    background-color: var(--primary-color);
    color: var(--white);
    border-radius: 20px;
    font-weight: 600;
  }

  .no-selection {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }

  .empty-icon {
    font-size: 60px;
    margin-bottom: 15px;
  }

  .no-selection p:last-child {
    color: var(--text-secondary);
  }
</style>

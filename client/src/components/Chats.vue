<template>
  <div class="chats">
    <div class="left">
      <div class="top-label">
        <a href="/"><img src="../assets/balls.png" width="60" /></a>
        <span class="top_label_text">Chats</span>
      </div>
      <div class="chats-list">
        <input class="search_input" placeholder="Search chats" />
        <Chat
            v-bind:title="text"
        ></Chat>

      </div>
    </div>
    <transition>
      <div class="right" v-if="show">
        <h1>Hello this is a chat history</h1>
      </div>
    </transition>

  </div>

</template>
<script>
export default{
  data() {
    return {
      show: true
    }
  }
}
import styles from '../assets/chats.css'
</script>
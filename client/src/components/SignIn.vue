<template>
  <div class="auth">
    <div class="form">
      <span style="color: #00005C; font-size: 60px;">Glad to see You,</span> <br>
      <span class="auth_text">please continue to signin</span>
      <form class="auth_form">
        <label>Email</label>
        <input class="auth_input" required placeholder="Type your email" type="email">
        <label>Password</label>
        <input class="auth_input" required placeholder="Type your password" type="password">
        <button class="auth_button">Connect to my account</button>
      </form>
    </div>
    <div class="gradient">

      <div class="auth_help">
        <img src="../assets/logo.png" width="100"/>
        <h1 class="auth_help_text">You want to create a new account?</h1>
        <a href="/signup" class="auth_help_button">Sign Up</a>
      </div>
    </div>
  </div>
</template>
<script>
export default {};
import styles from '../assets/auth.css'
</script>
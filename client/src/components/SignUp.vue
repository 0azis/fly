<template>
  <div class="auth">
    <div class="form">
      <span style="color: #00005C; font-size: 60px;">Hello,</span> <br>
      <span class="auth_text">please continue to signup</span>
      <form class="auth_form">
        <label>Nick</label>
        <input class="auth_input" required placeholder="How another people will be see you" type="text" name="nick">
        <label>Email</label>
        <input class="auth_input" required placeholder="For security" type="email">
        <label>Password</label>
        <input class="auth_input" required placeholder="Type a hard password" type="password">
        <button class="auth_button">Create a secure account</button>
      </form>
    </div>
    <div class="gradient">

      <div class="auth_help">
        <img src="../assets/logo.png" width="100"/>
        <h1 class="auth_help_text">Already communicate with Fly?</h1>
        <a href="/signin" class="auth_help_button">Sign In</a>
      </div>
    </div>
  </div>
</template>
<script>
export default {};
import styles from '../assets/auth.css'
</script>
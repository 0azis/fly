<template>
  <div class="auth">
    <div class="form">
      <span style="color: #00005C; font-size: 60px;">Glad to see You,</span> <br>
      <span class="auth_text">please continue to signin</span>
      <form class="auth_form" v-on:submit.prevent="sendData">
        <label>Login</label>
        <input class="auth_input" v-bind:value="Login" v-on:input="changeLogin" required placeholder="Type email or nick" type="text">
        <label>Password</label>
        <input class="auth_input" v-bind:value="Password" v-on:input="changePassword" required placeholder="Type your password" type="password">
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
import axios from "axios";

export default {
  data() {
    return {
      Login: '',
      Password: ''
    }
  },
  methods: {
    sendData() {
      axios({
        method: 'post',
        url: 'http://localhost:8080/api/auth/signin',
        data: {
          login: this.Login,
          password: this.Password
        }
      }, {withCredentials: true}).then(function (response) {
        if (response.status == 200) {
          localStorage.setItem("token", response.data["access_token"])
        }
      });
    },
    changeLogin(event) {this.Login = event.target.value},
    changePassword(event) {this.Password = event.target.value}
  }
};
import styles from '../assets/auth.css'
</script>
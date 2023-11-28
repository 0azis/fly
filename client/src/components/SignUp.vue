<template>
  <div class="auth">
    <div class="form">
      <span style="color: #00005C; font-size: 60px;">Hello,</span> <br>
      <span class="auth_text">please continue to signup</span>
      <form class="auth_form" v-on:submit.prevent="sendData">
        <label>Nick</label>
        <input class="auth_input" v-bind:value="Nick" v-on:input="changeNick" required placeholder="How another people will be see you" type="text" name="nick">
        <label>Email</label>
        <input class="auth_input" v-bind:value="Email" v-on:input="changeEmail" required placeholder="For security" type="email">
        <label>Password</label>
        <input class="auth_input" v-bind:value="Password" v-on:input="changePassword" required placeholder="Type a hard password" type="password">
        <button type="submit" class="auth_button">Create a secure account</button>
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
import {ref} from "vue";

export default {
  data() {
    return {
      Nick: '',
      Email: '',
      Password: ''
    }
  },
  methods: {
    sendData() {
      console.log(this.Nick, this.Email, this.Password)
      axios({
        method: 'post',
        url: 'http://localhost:8080/api/auth/signup',
        data: {
          nick: this.Nick,
          email: this.Email,
          password: this.Password
        }
      }, {withCredentials: true}).then(function (response) {
        if (response.status == 200) {
          localStorage.setItem("token", response.data["access_token"])

        }
      });
    },
    changeNick(event) {this.Nick = event.target.value},
    changeEmail(event) {this.Email = event.target.value},
    changePassword(event) {this.Password = event.target.value}

  }
};
import styles from '../assets/auth.css'
import axios from "axios";
</script>
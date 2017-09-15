<template>
  <div class='container'>
    <div class='card'>
      <section class='header'>
        <h4> Create Your Account </h4>
      </section>
      <section class='form'>
        <input placeholder='Username' v-model="user.Username">
        <input placeholder='someone@email.com' v-model="user.Email">
        <input placeholder='Password' v-model="user.Password">
        <input placeholder='Confirm Password' v-model="user.ConfirmPassword">
      </section>
      <div class='menu-bar'>
        <button class='menu-button' @click="Handler" :disabled="!valid"> Create </button>
        <router-link to='/'> Cancel </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      user: {},
      valid: false
    };
  },
  methods: {
    CreateAccount: function () {

    },
    ConfirmPasswords: function () {
      return this.user.Password === this.user.ConfirmPassword;
    }
  }
};
</script>

<style scoped>
.container {
  margin: 50px;
}
.card {
  padding: 30px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
input {
  display: block;
  font-size: 16px;
  padding: 5px 10px;
  width: 500px;
  border: none;
  font-weight: 100;
  outline: none;
  color: #325778;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
  margin-top: 10px;
  transition: all 0.5s linear;
}
input:focus {
  border-bottom: 1px solid #325778;
}
a, a:visited {
  color: #325778;
}
a:hover {
  color: #1B3F60;
}
button, a {
  border-radius: 0;
  border: none;
}
button:hover {
  cursor: pointer;
}
.menu-bar {
  display: flex;
  margin-top: 25px;
  align-items: center;
  margin-left: 10px;
}
.menu-button {
  background-color: #529A7F;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  color: #fff;
  margin-right: 10px;
}
.menu-button:disabled,
.menu-button:disabled:hover {
  background-color: #919191;
  box-shadow: none;
  cursor: not-allowed;
}
.menu-button:hover {
  background-color: #176548;
}
</style>
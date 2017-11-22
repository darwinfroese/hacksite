<template>
  <div>
    <LoginHeader />
    <div class='container'>
      <div class='card' v-on:keyup.enter = "enterKeyPressHandler">
        <section class='header'>
          <h4> Create Your Account </h4>
        </section>
        <section class='form'>
          <div class='input-container'>
            <label class="label" for="name">Name</label>
            <p :class="{ 'control': true }">
              <input id="name" placeholder='Name' v-model="account.Name" v-validate="'required|min:1|max:64'" :class="{'input': true, 'is-danger': errors.has('name') }" name="name" type="text">
              <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </p>
          </div>
          <div class='input-container'>
            <label class="label" for="username">Username</label>
            <p :class="{ 'control': true }">
              <input id="username" placeholder='Username' v-model="account.Username" v-validate="'required|alpha_num|min:3|max:64'" :class="{'input': true, 'is-danger': errors.has('username') }" name="username" type="text">
              <span v-show="errors.has('username')" class="help is-danger">{{ errors.first('username') }}</span>
            </p>
          </div>
          <div class='input-container'>
            <label class="label" for="email">Email</label>
            <p :class="{ 'control': true }">
              <input id="email" placeholder='someone@email.com' v-model="account.Email" v-validate="'required|email'" :class="{'input': true, 'is-danger': errors.has('email') }" name="email" type="text">
              <span v-show="errors.has('email')" class="help is-danger">{{ errors.first('email') }}</span>
            </p>
          </div>
          <div class='input-container'>
            <label class="label" for="password">Password</label>
            <p :class="{ 'control': true }">
              <input id="password" placeholder='Password' v-model="account.Password" v-validate="'required|min:8|max:256'" :class="{'input': true, 'is-danger': errors.has('password') }" name="password" type="password">
              <span v-show="errors.has('password')" class="help is-danger">{{ errors.first('password') }}</span>
            </p>
          </div>
        </section>
        <section class='message-container'>
          <span class='message' v-bind:class="{success: success}"> {{ message }} </span>
        </section>
        <div class='menu-bar'>
          <button class='menu-button' @click="Create" :disabled="!valid" v-on:click="$root.loading = true"> Create </button>
          <router-link to='/' v-on:click="$root.loading = true"> Cancel </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import router from '@/router';
import { CreateAccount } from '@/database';
import LoginHeader from '@/components/elements/LoginHeader';

export default {
  components: {
    'LoginHeader': LoginHeader
  },
  data () {
    return {
      account: {
        Username: '',
        Email: '',
        Password: '',
        Name: ''
      },
      valid: false,
      message: undefined,
      success: false
    };
  },
  updated () {
    // TODO: Username validation (no spaces)
    // TODO: Password validation (length, characters)
    // TODO: Shouldn't send "ConfirmPassword" to the API since it isn't needed
    this.valid =
      this.account.Username !== '' &&
      this.account.Email !== '' &&
      this.account.Password !== '';
  },
  methods: {
    Create: function () {
      CreateAccount(this.account)
        .then((response) => {
          if (response.status === 201) {
            this.message = 'Account successfully created. Redirecting to login page...';
            this.success = true;
            this.$root.loading = false;
            this.Redirect();
          } else {
            response.text().then((text) => { this.message = text; });
            this.success = false;
            this.$root.loading = false;
          }
        });
    },
    Redirect: function () {
      setTimeout(() => {
        router.push('/login');
      }, 3000);
    },
    enterKeyPressHandler: function () {
      if (this.valid) {
        this.Create();
      }
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
.input-container {
  margin-top: 35px;
}
.label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
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
.message-container {
  margin-top: 25px;
  margin-left: 25px;
}
.message {
  display: inline-block;
  font-style: italic;
  font-size: 14px;
  color: #ff4949;
}
.success {
  color: #4e9155;
}

.help.is-danger {
  display: inline-block;
  font-style: italic;
  font-size: 14px;
  color: #ff4949;
  position: absolute;
}
</style>

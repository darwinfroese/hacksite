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
            <ValidatedTextBox 
            :fieldLabel = "'Username *'" 
            :placeholder = "'Username'" 
            :model= 'account.Username' 
            :fieldName = "'Username'"
            :validateExpression = "'required|alpha_num|min:3'" />
          </div>
          <div class='input-container'>
            <ValidatedTextBox 
            :fieldLabel = "'Email *'" 
            :placeholder = "'someone@email.com'" 
            :model= 'account.Email' 
            :fieldName = "'email'"
            :validateExpression = "'required|email'" />       
          </div>
          <div class='input-container'>
            <ValidatedTextBox 
            :fieldLabel = "'Password *'" 
            :placeholder = "'Password'" 
            :model= 'account.Password' 
            :fieldName = "'password'"
            :validateExpression = "'required'" />
          </div>
          <div class='input-container'>
            <ValidatedTextBox 
            :fieldLabel = "'Confirm your password *'" 
            :placeholder = "'Confirm Password'" 
            :model= 'account.ConfirmPassword' 
            :fieldName = "'ConfirmPassword'"
            :validateExpression = "'required|confirmed:password'" />
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
import ValidatedTextBox from '@/components/elements/ValidatedTextBox';

export default {
  components: {
    'LoginHeader': LoginHeader,
    'ValidatedTextBox': ValidatedTextBox
  },
  data () {
    return {
      account: {
        Username: '',
        Email: '',
        Password: '',
        ConfirmPassword: ''
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
      this.account.Password !== '' &&
      this.account.ConfirmPassword !== '' &&
      this.ConfirmPasswords();
  },
  methods: {
    Create: function () {
      CreateAccount(this.account)
        .then((response) => {
          if (response.status === 201) {
            this.message = 'Account successfully created. Redirecting to login page...';
            this.success = true;
            this.Redirect();
          }
          return response.json();
        }).then((data) => {
          if (data.ErrorMessage) {
            this.message = data.ErrorMessage;
            this.success = false;
          }
        });
    },
    ConfirmPasswords: function () {
      return this.account.Password === this.account.ConfirmPassword;
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
</style>

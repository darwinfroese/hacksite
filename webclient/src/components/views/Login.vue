<template>
  <div>
    <LoginHeader />
    <!-- TODO: All of these inputs should accept 'enter' as the default action -->
    <div class='container'>
      <div class='card'>
        <div class='input-container'>
          <ValidatedTextBox 
          :fieldLabel = "'Username *'" 
          :placeholder = "'Username'" 
          :model= 'account.Username' 
          :fieldName = "'Username'"
          :validateExpression = "'required'" />
        </div>
        <div class='input-container'>
          <ValidatedTextBox 
          :fieldLabel = "'Password *'" 
          :placeholder = "'Password'" 
          :model= 'account.Password' 
          :fieldName = "'Password'"
          :validateExpression = "'required'" />
        </div>
        <div class='message-container'>
          <span class='message' v-bind:class="{success: success}"> {{ message }} </span>
        </div>
        <div class='create-account'>
          <span class='helper-text'> Don't have an account? </span>
          <router-link to='/createaccount'> Create one now. </router-link>
        </div>
        <div class='menu-bar'>
          <button class='menu-button' @click="Login" :disabled="!valid"> Login </button>
          <router-link to='/'> Cancel </router-link>          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import router from '@/router';
import { Login } from '@/database';
import LoginHeader from '@/components/elements/LoginHeader.vue';
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
        Password: ''
      },
      valid: false,
      message: '',
      success: false
    };
  },
  methods: {
    Login: function () {
      Login(this.account)
        .then((response) => {
          if (response.status === 200) {
            this.message = 'Successfully logged in. Redirecting...';
            this.success = true;
            this.Redirect();
          } else if (response.status === 401) {
            this.message = 'Incorrect username or password';
          } else {
            this.message = 'Something went wrong, please try again';
          }
        }, (error) => {
          console.log('Error: ', error);
        });
    },
    Validate: function () {
      this.valid = this.account.Username !== '' && this.account.Password !== '';
    },
    Redirect: function () {
      setTimeout(() => {
        router.push('/');
      }, 3000);
    }
  },
  updated () {
    this.Validate();
  }
};
</script>

<style scoped>
/* TODO: Styles used in multiple components should probably be exported */
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
  margin-top: 25px;
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
.create-account {
  padding: 10px;
  font-size: 14px;
  font-style: italic;
}
.helper-text {
  font-style: normal;
}
</style>

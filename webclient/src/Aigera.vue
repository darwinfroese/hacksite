<template>
  <div id="app">
    <div class="full-screen-modal" v-if="$root.loading" @click.stop.prevent>
      <vue-simple-spinner class="simple-spinner" line-fg-color="#325778"></vue-simple-spinner>
    </div>
    <div class="primary inverted header">
        <h1>AIGERA</h1>
        <h4 class="button" v-if="loggedIn" @click="ViewAccount"><i class="fas fa-user-circle"></i> Username</h4>
        <h4 class="button" v-if="loggedIn" @click="Logout">Logout</h4>
    </div>
    <div class="body-container">
      <router-view></router-view>
      <InfoFooter />
    </div>
  </div>
</template>

<script>
import InfoFooter from '@/components/elements/InfoFooter';
import Spinner from 'vue-simple-spinner';
import router from '@/router';
import { Logout } from '@/database';

export default {
  name: 'app',
  components: {
    InfoFooter: InfoFooter,
    'vue-simple-spinner': Spinner
  },
  methods: {
    LogOut: function () {
      Logout()
        .then((resp) => {
          router.push('/login');
        });
    },
    ViewAccount: function () {
      // TODO: Navigate to login page
    }
  }
};
</script>

<style>
* {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
#app {
  min-height: 100%;
  margin: 0;
  position: relative;
}
a {
  text-decoration: none;
}
hr {
  border: none;
  border-bottom: 1px solid #000;
}
.full-screen-modal {
  height: 100%;
  position: fixed;
  width: 100%;
  top: 0;
  background: rgba(0, 0, 0, 0.25);
  z-index: 2;
}
.full-screen-modal .simple-spinner {
  height: 100%;
  position: fixed;
  width: 100%;
  top: 50%;
}
.body-container {
  margin: 0 10%;
}
</style>

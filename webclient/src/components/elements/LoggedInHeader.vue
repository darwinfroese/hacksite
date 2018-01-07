<template>
  <div class='header'>
    <div class='bar'>
      <span class='spacer'> &nbsp; </span>
      <span class='title'>
        <span @click="navHome"> Aigera </span>
      </span>
      <span class='account'>
        <span @click="logout" class='logout-button' v-on:click="$root.loading = true">
          Logout
          <i class='fa fa-sign-out icon'></i>
        </span>
      </span>
    </div>
  </div>
</template>

<script>
import { Logout } from '@/database';
import router from '@/router';

export default {
  methods: {
    logout: function () {
      Logout()
        .then((resp) => {
          router.push('/login');
        });
    },
    navHome: function () {
      router.push('/');
    }
  }
};
</script>

<style scoped>
.bar {
  display: flex;
  align-items: center;
}
.spacer,
.account {
  flex-grow: 1;
}
.title {
  flex-grow: 2;
}
.header {
  display: block;
  margin: 0;
  padding: 30px;
  text-align: center;
  color: #fff;
  background-color: #325778;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
}
.title > span {
  font-size: 24px;
  margin: 0;
}
.title > span:hover {
  cursor: pointer;
}
.logout-button:hover {
  cursor: pointer;
}
.icon {
  padding-left: 15px;
}
</style>

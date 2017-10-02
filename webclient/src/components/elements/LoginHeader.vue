<template>
  <div class='header'>
    <div class='bar'>
      <h3 @click="navHome"> Hacksite </h3>
    </div>
  </div>
</template>

<script>
import router from '@/router';

export default {
  methods: {
    navHome: function () {
      router.push('/');
    }
  }
};
</script>

<style scoped>
.header {
  display: block;
  margin: 0;
  padding: 30px;
  text-align: center;
  color: #fff;
  background-color: #325778;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
h3 {
  font-size: 24px;
  margin: 0;
}
h3:hover {
  cursor: pointer;
}
</style>

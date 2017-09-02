<template>
  <div id="app">
    <router-view></router-view>
    <Info />
  </div>
</template>

<script>
import Info from '@/components/Info';

export default {
  name: 'app',
  components: {
    'Info': Info
  }
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  min-height: 100%;
}
a {
  text-decoration: none;
}
</style>

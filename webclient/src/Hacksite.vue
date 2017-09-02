<template>
  <div id="app">
    <HeaderBar />
    <router-view></router-view>
    <Info />
  </div>
</template>

<script>
import Info from '@/components/Info';
import HeaderBar from '@/components/HeaderBar';

export default {
  name: 'app',
  components: {
    'Info': Info,
    'HeaderBar': HeaderBar
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
  margin: 0;
  position: relative
}
a {
  text-decoration: none;
}
hr {
  border: none;
  border-bottom: 1px solid #000;
}
</style>

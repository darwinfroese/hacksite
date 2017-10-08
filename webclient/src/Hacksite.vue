<template>
  <div id="app">
    <router-view></router-view>
    <InfoFooter />
  </div>
</template>

<script>
import InfoFooter from '@/components/elements/InfoFooter';

export default {
  name: 'app',
  components: {
    'InfoFooter': InfoFooter
  }
};
</script>

<style>
* {
  font-family: 'Roboto', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
#app {
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

<template>
  <div id="app">
    <TitleHeader />
    <router-view></router-view>
    <InfoFooter />
  </div>
</template>

<script>
import InfoFooter from '@/components/elements/InfoFooter';
import TitleHeader from '@/components/elements/TitleHeader';

export default {
  name: 'app',
  components: {
    'InfoFooter': InfoFooter,
    'TitleHeader': TitleHeader
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

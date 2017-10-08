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
.error {
    color : red;
    font-size : 15px;
}
input, textarea {
  display: block;
  font-size: 16px;
  padding: 5px 10px;
  width: 500px;
  border: none;
  font-weight: 100;
  outline: none;
  color: #325778;
}
input:focus {
  border-bottom: 1px solid #325778;
}
input {
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
  margin-top: 10px;
  transition: all 0.5s linear;
}
.input-label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
}
</style>

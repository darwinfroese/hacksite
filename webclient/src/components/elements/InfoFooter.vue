<template>
  <div class='bottom'>
    <p><router-link to='/releasenotes'>View the Release Notes</router-link></p>
    <p> Have an problem? <a href="https://github.com/darwinfroese/hacksite/issues"> Submit an issue </a> </p>
  </div>
</template>

<style scoped>
p {
  font-size: 12px; 
  font-style: italic;
  margin: 0;
}
.bottom {
  margin-top: 30px;
  text-align: center;
  width: auto;
  display: block;
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
</style>

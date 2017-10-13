<template>
  <div>
    <LoginHeader />
    <div class='container'>
      <h3> Hacksite Beta (v0.5) - Release Notes </h3>
      <hr />
      <div>
        <ul>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/14'>#14</a> - Moved to using DynamoDB </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/13'>#13</a> - Implemented middleware for logging and applying CORS headers </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/1'>#1</a> - Accounts need to have an unique username and password </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/19'>#19</a> - Deleting a project prompts for confirmation </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/10'>#10</a> - Switched from console output logging to file logging </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/3'>#3</a> - Setup SSL for security </li>
          <li>
            General Quality of Development Improvements
            <ul>
              <li> <a href='https://github.com/darwinfroese/hacksite/issues/2'>#2</a> - Build Script Fixes </li>
              <li> <a href='https://github.com/darwinfroese/hacksite/issues/24'>#24</a> - Build Script Fixes </li>
              <li> <a href='https://github.com/darwinfroese/hacksite/issues/26'>#26</a> - Continuous Integration and Continuous Delivery </li>
              <li> <a href='https://github.com/darwinfroese/hacksite/issues/9'>#9</a> - Unit tests for server code </li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import LoginHeader from '@/components/elements/LoginHeader';

export default {
  components: {
    'LoginHeader': LoginHeader
  }
};
</script>

<style scoped>
.container {
  margin: 50px;
}
a, a:visited {
  color: #325778;
}
a:hover {
  color: #1B3F60;
}
</style>

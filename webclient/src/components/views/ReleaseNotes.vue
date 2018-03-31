<template>
  <div>
    <div class='container'>
      <h3> Hacksite Beta (v0.5) - Release Notes </h3>
      <hr />
      <div>
        <h4> New Features </h4>
        <ul>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/1'>#1</a> - Accounts need to have an unique username and password </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/3'>#3</a> - Setup SSL for security </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/10'>#10</a> - Switched from console output logging to file logging </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/12'>#12</a> - Wait indicators are displayed an all requests </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/13'>#13</a> - Implemented middleware for logging and applying CORS headers </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/14'>#14</a> - Moved to using DynamoDB </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/15'>#15</a> - Implemented remember me that will prevent logging the user out for a long time </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/19'>#19</a> - Deleting a project prompts for confirmation </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/41'>#41</a> - Add and Update accounts API endpoints, case insensitive usernames </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/51'>#51</a> - Release notes are displayed in the webclient </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/67'>#67</a> - Added a name field to account model </li>
        </ul>
        <h4> Bug Fixes </h4>
        <ul>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/18'>#18</a> - Completing an iteration selects the next iteration if one exists </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/45'>#45</a> - Enter key triggers the primary action button on forms </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/55'>#55</a> - Project ID is no longer displayed on the project details page </li>
        </ul>
        <h4> General Quality of Development Changes </h4>
        <ul>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/2'>#2</a> - Build Script Fixes </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/9'>#9</a> - Unit tests for server code </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/24'>#24</a> - Build Script Fixes </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/26'>#26</a> - Continuous Integration and Continuous Delivery </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/63'>#63</a> - Cleaned up the organization of the models package </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/64'>#64</a> - Added a POST method to Login; will deprecate GET </li>
          <li> <a href='https://github.com/darwinfroese/hacksite/issues/80'>#80</a> - ResponseWriter wrapper for future usage </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>

export default {
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

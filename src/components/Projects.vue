<template>
  <div>
    <h2> Projects </h2>
    <hr />
    <div>
      <router-link to='/create'>Add a project</router-link>
    </div>
    <ul>
      <li v-for="project in projects" v-bind:key="project.id">
        <router-link :to="project.details"> {{ project.name}} </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import database from '@/database';

export default {
  data () {
    return {
      projects: database.GetProjects()
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
  text-align: center;
}
ul {
  text-decoration: none;
  list-style: none;
}
a {
  width: 90%;
  display: block;
}
</style>

<template>
  <div>
    <h3> Projects </h3>
    <div>
      <router-link to='/create'>Add a project</router-link>
    </div>
    <ul>
      <Project v-for="project in projects" :project="project" :key="project.ID" v-on:update="Update" />
    </ul>
  </div>
</template>

<script>
import database from '@/database';
import Project from '@/components/Project';

export default {
  components: {
    'Project': Project
  },
  data () {
    return {
      projects: []
    };
  },
  methods: {
    Update: function () {
      let promise = database.GetProjects();
      promise.then((response) => {
        return response.json();
      }).then((json) => {
        this.projects = json;
      });
    }
  },
  mounted () {
    this.Update();
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
.project {
  width: 90%;
  position: relative;
  margin: 0;
  display: inline-block;
}
.removeButton {
  display: none;
}
.removeButton:hover {
  cursor: pointer;
}
</style>

<template>
  <div>
    <h2> Projects </h2>
    <hr />
    <div>
      <router-link to='/create'>Add a project</router-link>
    </div>
    <ul>
      <Project v-for="project in projects" :project="project" :key="project.id" v-on:update="Update" />
    </ul>
  </div>
</template>

<script>
import database from '@/database';
import router from '@/router';
import Project from '@/components/Project';

export default {
  components: {
    'Project': Project
  },
  data () {
    return {
      projects: database.GetProjects()
    };
  },
  methods: {
    Update: function () {
      this.projects = database.GetProjects();
    },
    RemoveProject: function (id) {
      database.RemoveProject(id);
    },
    ViewProject: function (id) {
      router.push('/details/' + id);
    }
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

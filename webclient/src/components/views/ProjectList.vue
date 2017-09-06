<template>
  <div class='list'>
    <h3> Projects </h3>
    <div class='menu-bar'>
        <router-link class='menu-button' to='/create'>
          <i class='fa fa-plus'></i>
          Add a project
        </router-link>
      </span>
    </div>
    <div class='list-container'>
      <Project v-for="project in projects" :project="project" :key="project.ID" v-on:update="Update" />
    </div>
  </div>
</template>

<script>
import database from '@/database';
import Project from '@/components/elements/Project';

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
  display: inline-block;
}
.list {
  margin: 50px;
}
.list-container {
  margin: 10px;
}
.menu-bar {
  display: flex;
  flex-direction: row-reverse;
  margin: 0 5%;
}
.menu-button {
  background-color: #3764ad;
  padding: 7px 10px;
  line-height: 16px;
  font-size: 16px;
}
.menu-button > i {
  padding-right: 5px;
  width: 16px;
  height: 16px;
}
.menu-button:hover {
  background-color: #25467c;
}
.menu-button:visited {
  color: #fff;
}
.removeButton {
  display: none;
}
.removeButton:hover {
  cursor: pointer;
}
</style>

<template>
  <div>
    <!-- Log in should be one "page" and logged in should be another -->
    <LoggedInHeader />
    <div class='list' v-if="loading">
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
    <div v-if="!loading">
      <vue-simple-spinner></vue-simple-spinner>
    </div>
  </div>
</template>

<script>
import { GetProjects } from '@/database';
import Project from '@/components/elements/Project';
import LoggedInHeader from '@/components/elements/LoggedInHeader';
import VueSimpleSpinner from 'vue-simple-spinner/src/components/Spinner';

export default {
  components: {
    VueSimpleSpinner,
    'Project': Project,
    'LoggedInHeader': LoggedInHeader
  },
  data () {
    return {
      projects: [],
      loading: false
    };
  },
  methods: {
    Update: function () {
      let self = this;
      let promise = GetProjects();
      promise.then((response) => {
        return response.json();
      }).then((json) => {
        this.projects = json;
        self.loading = true;
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
  color: #fff;
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
  margin: 25px 5%;
}
.menu-button {
  background-color: #529A7F;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.menu-button > i {
  padding-right: 5px;
  width: 16px;
  height: 16px;
}
.menu-button:hover {
  background-color: #176548;
}
.menu-button:visited {
  color: #fff;
}
</style>

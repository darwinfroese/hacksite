<template>
  <div class='details-container'>
    <div class='menu-bar'>
      <router-link class='menu-button' to='/'>
        <i class='fa fa-chevron-left'></i>
        Back to Projects
      </router-link>
    </div>
    <br>
    <div>
      <div> [ {{ project.ID }} ] {{ project.Name }} </div>
      <span class='description'>
        {{ project.Description }}
      </span>
      <div class='tasks'>
        <Task v-for="task in project.Tasks" v-bind:key='task.ID' :task="task" :pid="project.ID" v-on:GetProject="GetProject" />
      </div>
    </div>
  </div>
</template>

<script>
import database from '@/database';
import Task from '@/components/elements/Task';

export default {
  components: {
    'Task': Task
  },
  data () {
    return {
      project: {}
    };
  },
  props: ['pid'],
  methods: {
    GetProject: function () {
      database.GetProject(this.pid)
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        this.project = json;
      });
    }
  },
  mounted () {
    this.GetProject();
  }
};
</script>

<style scoped>
.details-container {
  margin: 50px;
}
.description {
  margin-left: 30px;
  font-style: italic;
}
.tasks {
  margin: 25px;
}
.checkbox {
  display: inline;
  margin: 0;
}
label:hover {
  cursor: pointer;
}
.menu-bar {
  display: flex;
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


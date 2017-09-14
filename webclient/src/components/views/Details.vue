<template>
  <div class='details-container'>
    <div class='menu-bar'>
      <router-link class='menu-button' to='/'>
        <i class='fa fa-chevron-left'></i>
        Back to Projects
      </router-link>
      <router-link class='menu-button' :to="editRoute">
        <i class='fa fa-pencil'></i>
        Edit Project
      </router-link>
      <router-link class='menu-button' :to="iterationRoute">
        <i class='fa fa-plus'></i>
        Add Iteration
      </router-link>
      <router-link class='menu-button' :to="allIterations">
        <i class='fa fa-history'></i>
        View iterations
      </router-link>
    </div>
    <br>
    <div class='detail-card'>
      <div class='project-header'>
        [ {{ project.ID }} ] {{ project.Name }}
        <span class='iteration' v-if="project.Iteration" title='Current Iteration'>
           ( Iteration {{ project.Iteration.Number }} )
        </span>
      </div>
      <div class='description'>
        {{ project.Description }}
      </div>
      <div class='tasks'>
        <Task v-for="task in project.Tasks" v-bind:key='task.ID' :task="task" :pid="project.ID" v-on:GetProject="Update" />
      </div>
    </div>
  </div>
</template>

<script>
import { GetProject } from '@/database';
import Task from '@/components/elements/Task';

export default {
  components: {
    'Task': Task
  },
  computed: {
    editRoute: function () {
      return '/edit/' + this.project.ID;
    },
    iterationRoute: function () {
      return '/iteration/' + this.project.ID;
    },
    allIterations: function () {
      return '/iterations/' + this.project.ID;
    }
  },
  data () {
    return {
      project: {}
    };
  },
  props: ['pid'],
  methods: {
    Update: function () {
      GetProject(this.pid)
        .then((response) => {
          return response.json();
        })
        .then((json) => {
          this.project = json;
        });
    }
  },
  mounted () {
    this.Update();
  }
};
</script>

<style scoped>
a {
  color: #fff;
}
.details-container {
  margin: 50px;
}
.detail-card {
  padding: 25px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.description {
  margin-left: 30px;
  margin-top: 20px;
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
  margin: 0 10px;
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
.project-header {
  font-size: 18px;
  font-weight: 500;
}
.iteration {
  margin-left: 25px;
  font-size: 14px;
}
</style>


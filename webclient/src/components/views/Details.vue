<template>
  <div>
    <router-link to='/'>Back to Projects</router-link>
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
</style>


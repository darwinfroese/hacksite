<template>
  <div>
    <router-link to='/'>Back to Projects</router-link>
    <br>
    <div>
      <div> [ {{ project.id }} ] {{ project.name }} </div>
      <span class='description'>
        {{ project.description }}
      </span>
      <div class='tasks'>
        <Task v-for="task in project.tasks" v-bind:key='task.id' :task="task" :pid="project.id" />
      </div>
    </div>
  </div>
</template>

<script>
import database from '@/database';
import Task from '@/components/Task';

export default {
  components: {
    'Task': Task
  },
  data () {
    return {
      project: database.GetProject(this.pid)
    };
  },
  props: ['pid']
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


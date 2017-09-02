<template>
  <div>
    <router-link to='/'>Back to Projects</router-link>
    <br>
    <div>
      <div> [ {{ project.id }} ] {{ project.name }} </div>
      <span class='description'>
        {{ project.description }}
      </span>
      <ul>
        <li v-for="task in project.tasks" v-bind:key='task.id' class='task'>
          <label>
            <input type='checkbox' class='checkbox' v-model="task.completed" @click="Update(task)">
            <span v-bind:class="{ completed: task.completed }"> {{task.task}} </span>
          </label>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import database from '@/database';

export default {
  data () {
    return {
      project: database.GetProject(this.pid)
    };
  },
  props: ['pid'],
  methods: {
    Update: function (task) {
      database.UpdateTask(this.pid, task);
    }
  }
};
</script>

<style scoped>
.description {
  margin-left: 30px;
  font-style: italic;
}
.completed {
  text-decoration: line-through;
  color: slategray;
}
.task {
  display: block;
  font-size: 18px;
  line-height: 18px;
}
.checkbox {
  display: inline;
  margin: 0;
}
label:hover {
  cursor: pointer;
}
</style>


<template>
  <div class='task'>
    <label class='taskItem'>
      <input type='checkbox' class='checkbox' v-model="task.completed" @click="Update()">
      <span v-bind:class="{ completed: task.completed }"> {{ task.task }} </span>
    </label>
    <span class='removeButton' @click="RemoveTask()"> x </span>
  </div>
</template>

<script>
import database from '@/database';

export default {
  props: ['task', 'pid'],
  methods: {
    Update: function () {
      database.UpdateTask(this.pid, this.task);
    },
    RemoveTask: function () {
      database.RemoveTask(this.pid, this.task);
    }
  }
};
</script>

<style scoped>
.task {
  display: block;
  font-size: 18px;
  line-height: 18px;
  width: 90%;
  position: relative;
}
.taskItem:hover {
  cursor: pointer;
}
.completed {
  text-decoration: line-through;
  color: slategray;
}
.removeButton {
  position: absolute;
  right: 10px;
  display: none;
}
.removeButton:hover {
  cursor: pointer;
}
.task:hover .removeButton {
  display: inline;
}
</style>

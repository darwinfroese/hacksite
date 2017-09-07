<template>
  <div class='task'>
    <label class='taskItem'>
      <input type='checkbox' class='checkbox' v-model="task.Completed" @click="Update()">
      <span v-bind:class="{ completed: task.Completed }"> {{ task.Task }} </span>
    </label>
    <span class='removeButton' @click="RemoveTask()"> <i class='fa fa-times'></i> </span>
  </div>
</template>

<script>
import database from '@/database';

export default {
  props: ['task', 'pid'],
  methods: {
    Update: function () {
      database.UpdateTask(this.task);
    },
    RemoveTask: function () {
      database.RemoveTask(this.task)
        .then((response) => {
          this.$emit('GetProject');
        });
    }
  }
};
</script>

<style scoped>
input {
  display: inline-block;
  width: 16px;
  height: 16px;
}
.task {
  display: block;
  font-size: 16px;
  line-height: 16px;
  width: 90%;
  position: relative;
  padding: 2px;
}
.taskItem:hover {
  cursor: pointer;
}
.completed {
  text-decoration: line-through;
  font-style: italic;
  color: slategray;
}
.removeButton {
  display: none;
  width: 16px;
  height: 16px;
  line-height: 16px;
  vertical-align: middle;
  float: right;
  margin: auto;
}
.removeButton:hover {
  cursor: pointer;
}
.removeButton > i {
  height: 16px; 
  width: 16px;
}
.task:hover .removeButton {
  display: inline;
}
</style>

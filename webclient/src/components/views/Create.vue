<template>
  <div>
    <span>Create a new project</span>
    <br>
    <input type='text' placeholder='Project Name' v-model='project.Name'>
    <input type='text' placeholder='Project Description' v-model='project.Description'>
    <section id='tasks'>
      <input type='text' placeholder='Task Description' name='taskInput'>
    </section>
    <p v-if="warningDisplayed"> Task Limit Reached! </p>
    <button name='addTaskButton' id='addTaskButton' @click="AddTask()"> Add Task </button>
    <hr />
    <div>
      <button @click="SaveProject()"> Save Project </button>
      <router-link to='/'> Cancel </router-link>
    </div>
  </div>
</template>

<script>
import database from '@/database';
import router from '@/router';

export default {
  data () {
    return {
      project: {
        Name: '',
        Description: '',
        Tasks: [],
        Completed: false
      },
      taskCount: 1,
      warningDisplayed: false
    };
  },
  methods: {
    SaveProject: function () {
      let inputs = document.getElementsByName('taskInput');

      inputs.forEach((i, idx) => {
        this.project.Tasks.push({'task': i.value, 'id': idx, 'completed': false});
      });

      database.AddProject(this.project)
        .then(() => {
          console.log('project added.');
          router.push('/');
        });
    },
    AddTask: function () {
      if (this.taskCount === 4) {
        document.getElementById('addTaskButton').disabled = true;
        this.warningDisplayed = true;
        return;
      }

      let node = document.createElement('input');
      node.type = 'text';
      node.placeholder = 'Task Description';
      node.name = 'taskInput';
      document.getElementById('tasks').appendChild(node);
      this.taskCount++;
    }
  }
};
</script>

<style>
  input {
    display: block;
  }
</style>

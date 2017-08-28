<template>
  <div>
    <span>Create a new project</span>
    <br>
    <input type='text' placeholder='Project Name' v-model='project.name'>
    <input type='text' placeholder='Project Description' v-model='project.description'>
    <section id='tasks'>
      <input type='text' placeholder='Task Description' name='taskInput'>
    </section>
    <button name='addTaskButton' id='addTaskButton' @click="AddTask"> Add Task </button>
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
        name: '',
        description: '',
        tasks: []
      },
      taskCount: 1,
      warningDisplayed: false
    };
  },
  methods: {
    SaveProject: function () {
      let inputs = document.getElementsByName('taskInput');

      inputs.forEach((i) => {
        this.project.tasks.push(i.value);
      });

      database().AddProject(this.project);
      router.push('/');
    },
    AddTask: function () {
      if (this.taskCount === 4) {
        let node = document.createElement('p');
        node.appendChild(document.createTextNode('Task limit reached'));
        document.getElementById('tasks').appendChild(node);
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

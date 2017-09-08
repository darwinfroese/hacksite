<template>
  <div class='create-container'>
    <div class='create-card'>
      <h4>Create a new project</h4>
      <div class='create-form'>
        <div class='input'>
          <span class='input-label'> Project Name </span>
          <input type='text' placeholder='Project Name' v-model='project.Name'>
        </div>
        <div class='input'>
          <span class='input-label'> Project Description </span>
          <textarea placeholder='Enter your project description' v-model='project.Description'></textarea>
        </div>
        <section id='tasks' class='tasks input'>
          <span class='input-label'>
            Project Tasks
            <i class='fa fa-info-circle info'>
              <span class='info-text'>
                Hacksite only lets you select four tasks for your projects to keep them
                small and achievable!
              </span>
            </i>
          </span>
          <input type='text' placeholder='Task 1 Description' name='taskInput'>
          <input type='text' placeholder='Task 2 Description' name='taskInput'>
          <input type='text' placeholder='Task 3 Description' name='taskInput'>
          <input type='text' placeholder='Task 4 Description' name='taskInput'>
        </section>
        <div class='menu-bar'>
          <button class='menu-button' @click="SaveProject()"> Add Project </button>
          <router-link to='/'> Cancel </router-link>
        </div>
      </div>
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
        let contents = i.value;

        if (/\S/.test(contents)) {
          this.project.Tasks.push({'task': contents, 'id': idx, 'completed': false});
        }
      });

      database.AddProject(this.project)
        .then(() => {
          console.log('project added.');
          router.push('/');
        });
    }
  }
};
</script>

<style>
input, textarea {
  display: block;
  font-size: 16px;
  padding: 2px 10px;
  width: 500px;
  border: none;
  font-weight: 100;
  outline: none;
  color: #325778;
}
textarea {
  margin: 0;
  height: 75px;
  width: 500px!important;
  border: 1px solid #eee;
  resize: none;
  overflow: auto;
}
h4 {
  margin-left: 15px;
}
a, a:visited {
  color: #325778;
}
a:hover {
  color: #1B3F60;
}
button, a {
  border-radius: 0;
  border: none;
}
button:hover {
  cursor: pointer;
}
.input {
  margin-top: 20px;
}
.info {
  position: relative;
}
.info .info-text {
  visibility: hidden;
  width: 400px;
  background-color: #325778;
  color: #fff;
  text-align: center;
  padding: 10px;
  border-radius: 6px;
  font-size: 12px;
  font-style: italic;

  /* Position the tooltip text - see examples below! */
  position: absolute;
  left: 25px;
  z-index: 1;
}
.info:hover .info-text{
  visibility: visible;
}
.tasks {
  margin-top: 10px;
}
.tasks > input {
  margin: 10px 0;
}
.input-label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
}
.create-form {
  padding-left: 15px;
}
.create-container {
  margin: 50px;
}
.create-card {
  padding: 25px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.menu-bar {
  display: flex;
  margin-top: 25px;
  align-items: center;
  margin-left: 10px;
}
.menu-button {
  background-color: #529A7F;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  color: #fff;
  margin-right: 10px;
}
.menu-button:hover {
  background-color: #176548;
}
</style>

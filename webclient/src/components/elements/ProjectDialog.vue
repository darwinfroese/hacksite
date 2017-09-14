<template>
  <div class='card'>
    <h4>{{ title }}</h4>
    <div class='form'>
      <div class='input'>
        <span class='input-label'> Project Name *</span>
        <input type='text' placeholder='Project Name' v-model='project.Name'>
      </div>
      <div class='input'>
        <span class='input-label'> Project Description </span>
        <textarea placeholder='Enter your project description' v-model='project.Description'></textarea>
      </div>
      <TaskInputs />
      <section id='messageSection' class='messageSection'>
        <div class='infoMessage'> * indicates a required field </div>
      </section>
      <div class='menu-bar'>
        <button class='menu-button' @click="Handler" :disabled="!valid"> {{ buttonText }} </button>
        <router-link to='/'> Cancel </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import TaskInputs from '@/components/elements/TaskInputs';

export default {
  props: ['title', 'project', 'buttonText'],
  components: {
    'TaskInputs': TaskInputs
  },
  methods: {
    Handler: function () {
      this.$emit('Handle');
    },
    SetTaskValues: function () {
      let tasks = this.project.Tasks;

      tasks.forEach((task, idx) => {
        document.getElementById('taskInput' + (idx + 1)).value = task.Task;
      });
    },
    ValidateInput: function () {
      if (/\S/.test(this.project.Name)) {
        this.valid = true;
      } else {
        this.valid = false;
      }
    }
  },
  data () {
    return {
      valid: false
    };
  },
  mounted () {
    this.SetTaskValues();
    this.ValidateInput();
  },
  updated () {
    this.ValidateInput();
  }
};
</script>

<style scoped>
input, textarea {
  display: block;
  font-size: 16px;
  padding: 5px 10px;
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
  transition: all 0.5s linear;
}
textarea:focus {
  border: 1px solid #325778;
}
input {
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
  margin-top: 10px;
  transition: all 0.5s linear;
}
input:focus {
  border-bottom: 1px solid #325778;
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
.input-label {
  font-size: 14px;
  margin-left: 5px;
  font-weight: 500;
}
.form {
  padding-left: 15px;
}
.card {
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
.menu-button:disabled,
.menu-button:disabled:hover {
  background-color: #919191;
  box-shadow: none;
  cursor: not-allowed;
}
.menu-button:hover {
  background-color: #176548;
}
.messageSection {
  margin: 25px;
}
.infoMessage {
  font-style: italic;
  font-size: 14px;
}
</style>
<template>
  <div>
    <LoggedInHeader />
    <div class='container'>
      <div class='card'>
        <h4> Enter Evolution Information </h4>
        <div class='form'>
          <div class='field'>
            <span class='label'> Project Name </span>
            <span class='value'> {{ this.project.Name }} </span>
          </div>
          <div class='field'>
            <span class='label'> Project Description </span>
            <span class='value'> {{ this.project.Description }} </span>
          </div>
          <TaskInputs />
          <div class='menu-bar'>
            <button class='menu-button' @click="AddEvolution"> Start Evolution </button>
            <router-link to='/'> Cancel </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import router from '@/router';
import { GetProject, AddEvolution } from '@/database';
import TaskInputs from '@/components/elements/TaskInputs';
import LoggedInHeader from '@/components/elements/LoggedInHeader';

export default {
  components: {
    'TaskInputs': TaskInputs,
    'LoggedInHeader': LoggedInHeader
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
    },
    GetTasks: function () {
      let inputs = document.getElementsByName('taskInput');
      let tasks = [];

      inputs.forEach((i, idx) => {
        let contents = i.value;

        if (/\S/.test(contents)) {
          if (idx >= tasks.length) {
            tasks.push({'task': contents, 'id': idx, 'completed': false, ProjectID: this.project.ID});
          }
        }

        this.project.CurrentEvolution.Tasks = tasks;
      });

      this.project.CurrentEvolution.Tasks = tasks;
    },
    AddEvolution: function () {
      this.GetTasks();
      let evolution = this.project.CurrentEvolution;
      evolution.Number++;

      AddEvolution(evolution)
        .then((response) => {
          return response.json();
        })
        .then((project) => {
          router.push('/details/' + project.ID);
        });
    }
  },
  mounted () {
    this.Update();
  }
};
</script>

<style scoped>
.container {
  margin: 50px;
}
.form {
  padding-left: 15px;
}
.card {
  padding: 25px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
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
.field {
  margin: 10px;
}
.label {
  display: inline-block;
  font-weight: 600;
  width: 150px;
}
.value {
  margin-left: 15px;
}
</style>

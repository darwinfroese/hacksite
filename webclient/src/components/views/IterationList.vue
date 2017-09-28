<template>
  <div>
    <LoggedInHeader />
    <div class='container'>
      <div class='menu-bar'>
        <router-link class='menu-button' :to="detailsRoute">
          <i class='fa fa-chevron-left'></i>
          Back to Project
        </router-link>
      </div>
      <br>
      <div class='card'>
        <div class='content'>
          <div class='field'>
            <span class='label'> Project Name </span>
            <span class='value'> {{ project.Name }} </span>
          </div>
        </div>
        <div class='iteration-container' v-if="project.Iterations">
          <section v-for="iteration in project.Iterations" :key="iteration.Number">
            <div class='iteration-label'> Iteration {{ iteration.Number }} </div>
            <div v-for="task in iteration.Tasks" :key="task.ID" class='iteration' v-bind:class="{ completed: task.Completed }">
              {{ task.Task }}
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { GetProject } from '@/database';
import LoggedInHeader from '@/components/elements/LoggedInHeader';

export default {
  components: {
    'LoggedInHeader': LoggedInHeader
  },
  data () {
    return {
      project: {}
    };
  },
  computed: {
    detailsRoute: function () {
      return '/details/' + this.project.ID;
    }
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
.card {
  padding: 25px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.menu-bar {
  display: flex;
}
.menu-button {
  background-color: #529A7F;
  margin: 0 10px;
  padding: 10px;
  line-height: 16px;
  font-size: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.menu-button > i {
  padding-right: 5px;
  width: 16px;
  height: 16px;
}
.menu-button:hover {
  background-color: #176548;
}
.menu-button:visited {
  color: #fff;
}
.content {
  margin-top: 25px;
  padding-left: 15px;
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
.iteration-container {
  margin: 25px;
  margin-top: 35px;
}
.iteration-label {
  font-weight: 600;
}
.iteration {
  margin-left: 10px;
  padding: 2px;
  font-size: 16px;
  line-height: 16px;
}
.completed {
  font-style: italic;
  text-decoration: line-through;
  color: slategray;
}
</style>

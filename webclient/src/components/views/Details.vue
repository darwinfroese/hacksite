<template>
  <div>
    <LoggedInHeader />
    <div class='details-container'>
      <div class='menu-bar'>
        <router-link class='menu-button' to='/' v-on:click.native="$root.loading = true">
          <i class='fa fa-chevron-left'></i>
          Back to Projects
        </router-link>
        <router-link class='menu-button' :to="editRoute" v-on:click.native="$root.loading = true">
          <i class='fa fa-pencil'></i>
          Edit Project
        </router-link>
        <router-link class='menu-button' :to="evolutionRoute" v-on:click.native="$root.loading = true">
          <i class='fa fa-plus'></i>
          Add Evolution
        </router-link>
        <router-link class='menu-button' :to="allEvolutions" v-on:click.native="$root.loading = true">
          <i class='fa fa-history'></i>
          View evolutions
        </router-link>
      </div>
      <br>
      <div class='detail-card'>
        <div class='project-header'>
          {{ project.Name }}
          <span class='evolution' v-if="project.CurrentEvolution" title='Current Evolution'>
            ( Evolution {{ project.CurrentEvolution.Number }} )
          </span>
          <span class='swap-link' v-if="swappable" title='Swap Evolutions'>
            Swap Evolutions
            <select @change="SwapEvolutions" v-model="selectedEvolution">
              <option v-for="evolution in project.Evolutions" :key="evolution.Number" :value="evolution.Number">
                {{ evolution.Number }}
              </option>
            </select>
          </span>
        </div>
        <div class='description'>
          {{ project.Description }}
        </div>
        <div class='tasks' v-if="project.CurrentEvolution">
          <Task v-for="task in project.CurrentEvolution.Tasks" v-bind:key='task.ID' :task="task" :pid="project.ID" :pname="project.Name" v-on:GetProject="Update" v-on:click="$root.loading = true" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { GetProject, ChangeCurrentEvolution } from '@/database';
import Task from '@/components/elements/Task';
import LoggedInHeader from '@/components/elements/LoggedInHeader';

export default {
  components: {
    'Task': Task,
    'LoggedInHeader': LoggedInHeader
  },
  computed: {
    editRoute: function () {
      return '/edit/' + this.project.ID;
    },
    evolutionRoute: function () {
      return '/evolution/' + this.project.ID;
    },
    allEvolutions: function () {
      return '/evolutions/' + this.project.ID;
    },
    swappable: function () {
      return this.project.Evolutions.length > 1;
    }
  },
  data () {
    return {
      project: {
        Evolutions: []
      },
      selectedEvolution: {}
    };
  },
  props: ['pid'],
  methods: {
    Update: function () {
      GetProject(this.pid)
        .then((response) => {
          this.$root.loading = false;
          return response.json();
        })
        .then((json) => {
          this.project = json;
          this.selectedEvolution = this.project.CurrentEvolution.Number;
        }).catch(() => {
          this.$root.loading = false;
        });
    },
    SwapEvolutions: function () {
      let selected = this.project.Evolutions.filter((iter) => {
        return iter.Number === this.selectedEvolution;
      })[0];
      ChangeCurrentEvolution(selected)
        .then((response) => {
          return response.json();
        })
        .then((json) => {
          this.project = json;
          this.selectedEvolution = this.project.CurrentEvolution.Number;
        });
    }
  },
  mounted () {
    this.$root.loading = true;
    this.Update();
  }
};
</script>

<style scoped>
a {
  color: #fff;
}
.details-container {
  margin: 50px;
}
.detail-card {
  padding: 25px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
}
.description {
  margin-left: 30px;
  margin-top: 20px;
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
.project-header {
  font-size: 18px;
  font-weight: 500;
}
.evolution {
  margin-left: 25px;
  font-size: 14px;
}
.swap-link {
  margin-left: 25px;
  font-size: 14px;
}
</style>


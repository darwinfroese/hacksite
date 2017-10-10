<template>
  <div>
    <LoggedInHeader />
    <div class='create-container' v-if="loading">
      <ProjectDialog
        :title="'Create A New Project'"
        :buttonText="'Add Project'"
        :project="project"
        v-on:Handle="SaveProject" />
    </div>
    <div v-if="!loading">
      <vue-simple-spinner></vue-simple-spinner>
    </div>
  </div>
</template>

<script>
import { AddProject } from '@/database';
import router from '@/router';
import ProjectDialog from '@/components/elements/ProjectDialog';
import LoggedInHeader from '@/components/elements/LoggedInHeader';

export default {
  components: {
    'ProjectDialog': ProjectDialog,
    'LoggedInHeader': LoggedInHeader
  },
  data () {
    return {
      project: {
        Name: '',
        Description: '',
        CurrentIteration: {
          Tasks: []
        },
        Completed: false,
        loading: false
      },
      taskCount: 1,
      warningDisplayed: false
    };
  },
  methods: {
    SaveProject: function () {
      var self = this;
      let inputs = document.getElementsByName('taskInput');

      inputs.forEach((i, idx) => {
        let contents = i.value;

        if (/\S/.test(contents)) {
          this.project.CurrentIteration.Tasks.push({'task': contents, 'id': idx, 'completed': false});
        }
      });

      AddProject(this.project)
        .then((response) => {
          return response.json();
        })
        .then((project) => {
          router.push('/details/' + project.ID);
          self.loading = true;
        });
    }
  }
};
</script>

<style scoped>
.create-container {
  margin: 50px;
}
</style>

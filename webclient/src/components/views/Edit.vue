<template>
  <div class='edit-container'>
    <ProjectDialog
      v-if="hasProject"
      :title="editTitle"
      :buttonText="'Update'"
      :project="project"
      v-on:Handle="Update" />
  </div>
</template>

<script>
import { GetProject, UpdateProject } from '@/database';
import router from '@/router';
import ProjectDialog from '@/components/elements/ProjectDialog';

export default {
  components: {
    'ProjectDialog': ProjectDialog
  },
  props: ['pid'],
  data () {
    return {
      project: {},
      hasProject: false
    };
  },
  mounted () {
    this.GetProject();
  },
  methods: {
    Update: function () {
      let inputs = document.getElementsByName('taskInput');
      let tasks = this.project.Tasks;

      inputs.forEach((i, idx) => {
        let contents = i.value;

        if (/\S/.test(contents)) {
          if (idx >= tasks.length) {
            tasks.push({'task': contents, 'id': idx, 'completed': false, ProjectID: this.project.ID});
          } else {
            tasks[idx].Task = contents;
          }
        }
      });

      this.project.Tasks = tasks;

      UpdateProject(this.project)
        .then(() => {
          router.push('/details/' + this.project.ID);
        });
    },
    GetProject: function () {
      GetProject(this.pid)
      .then((response) => {
        return response.json();
      })
      .then((json) => {
        this.project = json;
        this.hasProject = true;
      });
    }
  },
  computed: {
    editTitle: function () {
      return 'Edit Project ' + this.project.ID;
    }
  }
};
</script>

<style scoped>
.edit-container {
  margin: 50px;
}
</style>
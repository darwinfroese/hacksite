<template>
  <div>
    <div class='edit-container'>
      <ProjectDialog
        v-if="hasProject"
        :title="editTitle"
        :buttonText="'Update'"
        :project="project"
        v-on:Handle="Update" />
    </div>
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
    this.$root.loading = true;
    this.GetProject();
  },
  methods: {
    Update: function () {
      let inputs = document.getElementsByName('taskInput');
      let tasks = this.project.CurrentEvolution.Tasks;

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

      this.project.CurrentEvolution.Tasks = tasks;

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
        this.$root.loading = false;
        this.project = json;
        this.hasProject = true;
      }).catch(() => {
        this.$root.loading = false;
      });
    }
  },
  computed: {
    editTitle: function () {
      return 'Edit Project';
    }
  }
};
</script>

<style scoped>
.edit-container {
  margin: 50px;
}
</style>

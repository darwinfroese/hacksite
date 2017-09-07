<template>
  <div class='project'>
    <span class='status-icon'>
      <i class='fa fa-check' aria-hidden='true' v-if="project.Completed"></i>
      <i class='fa fa-clock-o' aria-hidden='true' v-if="!project.Completed"></i>
    </span>
    <span class='name' @click="ShowProject()" v-bind:class="{completed: project.Completed}">
      {{ project.Name }}
    </span>
    <span class='description'> {{ project.Description }} </span>
    <span class='control-icon'>
      <i class='fa fa-eye' title='View Project' @click="ShowProject()"></i>
      <i class='fa fa-trash-o' title='Delete Project' @click="RemoveProject()"></i>
    </span>
  </div>
</template>

<script>
import router from '@/router';
import database from '@/database';

export default {
  data () {
    return {
      active: false
    };
  },
  props: ['project'],
  methods: {
    ToggleMenu: function () {
      this.active = !this.active;
    },
    ShowProject: function () {
      router.push(this.project.Details);
    },
    RemoveProject: function () {
      this.active = false;
      database.RemoveProject(this.project)
        .then((response) => {
          this.$emit('update');
        });
    }
  }
};
</script>

<style scoped>
i {
  width: 16px;
  height: 16px;
}
.project {
  display: flex;
  padding: 5px 50px;
  margin: 5px;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
}
.project:hover {
  box-shadow: 0 14px 28px rgba(0,0,0,0.25), 0 10px 10px rgba(0,0,0,0.22);
}
.completed {
  font-style: italic;
  font-weight: 600;
}
.name {
  margin-left: 5%;
  flex-grow: 1;
  max-width: 300px;
}
.name:hover {
  cursor: pointer;
}
.description {
  flex-grow: 1;
  margin: auto;
  color: #325778;
  text-align: left;
  font-size: 13px;
  font-style: italic;
  vertical-align: middle;
}
.span-icon {
  flex-grow: 1;
}
.control-icon {
  flex-grow: 1;
  text-align: right;
}
.control-icon > i {
  margin: 0 5px;
}
.control-icon > i:hover {
  cursor: pointer;
}
/* Icon (Color) Overrides */
.fa-check { 
  color: green;
}
.fa-clock-o {
  color: goldenrod;
}
</style>

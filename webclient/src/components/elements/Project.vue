<template>
  <div class='project'>
    <i class='fa fa-check' aria-hidden='true' v-if="project.Completed"></i>
    <i class='fa fa-clock-o' aria-hidden='true' v-if="!project.Completed"></i>
    <span class='name' @click="ShowProject()" v-bind:class="{completed: project.Completed}">
      {{ project.Name }}
    </span>
    <span class='dropdown'>
      <i class="fa fa-chevron-down" id='dropdown' @click="ToggleMenu()" v-if="!active"></i>
      <i class="fa fa-chevron-up" id='dropdown' @click="ToggleMenu()" v-if="active"></i>
    </span class='dropdown'>
    <div v-if="active" class='menu'>
      <i class='menu-item fa fa-eye' @click="ShowProject()"></i>
      <i class='menu-item fa fa-trash-o' @click="RemoveProject()"></i>
    </div>
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
  padding: 2px 50px;
}
.project:hover {
  background-color: #e5e6e8;
}
.completed {
  font-style: italic;
  font-weight: 600;
}
.name {
  margin-left: 5%;
}
.name:hover {
  cursor: pointer;
}
.menu {
  display: flex;
  flex-direction: row;
}
.menu-item {
  flex-grow: 1;
  text-align: center;
  padding: 10px 0;
}
.menu-item:hover {
  background-color: #eee;
  cursor: pointer;
}
.dropdown {
  float: right;
}
/* Icon (Color) Overrides */
.fa-check { 
  color: green;
}
.fa-clock-o {
  color: goldenrod;
}
</style>

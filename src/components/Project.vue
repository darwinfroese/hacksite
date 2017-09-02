<template>
  <div class='project'>
    <span class='name' @click="ShowProject()"> {{ project.name }} </span>
    <span v-if="project.completed" class='completed'> COMPLETED </span>
    <button class="dropdown" id='dropdown' @click="ToggleMenu()"> ... </button>
    <div v-if="active" class='menu'>
      <span class='menu-item' @click="ShowProject()"> View Project </span>
      <span class='menu-item' @click="RemoveProject()"> Delete Project </span>
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
      router.push(this.project.details);
    },
    RemoveProject: function () {
      this.active = false;
      database.RemoveProject(this.project.id);
      this.$emit('update');
    }
  }
};
</script>

<style scoped>
.project {
  width: 90%;
  position: relative;
}
.completed {
  margin-left: 5%;
  font-style: italic;
  font-weight: 600;
}
.name:hover {
  cursor: pointer;
}
.menu {
  right: 0;
  position: absolute;
  border: 1px solid #000;
  padding: 10px;
  background-color: #fff;
  z-index: 10;
}
.menu-item {
  display: block;
  margin: 0 -10px;
  padding: 0 10px;
}
.menu-item:hover {
  background-color: #eee;
}
.dropdown {
  float: right;
}
</style>

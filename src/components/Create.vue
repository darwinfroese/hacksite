<template>
  <div>
    <span>Create a new project</span>
    <br>
    <input type='text' placeholder='Project Name' v-model='name'>
    <hr />
    <div>
      <button @click="add()"> Add </button>
      <router-link to='/'> Cancel </router-link>
    </div>
  </div>
</template>

<script>
import database from '@/database';
import router from '@/router';

export default {
  data: () => ({
    name: ''
  }),
  methods: {
    add: function () {
      database().AddProject(this.name);
      router.push('/');
    }
  }
};
</script>

<style scoped>

</style>

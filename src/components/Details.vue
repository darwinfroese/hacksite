<template>
  <div>
    <router-link to='/'>Back to Projects</router-link>
    <br>
    <div>
      <span> [ {{ project.id }} ] {{ project.name }} </span>
    </div>
  </div>
</template>

<script>
import database from '@/database';

export default {
  data () {
    return {
      project: database().GetProject(this.pid)
    };
  },
  props: ['pid']
};
</script>

<style scoped>

</style>


// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import Aigera from './Aigera';
import router from './router';

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#aigera',
  router,
  template: '<Aigera />',
  components: {
    Aigera
  },
  data: {
    loading: false
  }
});

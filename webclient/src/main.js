// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import Hacksite from './Hacksite';
import router from './router';
import VeeValidate from 'vee-validate';

Vue.config.productionTip = false;
Vue.use(VeeValidate);

/* eslint-disable no-new */
new Vue({
  el: '#hacksite',
  router,
  template: '<Hacksite/>',
  components: { Hacksite },
  data: {loading: false}
});

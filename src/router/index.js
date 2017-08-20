import Vue from 'vue';
import Router from 'vue-router';
import Projects from '@/components/Projects';
import Details from '@/components/Details';
import Create from '@/components/Create';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Projects',
      component: Projects
    },
    {
      path: '/details/:pid',
      name: 'Details',
      component: Details,
      props: (route) => {
        return { pid: parseInt(route.params.pid) };
      }
    },
    {
      path: '/create',
      name: 'Create',
      component: Create
    }
  ]
});

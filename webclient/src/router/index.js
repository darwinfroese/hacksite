import Vue from 'vue';
import Router from 'vue-router';
import Projects from '@/components/views/ProjectList';
import Details from '@/components/views/Details';
import Create from '@/components/views/Create';
import Edit from '@/components/views/Edit';

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
      path: '/edit/:pid',
      name: 'Edit',
      component: Edit,
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

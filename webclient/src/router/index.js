import Vue from 'vue';
import Router from 'vue-router';
import Projects from '@/components/views/ProjectList';
import Details from '@/components/views/Details';
import Create from '@/components/views/Create';
import Edit from '@/components/views/Edit';
import Iteration from '@/components/views/Iteration';
import IterationList from '@/components/views/IterationList';
import CreateAccount from '@/components/views/CreateAccount';

Vue.use(Router);

// TODO: Route better - use the project id inside the
// route better (ie. :pid/details, :pid/iterations, etc.)
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
    },
    {
      path: '/iteration/:pid',
      name: 'CreateIteration',
      component: Iteration,
      props: (route) => {
        return { pid: parseInt(route.params.pid) };
      }
    },
    {
      path: '/iterations/:pid',
      name: 'AllIterations',
      component: IterationList,
      props: (route) => {
        return { pid: parseInt(route.params.pid) };
      }
    },
    {
      path: '/createaccount',
      name: 'CreateAccount',
      component: CreateAccount
    }
  ]
});

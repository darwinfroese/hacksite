import Vue from 'vue';
import Router from 'vue-router';
import Projects from '@/components/views/ProjectList';
import Details from '@/components/views/Details';
import Create from '@/components/views/Create';
import Edit from '@/components/views/Edit';
import Iteration from '@/components/views/Iteration';
import IterationList from '@/components/views/IterationList';
import CreateAccount from '@/components/views/CreateAccount';
import Login from '@/components/views/Login';
import { Authenticate } from '@/database';

Vue.use(Router);

// TODO: Route better - use the project id inside the
// route better (ie. :pid/details, :pid/iterations, etc.)
const router = new Router({
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
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (to.path === '/login') {
    next();
    return;
  }

  Authenticate()
    .then((response) => {
      if (response.status === 401) {
        next('/login');
      }
      if (response.status === 200) {
        next();
      }
    });
});

export default router;

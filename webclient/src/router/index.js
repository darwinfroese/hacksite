import Vue from 'vue';
import Router from 'vue-router';
import Projects from '@/components/views/ProjectList';
import Details from '@/components/views/Details';
import Create from '@/components/views/Create';
import Edit from '@/components/views/Edit';
import Evolution from '@/components/views/Evolution';
import EvolutionList from '@/components/views/EvolutionList';
import CreateAccount from '@/components/views/CreateAccount';
import Login from '@/components/views/Login';
import ReleaseNotes from '@/components/views/ReleaseNotes';
import { Authenticate } from '@/database';

Vue.use(Router);

// TODO: Route better - use the project id inside the
// route better (ie. :pid/details, :pid/evolutions, etc.)
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
        return { pid: route.params.pid };
      }
    },
    {
      path: '/edit/:pid',
      name: 'Edit',
      component: Edit,
      props: (route) => {
        return { pid: route.params.pid };
      }
    },
    {
      path: '/create',
      name: 'Create',
      component: Create
    },
    {
      path: '/evolution/:pid',
      name: 'CreateEvolution',
      component: Evolution,
      props: (route) => {
        return { pid: route.params.pid };
      }
    },
    {
      path: '/evolutions/:pid',
      name: 'AllEvolutions',
      component: EvolutionList,
      props: (route) => {
        return { pid: route.params.pid };
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
    },
    {
      path: '/releasenotes',
      name: 'ReleaseNotes',
      component: ReleaseNotes
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (to.path === '/login' || to.path === '/createaccount') {
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

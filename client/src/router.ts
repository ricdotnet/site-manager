import 'vue-router';
import { createRouter, createWebHistory } from 'vue-router';
import { useUserStore } from '@stores';
import * as pages from './pages';

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth: boolean;
    isAdmin?: boolean;
    isAuthPage?: boolean;
  }
}

const routes = [
  {
    path: '/',
    component: pages.Home,
    meta: {
      requiresAuth: false,
    },
  },
  {
    path: '/login',
    component: pages.Login,
    meta: {
      requiresAuth: false,
      isAuthPage: true,
    },
  },
  {
    path: '/register',
    component: pages.Register,
    meta: {
      requiresAuth: false,
      isAuthPage: true,
    },
  },
  {
    path: '/dashboard',
    component: pages.Dashboard,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        component: pages.Overview,
        name: 'overview',
      },
      {
        path: 'sites',
        component: pages.Sites,
        name: 'sites',
      },
      {
        path: 'sites/:id',
        component: pages.Site,
        name: 'site-details',
      },
      {
        path: 'domains',
        component: pages.Domains,
        name: 'domains',
      },
      {
        path: 'settings',
        component: pages.Settings,
        name: 'settings',
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    component: pages.NotFound,
    name: 'not-found',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to) => {
  const userStore = useUserStore();

  if (!userStore.isAuthed && localStorage.getItem('token')) await userStore.tokenAuth();

  if (to.meta.requiresAuth && !to.meta.isAuthPage && !userStore.isAuthed) return '/login';

  if (to.meta.isAuthPage && userStore.isAuthed) return '/';
});

export { router };

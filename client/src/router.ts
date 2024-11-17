import * as pages from './pages';
import { createRouter, createWebHistory } from 'vue-router';
import { useUserStore } from '@stores';

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
    path: '/verify-code',
    component: pages.VerifyCode,
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
        path: 'domains/:domain',
        component: pages.DomainDetails,
        name: 'domain-details',
        children: [
          {
            path: ':type(a|aaaa|cname|mx|ns|txt|srv)',
            component: pages.DNSRecords,
            name: 'dns-records',
          },
        ],
      },
      {
        path: 'settings',
        component: pages.Settings,
        children: [
          {
            path: '',
            component: pages.NotBuiltYet, // this should default always to /profile
            name: 'settings',
          },
          {
            path: 'profile',
            component: pages.NotBuiltYet,
            name: 'profile',
          },
          {
            path: 'api-keys',
            component: pages.ApiKeys,
            name: 'api-keys',
          },
          {
            path: 'active-sessions',
            component: pages.ActiveSessions,
            name: 'active-sessions',
          },
        ],
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

  if (!userStore.isAuthed && !to.meta.isAuthPage) {
    await userStore.tokenAuth();
  }

  if (to.meta.requiresAuth && !to.meta.isAuthPage && !userStore.isAuthed) {
    return '/login';
  }

  if (to.meta.isAuthPage && userStore.isAuthed) return '/';
});

export { router };

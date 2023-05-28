import 'vue-router';
import { createRouter, createWebHistory } from "vue-router";
import * as pages from "./pages";
import { useUserStore } from "./stores/user.store.ts";

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth: boolean;
    isAdmin?: boolean;
    isAuthPage?: boolean;
  }
}

const routes = [
  {
    path: '/', component: pages.Home, meta: {
      requiresAuth: false,
    }
  },
  {
    path: '/login', component: pages.Login, meta: {
      requiresAuth: false,
      isAuthPage: true,
    }
  },
  {
    path: '/register', component: pages.Register, meta: {
      requiresAuth: false,
      isAuthPage: true,
    }
  }
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

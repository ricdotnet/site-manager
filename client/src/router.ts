import { createRouter, createWebHistory } from "vue-router";
import * as pages from "./pages";

const routes = [
  { path: '/', component: pages.Home },
  { path: '/login', component: pages.Login },
  { path: '/register', component: pages.Register }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export { router };

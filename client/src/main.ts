import { createApp } from 'vue';
import { router } from "./router.ts";
import './style.scss';
import App from './App.vue';

const app = createApp(App);

app.provide('base', import.meta.env.VITE_BASE);
app.provide('api', import.meta.env.VITE_API);

app.use(router)
  .mount('#app');

import './style.scss';
import App from './App.vue';
import { createApp } from 'vue';
import { createHead } from '@unhead/vue/client';
import { createPinia } from 'pinia';
import { router } from './router.ts';

const app = createApp(App);
const pinia = createPinia();
const head = createHead();

app.provide('base', import.meta.env.VITE_BASE);
app.provide('api', import.meta.env.VITE_API);

app.use(router).use(pinia).use(head).mount('#app');

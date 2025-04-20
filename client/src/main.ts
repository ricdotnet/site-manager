import './style.scss';
import App from './App.vue';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { router } from './router.ts';

const app = createApp(App);
const pinia = createPinia();

app.provide('base', import.meta.env.VITE_BASE);
app.provide('api', import.meta.env.VITE_API);

app.use(router).use(pinia).mount('#app');

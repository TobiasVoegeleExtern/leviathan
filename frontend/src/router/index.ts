import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/HomeView.vue';
import Haushalt  from '../views/Haushalt.vue';
const routes = [
  { path: '/', component: Home },
  { path: '/haushalt', component: Haushalt },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

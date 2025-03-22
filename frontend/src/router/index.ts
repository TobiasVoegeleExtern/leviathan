import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import Haushalt  from '../views/Haushalt.vue';
const routes = [
  { path: '/', component: HomeView },
  { path: '/haushalt', component: Haushalt },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

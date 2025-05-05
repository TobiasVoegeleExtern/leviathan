import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import Haushalt  from '../views/Haushalt/Haushalt.vue';
import LoginRegister from '../views/LoginRegister/LoginRegisterForm.vue'
import Profile from '../views/Profile/ProfileView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/haushalt', component: Haushalt },
  { path: '/loginregister', component: LoginRegister },
  { path: '/profile', component: Profile },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;

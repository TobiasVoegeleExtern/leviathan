<template>
  
  <Layout :menu-items="menuItems" app-name="Leviathan">
    <template #header>
      <h1>Welcome </h1>
    
    </template>

    <template #default>
      <div>
        <LoginRegister></LoginRegister>
        <p>Neuer Test</p>
      </div>
    </template>
  </Layout>
</template>

<script setup lang="ts">
import { ref} from 'vue';

import Layout from './components/AppLayout.vue';
import LoginRegister from './views/LoginRegister/LoginRegisterForm.vue';

// Menu items for the layout
const menuItems = ref<{ link: string; label: string }[]>([
{link:'/loginregister',label:'Login/Register'},  
{ link: '/profile', label: 'Profil' },
{ link: '/', label: 'Home' },
  { link: '/haushalt', label: 'Haushaltsbuch' },
]);
</script>

<style scoped>
h1 {
  color: #ffffff;
}
p {
  margin: 0.5rem 0;
}
</style>
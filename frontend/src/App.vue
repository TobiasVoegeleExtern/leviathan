<template>
  <Layout :menu-items="menuItems" app-name="Vue App">
    <template #header>
      <h1>Welcome to Vue with TypeScript</h1>
    
    </template>

    <template #default>
      <div>
        <p>Neuer Test</p>
      </div>
    </template>
  </Layout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

import Layout from '../src/components/AppLayout.vue';


const goMessage = ref<string>('Loading...');




// Menu items for the layout
const menuItems = ref<{ link: string; label: string }[]>([
  { link: '/home', label: 'Home' },
  { link: '/haushalt', label: 'Haushalt' },
  { link: '/contact', label: 'Contact' },
]);
</script>

<style scoped>
h1 {
  color: #42b983;
}
p {
  margin: 0.5rem 0;
}
</style>
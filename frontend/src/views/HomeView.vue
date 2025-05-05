<template>
  <div class="home">
   <h1>Component Test</h1>
    <Button 
      label="Ausgabe hinzufügen"
      :execute="toggleFormVisibility"
    />
    <hr>
    <h2>Data Table</h2>
    <div v-if="!showForm">
    <DataTable :columns="columns" :data="rows" :onDelete="handleDelete"/>
    <h2>Data Table pixel style</h2>
    <PixelTable :columns="columns" :data="rows" />

    
  </div>

  
    <hr>
    <h2>Modals</h2>
 
    <button @click="openBasicDialog" class="open-btn">Open Basic Dialog</button>
    <button @click="openFullscreenDialog" class="open-btn">Open Fullscreen Dialog</button>

    <!-- Basic Dialog -->
    <Dialog ref="basicDialog" title="Delete Item?" confirmText="Delete" @confirm="handleDelete">
      <p>Are you sure you want to delete this item? This action cannot be undone.</p>
    </Dialog>

    <!-- Fullscreen Dialog -->
    <Dialog ref="fullscreenDialog" title="Edit Details" fullscreen>
      <p>This is a full-screen modal for editing details.</p>
    </Dialog>

    <hr>
    <h2>Checkbox</h2>
    <CheckBox></CheckBox>
    <hr>
    <RadioButton></RadioButton>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

import Button from "../components/material/Button.vue"; 
import DataTable from "../components/material/DataTable.vue";
import Dialog from '../components/material/Modal.vue'
import PixelTable from "../components/pixel/PixelTable.vue";
import CheckBox from "../components/material/CheckBox.vue";
import RadioButton from "../components/material/RadioButton.vue";
const showForm = ref(false);


const toggleFormVisibility = async () => {
  showForm.value = !showForm.value;
};


const columns = ref([
  { label: 'Bezeichnung', field: 'description' },
  { label: 'Komplettbetrag', field: 'valuetotal' },
  { label: 'Monatliche Rate', field: 'valuerate' },
  { label: 'Kreditstart', field: 'creditstart' },
  { label: 'Kreditende', field: 'creditend' },
  { label: 'Typ', field: 'type' },
]);

// Table data
const rows = ref([
  { id:1,description: 'MacBook 10 Pro ', valuetotal: 12800, valuerate: 399.95,type: 'Kredit' },
  { id:2,description: 'RTX 5070 Ti', valuetotal: 999, type: 'Sofortkauf' },
  { id:3,description: 'Handy Rechnung', valuetotal: 59.99, type: 'monthlycosts' },

]);

const basicDialog = ref();
const fullscreenDialog = ref();

const openBasicDialog = () => basicDialog.value.openDialog();
const openFullscreenDialog = () => fullscreenDialog.value.openDialog();

const handleDelete = () => alert("Item deleted!");
</script>


<style scoped>
.home .dropdown {
  margin-left: 0; /* Align dropdown to the left */
}
</style>

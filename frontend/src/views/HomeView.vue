<template>
  <div class="home">
   <h1>Component Test</h1>
    <Button 
      label="Ausgabe hinzufügen"
      :execute="toggleFormVisibility"
    />
    <p></p>
    <div v-if="!showForm">
    <DataTable :columns="columns" :data="rows" />
  </div>
    <!-- Conditionally render the form -->
    <Form v-if="showForm" />

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
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Form from "../complexcomponents/AusgabenForm.vue";
import Button from "../components/material/Button.vue"; 
import DataTable from "../components/material/DataTable.vue";
import Dialog from '../components/material/Modal.vue'
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
  { description: 'MacBook 10 Pro ', valuetotal: 12800, valuerate: 399.95,type: 'Kredit' },
  { description: 'RTX 5070 Ti', valuetotal: 999, type: 'Sofortkauf' },
  { description: 'Handy Rechnung', valuetotal: 59.99, type: 'monthlycosts' },

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

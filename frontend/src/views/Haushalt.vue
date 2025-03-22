<template>
  <div class="home">
    <!-- Toggle button to show/hide the form -->
    <Button 
      label="Ausgabe hinzufügen"
      :execute="toggleFormVisibility"
    />
    <p></p>
    <div v-if="!showForm">
    <DataTable :columns="columns" :data="rows" />
  </div>

    <Form v-if="showForm" />

    
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Form from "../complexcomponents/AusgabenForm.vue";
import Button from "../components/material/Button.vue"; 
import DataTable from "../components/material/DataTable.vue";
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

</script>


<style scoped>
.home .dropdown {
  margin-left: 0; /* Align dropdown to the left */
}
</style>

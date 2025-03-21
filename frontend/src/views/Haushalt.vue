<template>
  <div class="home">
    <!-- Toggle button to show/hide the form -->
    <Button 
      label="Ausgabe hinzufügen"
      :execute="toggleFormVisibility"
    />

    <!-- Conditionally render the form -->
    <Form v-if="showForm" />
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Form from "../complexcomponents/Form.vue";
import Button from "../components/Button.vue"; 

// Local state to toggle the form visibility
const showForm = ref(false);

// Function to toggle form visibility
const toggleFormVisibility = () => {
  showForm.value = !showForm.value;
};
</script>

<style scoped>
.home .dropdown {
  margin-left: 0; /* Align dropdown to the left */
}
</style>

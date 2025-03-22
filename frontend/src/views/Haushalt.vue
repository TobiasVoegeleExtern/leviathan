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

const showForm = ref(false);

// Make the function async to match the expected type
const toggleFormVisibility = async () => {
  showForm.value = !showForm.value;
};
</script>


<style scoped>
.home .dropdown {
  margin-left: 0; /* Align dropdown to the left */
}
</style>

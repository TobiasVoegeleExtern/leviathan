<template>
    <form @submit.prevent="handleSubmit" class="form-wrapper">
      <!-- Form content is rendered here via slot -->
      <slot></slot> 
      
      <!-- Optional: Display error message if there's a form error -->
      <p v-if="formError" class="error-message">{{ formError }}</p>
      
      <!-- Form fields -->
      <TextInput 
        label="Bezeichnung" 
        inputType="text" 
        showError="true" 
        errorMessage="This field is required" 
      /> 
  
      <TextInput 
        v-model="inputValue" 
        label="Betrag in Euro" 
        inputType="number" 
        showError="true" 
        errorMessage="This field is required" 
      />
  
      <!-- Faelligkeitsdatum Input (conditionally rendered based on dropdown selection) -->
      <TextInput 
        v-if="isFaelligkeitsdatumVisible"
        v-model="inputValue"
       
        inputType="date" 
        showError="true" 
        errorMessage="This field is required" 
      />
  
      <Dropdown
        :options="dropdownOptions"
        :selected-option="selected"
        @update:selected-option="handleSelection"
      />
      
      <!-- Submit Button -->
      <Button 
        label="Submit" 
        :execute="submitForm"
        :disabled="!isFormValid"
      />
    </form>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue';
  import TextInput from '../components/TextInput.vue'; 
  import Button from '../components/Button.vue'; 
  import Dropdown from "../components/DropdownSelect.vue"; 
  
  interface Option {
    label: string;
    value: string;
  }
  
  // Form error message state
  const formError = ref<string | null>(null);
  
  // Form data
  const inputText = ref('');
  const selected = ref<Option | null>(null);
  
  // Dropdown options
  const dropdownOptions: Option[] = [
    { label: "Sonstiges", value: "allelse" },
    { label: "Kredite", value: "credit" },
    { label: "Laufende Kosten", value: "monthlycosts" },
  ];
  
  // Computed property to check if the form is valid
  const isFormValid = computed(() => {
    return inputText.value.trim() !== '' && selected.value !== null;
  });
  
  // Computed property to check if Faelligkeitsdatum should be shown
  const isFaelligkeitsdatumVisible = computed(() => {
    return selected.value?.value === "credit" || selected.value?.value === "monthlycosts";
  });
  
  // Handle form submission
  const handleSubmit = () => {
    console.log('Form submitted!');
  };
  
  // Reusable method for form submission execution (can be overridden)
  const submitForm = async () => {
    try {
      formError.value = null; // Reset previous errors
      await new Promise(resolve => setTimeout(resolve, 2000)); // Simulate a delay
      alert('Form submitted successfully!');
    } catch (error) {
      console.error('Error during submission', error);
      formError.value = 'There was an error during submission.';
    }
  };
  
  // Handle selection in the dropdown
  const handleSelection = (option: Option) => {
    selected.value = option;
    console.log('Selected option:', option);
  };
  </script>
  
  <style scoped>
  .form-wrapper {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 8px;
    max-width: 500px;
    margin: 0 auto;
  }
  
  .error-message {
    color: #e74c3c;
    font-size: 0.875rem;
    margin-top: 0.25rem;
    text-align: center;
  }
  
  /* Make the dropdown and its options the same width */
  .dropdown-wrapper {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    width: 100%;
  }
  
  .dropdown-wrapper select {
    width: 100%;
    padding: 0.5rem;
    border-radius: 4px;
    border: 1px solid #ccc;
  }
  
  /* Optional: Ensure the options of the dropdown are full width */
  .dropdown-wrapper option {
    width: 100%;
    padding: 0.5rem;
  }
  </style>
  
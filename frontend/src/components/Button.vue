<template>
    <button
      :class="['material-btn', { 'loading': loading }]"
      :disabled="loading"
      @click="handleClick"
      :aria-label="loading ? 'Loading... Please wait' : label" 
      :aria-busy="loading.toString()" 
    >
      <!-- Spinner for loading state -->
      <span v-if="loading" class="spinner" aria-hidden="true"></span>
      
      <!-- Button label when not loading -->
      <span v-else>{{ label }}</span>
    </button>
  </template>
  
  <script setup lang="ts">
  import { ref } from "vue";
  
  // Define props for label and execute function
  const props = defineProps<{
    label: string;  // The label text to display on the button
    execute: () => void;  // The function to be executed when button is clicked
  }>();
  
  // Local state to handle the loading state of the button
  const loading = ref(false);
  
  // Method to handle button click and execute the provided execute function
  const handleClick = async () => {
    loading.value = true; // Set loading state to true when button is clicked
    try {
      await props.execute();  // Execute the function passed through the `execute` prop
    } catch (error) {
      console.error("Error occurred:", error);
    } finally {
      loading.value = false;  // Set loading state to false after execution
    }
  };
  </script>
  
  <style scoped>
  /* Material Design Button */
  .material-btn {
    background-color: #6200ea;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    width: auto; /* Auto width to fit the label */
    height: 40px; /* Fixed height for consistency */
    min-width: 120px; /* Optional: minimum width to keep it consistent */
    max-width: 300px; /* Optional: max width to prevent it from becoming too wide */
    transition: background-color 0.3s, transform 0.3s ease;
  }
  
  /* Hover and active states */
  .material-btn:hover {
    background-color: #3700b3; /* Darker blue on hover */
  }
  
  .material-btn:active {
    transform: scale(0.98); /* Slight shrink effect on click */
  }
  
  .material-btn:disabled {
    background-color: #9e9e9e;
    cursor: not-allowed;
  }
  
  /* Spinner while loading */
  .spinner {
    border: 4px solid #f3f3f3; /* Light grey */
    border-top: 4px solid #6200ea; /* Purple color for the spinner */
    border-radius: 50%;
    width: 20px;
    height: 20px;
    animation: spin 1s linear infinite;
    position: absolute;
  }
  
  /* Spinner animation */
  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  </style>
  
<template>
  <form @submit.prevent="handleSubmit" class="form-wrapper">
    <!-- Dropdown moved to the top of the form -->
    <Dropdown
      :options="dropdownOptions"
      v-model="selected"
      @update:selected-option="handleSelection"
      class="dropdown-wrapper" 
    />

    <!-- Optional: Display error message if there's a form error -->
    <p v-if="formError" class="error-message">{{ formError }}</p>

    <!-- Bezeichnung Input with error handling -->
    <TextInput 
      label="Bezeichnung" 
      inputType="text" 
      v-model="inputText" 
      :showError="!isBezeichnungValid"
      errorMessage="This field is required" 
      class="input-with-margin"
    />

    <!-- Betrag Input with number validation -->
    <TextInput 
      label="Betrag" 
      inputType="text" 
      v-model="betragInput" 
      :showError="!isBetragValid"
      errorMessage="Please enter a valid number" 
    />

    <!-- Faelligkeitsdatum Input (conditionally rendered based on dropdown selection) -->
    <TextInput 
      v-if="isFaelligkeitsdatumVisible"
      label="Monatliches Faelligkeitsdatum"
      v-model="inputText"
      inputType="date" 
      :showError="!isFaelligkeitsdatumValid"
      errorMessage="This field is required" 
    />

    <!-- Additional fields for "Kredite" -->
    <TextInput 
      v-if="isKrediteSelected"
      label="Kreditbeginn"
      inputType="date" 
      v-model="abbezahltInput" 
      :showError="!isAbbezahltValid"
      errorMessage="Please enter a valid date" 
    />

    <TextInput 
      v-if="isKrediteSelected"
      label="Kreditende"
      inputType="date" 
      v-model="zahltagInput" 
      :showError="!isZahltagValid"
      errorMessage="Please enter a valid day of the month" 
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
const betragInput = ref('');
const abbezahltInput = ref('');
const zahltagInput = ref('');
const selected = ref<Option | null>(null);

// Dropdown options
const dropdownOptions: Option[] = [
  { label: "Sonstiges", value: "allelse" },
  { label: "Kredite", value: "credit" },
  { label: "Laufende Kosten", value: "monthlycosts" },
];

// Validation logic for Bezeichnung (required)
const isBezeichnungValid = computed(() => {
  return inputText.value.trim() !== '';
});

// Validation logic for Betrag (number and required)
const isBetragValid = computed(() => {
  return !isNaN(parseFloat(betragInput.value)) && betragInput.value.trim() !== ''; 
});

// Validation logic for Faelligkeitsdatum (required when dropdown option is selected)
const isFaelligkeitsdatumValid = computed(() => {
  return selected.value?.value === 'credit' || selected.value?.value === 'monthlycosts'
    ? inputText.value.trim() !== '' 
    : true; // Not required if dropdown selection doesn't match
});

// Validation logic for Abbezahlt bis (required if Kredite is selected)
const isAbbezahltValid = computed(() => {
  return isKrediteSelected.value ? abbezahltInput.value.trim() !== '' : true;
});

// Validation logic for Zahltag (required if Kredite is selected, and must be between 1 and 31)
const isZahltagValid = computed(() => {
  return isKrediteSelected.value && /^[1-9]$|^[12][0-9]$|^3[01]$/.test(zahltagInput.value);
});

// Computed property to check if the form is valid
const isFormValid = computed(() => {
  return isBezeichnungValid.value && isBetragValid.value && isFaelligkeitsdatumValid.value && isAbbezahltValid.value && isZahltagValid.value && selected.value !== null;
});

// Computed property to check if Faelligkeitsdatum should be shown
const isFaelligkeitsdatumVisible = computed(() => {
  return selected.value?.value === "credit" || selected.value?.value === "monthlycosts";
});

// Computed property to check if Kredite is selected
const isKrediteSelected = computed(() => {
  return selected.value?.value === "credit";
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
  position: relative;
  z-index: 1000; /* Ensure the dropdown has the highest z-index */
}

.dropdown-wrapper select {
  width: 100%;
  padding: 0.5rem;
  border-radius: 4px;
  border: 1px solid #ccc;
}

/* Add more margin to the first input field */
.input-with-margin {
  margin-top: 1rem; /* Add more space above the first input field */
}

/* Optional: Ensure the options of the dropdown are full width */
.dropdown-wrapper option {
  width: 100%;
  padding: 0.5rem;
}
</style>

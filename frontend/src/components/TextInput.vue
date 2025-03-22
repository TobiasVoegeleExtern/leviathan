<template>
  <div class="input-container">
    <!-- Floating Label -->
    <label
      v-if="label"
      class="input-label"
      :class="{ 'float-label': modelValue && modelValue !== '' }"
    >
      {{ label }}
    </label>

    <!-- Input Field -->
    <input
      :type="inputType"
      :placeholder="!modelValue ? placeholder : ''"
      v-model="modelValue"
      :disabled="disabled"
      class="input-field"
      @input="handleInput"
    />

    <!-- Error Message -->
    <div v-if="showError" class="error-message">{{ errorMessage }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits, watch } from 'vue';

const props = defineProps({
  modelValue: { type: String, default: '' },
  label: { type: String, default: '' },
  inputType: { type: String, default: 'text' }, // Allow text, number, password, or date
  placeholder: { type: String, default: '' },
  disabled: { type: Boolean, default: false },
  showError: { type: Boolean, default: false },
  errorMessage: { type: String, default: '' },
});

const emit = defineEmits(['update:modelValue']);

const modelValue = ref(props.modelValue);

// Sync v-model
watch(modelValue, (newVal) => {
  emit('update:modelValue', newVal);
});

// Handle input change for text, number, password, or date types (no formatting)
const handleInput = (event: Event) => {
  const input = event.target as HTMLInputElement;
  // Simply update the modelValue based on the input value
  modelValue.value = input.value;
};
</script>

<style scoped>
.input-container {
  display: flex;
  flex-direction: column;
  position: relative;
  margin-bottom: 1rem;
  width: 100%;
}

.input-label {
  font-size: 1rem; /* Label font size */
  color: #555;
  position: absolute;
  left: 12px;
  top: 16px; /* Adjusted to ensure proper label positioning */
  transition: 0.3s ease all;
  pointer-events: none; /* Prevent label from interfering with input */
  z-index: 2; /* Higher z-index to ensure label stays above input field */
}

/* Make the label float above the input field */
.input-label.float-label {
  top: -20px; /* Increased distance between label and input box */
  left: 8px; /* Shift the label more to the left */
  font-size: 1rem; /* Smaller font size when the label floats */
  color: #6200ea; /* Optional: Change color to purple when label floats */
}

.input-field {
  padding: 16px 12px;
  font-size: 1rem; /* Match placeholder and label font size */
  border: 1px solid #ccc;
  border-radius: 4px;
  transition: all 0.3s ease;
  outline: none;
  background-color: white;
  position: relative;
  z-index: 1; /* Lower z-index to ensure input stays below the label */
}

/* Ensure the placeholder matches the size of the label */
.input-field::placeholder {
  font-size: 1rem; /* Same font size as label */
  color: #aaa; /* Lighter color for placeholder */
}

.input-field:focus {
  border-color: #6200ea; /* Purple border on focus */
  box-shadow: 0 0 5px rgba(98, 0, 234, 0.3);
}

.input-field:disabled {
  background-color: #f5f5f5;
  border-color: #e0e0e0;
  cursor: not-allowed;
}

.error-message {
  color: #e74c3c;
  font-size: 0.875rem;
  margin-top: 0.25rem;
}
</style>

<template>
    <label class="checkbox-container">
      <input 
        type="checkbox" 
        v-model="checked" 
        class="checkbox-input"
        @change="emitChange"
      />
      <span class="checkbox-custom">
        <svg class="checkbox-icon" viewBox="0 0 24 24">
          <path d="M20.65 5.35a1 1 0 00-1.41 0L9 15.59l-4.24-4.24a1 1 0 00-1.41 1.41l5 5a1 1 0 001.41 0l11-11a1 1 0 000-1.41z"></path>
        </svg>
      </span>
      <span class="checkbox-label">{{ label }}</span>
    </label>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps, defineEmits } from "vue";
  
  const props = defineProps({
    modelValue: Boolean,
    label: { type: String, default: "" },
  });
  
  const emit = defineEmits(["update:modelValue"]);
  const checked = ref(props.modelValue);
  
  const emitChange = () => {
    emit("update:modelValue", checked.value);
  };
  </script>
  
  <style scoped>
  /* Checkbox Container */
  .checkbox-container {
    display: flex;
    align-items: center;
    cursor: pointer;
    font-family: "Roboto", sans-serif;
    user-select: none;
    gap: 8px;
  }
  
  /* Hide default checkbox */
  .checkbox-input {
    display: none;
  }
  
  /* Custom Checkbox */
  .checkbox-custom {
    width: 20px;
    height: 20px;
    border: 2px solid #6200ea; /* Material Purple */
    border-radius: 4px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
  }
  
  /* Checkmark Icon */
  .checkbox-icon {
    width: 16px;
    height: 16px;
    fill: white;
    transform: scale(0);
    transition: transform 0.2s ease-in-out;
  }
  
  /* Checked State */
  .checkbox-input:checked + .checkbox-custom {
    background-color: #6200ea;
    border-color: #6200ea;
  }
  
  .checkbox-input:checked + .checkbox-custom .checkbox-icon {
    transform: scale(1);
  }
  
  /* Ripple Effect */
  .checkbox-custom::before {
    content: "";
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgba(98, 0, 234, 0.3);
    border-radius: 50%;
    transform: scale(0);
    transition: transform 0.3s ease-out;
  }
  
  .checkbox-input:checked + .checkbox-custom::before {
    transform: scale(2);
  }
  
  /* Checkbox Label */
  .checkbox-label {
    font-size: 16px;
    color: #333;
    transition: color 0.3s ease;
  }
  
  .checkbox-input:checked ~ .checkbox-label {
    color: #6200ea;
  }
  </style>
  
<template>
    <label class="radio-container">
      <input 
        type="radio" 
        :value="value" 
        v-model="selected" 
        class="radio-input"
        @change="emitChange"
      />
      <span class="radio-custom"></span>
      <span class="radio-label">{{ label }}</span>
    </label>
  </template>
  
  <script setup lang="ts">
  import { defineProps, defineEmits, ref, watch } from "vue";
  
  const props = defineProps({
    modelValue: [String, Number],
    value: [String, Number],
    label: { type: String, default: "" },
  });
  
  const emit = defineEmits(["update:modelValue"]);
  const selected = ref(props.modelValue);
  
  const emitChange = () => {
    emit("update:modelValue", selected.value);
  };
  
  watch(() => props.modelValue, (newValue) => {
    selected.value = newValue;
  });
  </script>
  
  <style scoped>
  /* Radio Button Container */
  .radio-container {
    display: flex;
    align-items: center;
    cursor: pointer;
    font-family: "Roboto", sans-serif;
    user-select: none;
    gap: 8px;
  }
  
  /* Hide Default Radio */
  .radio-input {
    display: none;
  }
  
  /* Custom Radio Button */
  .radio-custom {
    width: 20px;
    height: 20px;
    border: 2px solid #6200ea; /* Material Purple */
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    transition: all 0.3s ease;
  }
  
  /* Inner Circle */
  .radio-custom::before {
    content: "";
    width: 12px;
    height: 12px;
    background: #6200ea;
    border-radius: 50%;
    transform: scale(0);
    transition: transform 0.2s ease-in-out;
  }
  
  /* Checked State */
  .radio-input:checked + .radio-custom::before {
    transform: scale(1);
  }
  
  /* Ripple Effect */
  .radio-custom::after {
    content: "";
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgba(98, 0, 234, 0.3);
    border-radius: 50%;
    transform: scale(0);
    transition: transform 0.3s ease-out;
  }
  
  .radio-input:checked + .radio-custom::after {
    transform: scale(2);
  }
  
  /* Radio Label */
  .radio-label {
    font-size: 16px;
    color: #333;
    transition: color 0.3s ease;
  }
  
  .radio-input:checked ~ .radio-label {
    color: #6200ea;
  }
  </style>
  
<template>
    <router-link
      :to="to"
      :class="['custom-link', customClass]"
      @mouseover="handleMouseOver"
      @mouseleave="handleMouseLeave"
    >
      {{ label }}
    </router-link>
  </template>
  
  <script setup lang="ts">
  import { defineProps } from 'vue';
  
  // Define props for the CustomLink component
  defineProps({
    to: {
      type: String,
      required: true
    },
    label: {
      type: String,
      required: true
    },
    customClass: {
      type: String,
      default: ''
    }
  });
  
  // Methods for mouse interactions
  const handleMouseOver = () => {
    console.log('Mouse over the link');
  };
  
  const handleMouseLeave = () => {
    console.log('Mouse leave the link');
  };
  </script>
  
  <style scoped>
  .custom-link {
    color: #3498db;
    text-decoration: none;
    font-size: 1rem;
    padding: 5px 10px;
    display: inline-block;
    border-radius: 4px;
    transition: background-color 0.3s, color 0.3s;
  }
  
  .custom-link:hover {
    background-color: #2980b9;
    color: white;
  }
  
  .custom-link:active {
    background-color: #1f6f94;
    color: white;
  }
  
  .custom-link:focus {
    outline: none;
  }
  </style>
  
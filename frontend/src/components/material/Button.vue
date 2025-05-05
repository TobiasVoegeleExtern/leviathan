<template>
  <button
    :class="buttonClass"
    :disabled="loading"
    @click="handleClick"
    :aria-label="loading ? 'Loading... Please wait' : label"
    :aria-busy="loading ? 'true' : 'false'"
  >
    <!-- Spinner for loading state -->
    <span v-if="loading" class="spinner" aria-hidden="true"></span>

    <!-- Button label when not loading -->
    <span v-else>{{ label }}</span>
  </button>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";

// Define props for label, function execution, color, and size
const props = defineProps<{
  label: string;
  execute: (() => void) | (() => Promise<void>);
  color?: "primary" | "secondary" | "success" | "warning" | "danger" | "info";
  size?: "small" | "medium" | "large" | "xlarge" | "xsmall" | "xxsmall" | "xxxsmall";
}>();

// Local state for loading indicator
const loading = ref(false);

// Computed property to generate dynamic classes
const buttonClass = computed(() => [
  "material-btn",
  props.color || "primary", // Default to primary if no color is set
  props.size || "medium", // Default to medium if no size is set
  { loading: loading.value }, // Adds 'loading' class when button is loading
]);

// Handle button click and execution
const handleClick = async () => {
  loading.value = true;
  try {
    await props.execute();
  } catch (error) {
    console.error("Error occurred:", error);
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* Base Button Style */
.material-btn {
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  transition: background-color 0.3s, transform 0.3s ease;
  font-weight: 600;
  text-transform: uppercase;
  min-width: 100px;
}

/* Colors */
.primary {
  background-color: #6200ea;
}
.primary:hover {
  background-color: #4500b5;
}

.secondary {
  background-color: #6c757d;
}
.secondary:hover {
  background-color: #565e64;
}

.success {
  background-color: #28a745;
}
.success:hover {
  background-color: #218838;
}

.warning {
  background-color: #ffc107;
  color: black;
}
.warning:hover {
  background-color: #e0a800;
}

.danger {
  background-color: #dc3545;
}
.danger:hover {
  background-color: #c82333;
}

.info {
  background-color: #17a2b8;
}
.info:hover {
  background-color: #138496;
}

/* Sizes */
.xsmall {
  padding: 4px 8px;
  font-size: 0.6rem;
  height: 22px;
}

.xxsmall {
  padding: 3px 6px;
  font-size: 0.5rem;
  height: 18px;
}

.xxxsmall {
  padding: 2px 4px;
  font-size: 0.4rem;
  height: 14px;
}

.small {
  padding: 6px 12px;
  font-size: 0.8rem;
  height: 30px;
}

.medium {
  padding: 10px 20px;
  font-size: 1rem;
  height: 40px;
}

.large {
  padding: 14px 28px;
  font-size: 1.2rem;
  height: 50px;
}

.xlarge {
  padding: 18px 32px;
  font-size: 1.5rem;
  height: 60px;
}

/* Disabled State */
.material-btn:disabled {
  background-color: #9e9e9e;
  cursor: not-allowed;
}

/* Spinner */
.spinner {
  border: 4px solid #f3f3f3;
  border-top: 4px solid currentColor;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  animation: spin 1s linear infinite;
  position: absolute;
}

/* Spinner Animation */
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>

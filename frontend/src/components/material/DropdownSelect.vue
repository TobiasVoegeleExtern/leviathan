<template>
  <div class="dropdown" ref="dropdownRef" v-click-outside="closeDropdown">
    <button class="dropdown-btn" @click="toggleDropdown">
      {{ selectedLabel || 'Select an option' }}
      <span class="arrow"></span>
    </button>
    <div v-if="isOpen" class="dropdown-menu">
      <ul>
        <li
          v-for="(option, index) in options"
          :key="index"
          @click="selectOption(option)"
          :class="{ 'selected': selectedLabel === option.label }"
        >
          {{ option.label }}
        </li>
      </ul>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, defineProps, defineEmits, watch, onMounted, onBeforeUnmount } from 'vue';

interface Option {
  label: string;
  value: string;
}

// Define Props and Emits
const props = defineProps<{
  options: Option[];
  selectedOption?: Option | null;
}>();

const emit = defineEmits<{
  (e: "update:selectedOption", option: Option): void;
}>();

// Define reactive properties
const isOpen = ref(false);
const selectedLabel = ref<string | null>(props.selectedOption?.label || null);
const dropdownRef = ref<HTMLElement | null>(null);

// Watch the selected option
watch(
  () => props.selectedOption,
  (newOption) => {
    selectedLabel.value = newOption?.label || null;
  }
);

// Toggle dropdown open/close
const toggleDropdown = () => {
  isOpen.value = !isOpen.value;
};

// Close dropdown
const closeDropdown = () => {
  isOpen.value = false;
};

// Select an option and update the selectedLabel
const selectOption = (option: Option) => {
  selectedLabel.value = option.label; 
  emit("update:selectedOption", option); 
  closeDropdown(); 
};

// Create the clickOutside directive
const clickOutside = (el: HTMLElement, binding: any) => {
  const onClickOutside = (event: MouseEvent) => {
    if (!el.contains(event.target as Node)) {
      binding.value(); // Call the passed method (closeDropdown)
    }
  };

  onMounted(() => {
    document.addEventListener('click', onClickOutside);
  });

  onBeforeUnmount(() => {
    document.removeEventListener('click', onClickOutside);
  });
};

// Register custom directive globally
defineExpose({
  directives: {
    clickOutside
  }
});
</script>


<style scoped>
.dropdown {
  position: relative;
  display: inline-block;
}

.dropdown-btn {
  padding: 10px 20px;
  background-color: #6200ea;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 4px;
  font-size: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 200px;
  transition: all 0.3s ease;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
}

.dropdown-btn:hover {
  background-color: #3700b3;
}

.arrow {
  margin-left: 10px;
  border-top: 5px solid white;
  border-left: 5px solid transparent;
  border-right: 5px solid transparent;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  background-color: white;
  border-radius: 4px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
  z-index: 1;
  overflow: hidden;
}

.dropdown-menu ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.dropdown-menu li {
  padding: 10px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.dropdown-menu li:hover {
  background-color: #f1f1f1;
}

.dropdown-menu li.selected {
  background-color: #6200ea;
  color: white;
}
</style>
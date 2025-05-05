<template>
  <div class="card" :class="{ 'full-width': fullWidth }">
    <img v-if="image" :src="image" alt="Card image" class="card-image">
    
    <div class="card-content">
      <h3 class="card-title">{{ title }}</h3>

      <!-- Horizontal line with the expand/collapse button on top (right corner) -->
      <div class="card-divider-wrapper">
        <button 
          @click="toggleDescription"
          class="card-toggle-button"
        >
          {{ isDescriptionVisible ? 'Collapse' : 'Expand' }}
        </button>
        <hr class="card-divider" />
      </div>

      <!-- Toggling visibility of description -->
      <p v-show="isDescriptionVisible" class="card-description">{{ description }}</p>
      <slot name="description" v-if="isDescriptionVisible"></slot>
    </div>
    <div class="card-footer">
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
  title: String,
  description: String,  // Optional if slot is used
  image: String,
  buttonText: {
    type: String,
    default: "Button"  
  },
  onClick: Function,
  fullWidth: {
    type: Boolean,
    default: false // Standard size by default
  }
});

// Track whether the description is visible or not
const isDescriptionVisible = ref(true);

// Toggle method for description visibility
const toggleDescription = () => {
  isDescriptionVisible.value = !isDescriptionVisible.value;
};
</script>

<style scoped>
.card {
  width: 300px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: box-shadow 0.3s ease-in-out;
}

.card.full-width {
  width: 100%;
}

.card:hover {
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.card-image {
  width: 100%;
  height: 180px;
  object-fit: cover;
}

.card-content {
  padding: 16px;
}

.card-title {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: 8px;
}

.card-divider-wrapper {
  position: relative;
  margin: 8px 0;
}

.card-divider {
  border: none;
  height: 1px;
  background-color: #ddd;
  margin: 0;
}

.card-toggle-button {
  position: absolute;
  top: -18px; /* Moves the button above the hr line */
  right: 10px;
  background: #6200ea;
  color: white;
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.8rem;
  transition: background 0.2s ease-in-out, transform 0.2s ease;
}

.card-toggle-button:hover {
  background: #4500b5;
  transform: scale(1.1);
}

.card-description {
  font-size: 0.9rem;
  color: #555;
}

.card-footer {
  padding: 16px;
  border-top: 1px solid #eee;
  text-align: right;
}
</style>

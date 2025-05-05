<template>
    <Teleport to="body">
      <div v-if="isOpen" class="backdrop" @click.self="closeDialog">
        <transition name="fade">
          <div :class="['dialog', { fullscreen }]">
            <!-- Header -->
            <div class="dialog-header">
              <h2>{{ title }}</h2>
              <button class="close-btn" @click="closeDialog">&times;</button>
            </div>
  
            <!-- Content Slot -->
            <div class="dialog-content">
              <slot></slot>
            </div>
  
            <!-- Actions (Only for Basic Dialog) -->
            <div v-if="!fullscreen" class="dialog-actions">
              <button class="btn cancel-btn" @click="closeDialog">Abbrechen</button>
              <button class="btn confirm-btn" @click="confirmAction">
                {{ confirmText }}
              </button>
            </div>
          </div>
        </transition>
      </div>
    </Teleport>
  </template>
  
  <script setup lang="ts">
  import { defineProps, defineEmits, ref } from "vue";
  
  const props = defineProps({
    title: { type: String, required: true },
    confirmText: { type: String, default: "Confirm" },
    fullscreen: { type: Boolean, default: false }
  });
  
  const emits = defineEmits(["confirm", "close"]);
  const isOpen = ref(false);
  
  const openDialog = () => (isOpen.value = true);
  const closeDialog = () => {
    isOpen.value = false;
    emits("close");
  };
  const confirmAction = () => {
    emits("confirm");
    closeDialog();
  };
  
  // Expose method for parent component
  defineExpose({ openDialog });
  </script>
  
  <style scoped>
  /* Backdrop */
  .backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }
  
  /* Dialog */
  .dialog {
    background: white;
    width: 90%;
    max-width: 500px;
    padding: 1.5rem;
    border-radius: 8px;
    box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.3);
    animation: scaleIn 0.2s ease-out;
  }
  
  /* Fullscreen Dialog */
  .dialog.fullscreen {
    width: 100vw;
    height: 100vh;
    max-width: none;
    border-radius: 0;
    display: flex;
    flex-direction: column;
  }
  
  /* Header */
  .dialog-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: bold;
    color: #333;
  }
  
  /* Close Button */
  .close-btn {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
  }
  
  /* Content */
  .dialog-content {
    margin: 1rem 0;
  }
  
  /* Actions */
  .dialog-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
  }
  
  /* Buttons */
  .btn {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    transition: background 0.2s;
  }
  
  .cancel-btn {
    background: #ccc;
  }
  
  .cancel-btn:hover {
    background: #bbb;
  }
  
  .confirm-btn {
    background: #6200ea;
    color: white;
  }
  
  .confirm-btn:hover {
    background: #4b00c1;
  }
  
  /* Animations */
  @keyframes scaleIn {
    from {
      transform: scale(0.9);
      opacity: 0;
    }
    to {
      transform: scale(1);
      opacity: 1;
    }
  }
  
  /* Transition */
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.2s ease;
  }
  
  .fade-enter-from, .fade-leave-to {
    opacity: 0;
  }
  </style>
  
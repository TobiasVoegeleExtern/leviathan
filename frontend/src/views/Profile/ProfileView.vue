<template>
    <div class="profile-view">
      <h1>User Profile</h1>
  
      <div v-if="user" class="profile-details">
        <p><strong>Name:</strong> {{ user.name }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
        <p><strong>Income:</strong> {{ user.income || 'Not available' }}</p>
        <p><strong>Account Balance:</strong> ${{ user.accountbalance ? user.accountbalance.toFixed(2) : '0.00' }}</p>
  
        <!-- Edit form -->
        <div v-if="isEditing">
          <input v-model="user.name" placeholder="Enter new name" />
          <input v-model="user.email" placeholder="Enter new email" />
          <input v-model="user.income" placeholder="Enter new income" type="number" />
          <input v-model="user.accountbalance" placeholder="Enter new account balance" type="number" step="0.01" />
  
          <button @click="updateUser">Save Changes</button>
        </div>
  
        <!-- Show button to edit or save -->
        <button @click="toggleEditMode">{{ isEditing ? 'Cancel' : 'Edit Profile' }}</button>
        <button @click="deleteUser" class="delete-button">Delete Account</button>
      </div>
  
      <div v-else>
        <p>Loading user data...</p>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref, onMounted } from 'vue';
  import { userService } from '../../services/UserService';
  import type { User } from '../../models/user';
  import { useUserStore } from '../../stores/userStore';
  
  export default defineComponent({
    name: 'ProfileView',
    setup() {
      const userStore = useUserStore();  // Get the store instance
  
      const user = ref<User | null>(userStore.user);  // Get the user from the store
      const isEditing = ref(false);  // Flag to toggle between view and edit mode
  
      // When the component mounts, load the user from storage (if any)
      onMounted(async () => {
        try {
          const userId = userStore.user?.id;  // Get the user ID from the store
          if (userId) {
            user.value = await userService.getUserById(userId);  // Fetch user details from the API
          } else {
            console.error("User is not authenticated");
          }
        } catch (error) {
          console.error("Error fetching user:", error);
        }
      });
  
      // Toggle edit mode
      const toggleEditMode = () => {
        isEditing.value = !isEditing.value;
      };
  
      // Update user profile
      const updateUser = async () => {
        if (user.value) {
          try {
            await userService.updateUser(user.value.id!, user.value);
            alert("Profile updated successfully!");
            isEditing.value = false;  // Exit edit mode after saving
          } catch (error) {
            console.error("Error updating user:", error);
          }
        }
      };
  
      // Delete user profile
      const deleteUser = async () => {
        if (user.value) {
          try {
            await userService.deleteUser(user.value.id!);
            alert("Account deleted successfully!");
            // Redirect user or clear data after deletion
            user.value = null;
          } catch (error) {
            console.error("Error deleting user:", error);
          }
        }
      };
  
      return {
        user,
        isEditing,
        toggleEditMode,
        updateUser,
        deleteUser
      };
    }
  });
  </script>
  
  <style scoped>
  .profile-view {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
  }
  
  .profile-details {
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 8px;
  }
  
  button {
    margin-top: 10px;
    padding: 10px;
    cursor: pointer;
  }
  
  button.delete-button {
    background-color: red;
    color: white;
  }
  
  button:hover {
    opacity: 0.8;
  }
  </style>
  
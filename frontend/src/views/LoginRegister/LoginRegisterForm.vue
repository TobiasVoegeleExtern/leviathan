<template>
  <div class="form-container">
    <form @submit.prevent="handleSubmit" class="form-wrapper">
      <!-- Registration: Username Field -->
      <div v-if="isRegister" class="input-wrapper">
        <TextInput
          label="Username"
          inputType="text"
          v-model="formData.name"
          :showError="isRegister && !isUsernameValid"
          errorMessage="This field is required"
          class="input-with-margin"
        />
      </div>

      <!-- Email or Username Field -->
      <div class="input-wrapper">
        <TextInput
          :label="isRegister ? 'Email' : 'Email or Username'"
          inputType="text"
          v-model="formData.identifier"
          :showError="!isIdentifierValid"
          errorMessage="Please enter a valid email or username"
          class="input-with-margin"
        />
      </div>

      <!-- Password Field -->
      <div class="input-wrapper">
        <TextInput
          label="Password"
          inputType="password"
          v-model="formData.password"
          :showError="!isPasswordValid"
          errorMessage="Password must be at least 6 characters long"
          class="input-with-margin"
        />
      </div>

      <!-- Error Message -->
      <p v-if="formError" class="error-message">{{ formError }}</p>

      <!-- Submit Button -->
      <Button label="Submit" :execute="handleSubmit" :disabled="!isFormValid" />

      <!-- Toggle Between Login and Register -->
      <p class="toggle-form">
        <span @click="toggleForm">
          {{ isRegister ? 'Already have an account? Login' : 'No account yet? Register' }}
        </span>
      </p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useUserStore } from '../../stores/userStore';
import { userService } from '../../services/UserService.ts';
import TextInput from '../../components/material/TextInput.vue';
import Button from '../../components/material/Button.vue';
import { useRouter } from 'vue-router'; // Import useRouter

const userStore = useUserStore();
const router = useRouter(); // Initialize the router
const formData = ref({
  name: '',
  identifier: '', // Can be email or username for login
  password: '',
});

const formError = ref<string | null>(null);
const isRegister = ref(false);

// Validation
const isUsernameValid = computed(() => (isRegister.value ? formData.value.name.trim() !== '' : true));
const isIdentifierValid = computed(() => formData.value.identifier.trim() !== '');
const isPasswordValid = computed(() => formData.value.password.length >= 6);
const isFormValid = computed(() => isIdentifierValid.value && isPasswordValid.value && isUsernameValid.value);


const handleSubmit = async () => {
  if (!isFormValid.value) return;

  try {
    if (isRegister.value) {
      // Register user
      await userService.register({
        name: formData.value.name,
        email: formData.value.identifier,
        password: formData.value.password,
      });
      alert('Registration successful!');
    } else {
      // Login user using the login method from the store
      await userStore.login(formData.value.identifier, formData.value.password);
      // Redirect to home page on successful login
      router.push('/'); // Redirect to home page after successful login
    }

    // Reset form data and error messages
    formData.value = { name: '', identifier: '', password: '' };
    formError.value = null;
  } catch (error: any) {
    formError.value = error.message || 'An error occurred';
  }
};

// Toggle Login/Register
const toggleForm = () => {
  isRegister.value = !isRegister.value;
  formData.value = { name: '', identifier: '', password: '' };
  formError.value = null; // Clear errors when toggling
};

// Load user from local storage when component is mounted
onMounted(() => {
  userStore.loadUserFromStorage();
});
</script>

<style scoped>
.form-container {
  display: flex;
  justify-content: flex-start;
  align-items: flex-start;
  height: 100vh;
  padding: 20px;
}

.form-wrapper {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  max-width: 500px;
  width: 100%;
  margin-top: 0;
}

.error-message {
  color: #e74c3c;
  font-size: 0.875rem;
  margin-top: 0.25rem;
  text-align: center;
}

.input-wrapper {
  display: flex;
  flex-direction: column;
}

.input-with-margin {
  margin-top: 1rem;
}

.toggle-form {
  margin-top: 1rem;
  text-align: center;
  cursor: pointer;
  color: #007bff;
}
</style>
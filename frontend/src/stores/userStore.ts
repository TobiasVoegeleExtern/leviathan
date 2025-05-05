import { defineStore } from 'pinia';
import { userService } from '../services/UserService.ts';
import type { User } from '../models/user';

// Define the state interface explicitly
interface UserStoreState {
  user: User | null;
  isAuthenticated: boolean;
}

export const useUserStore = defineStore('user', {
  // Explicitly typing the state
  state: (): UserStoreState => ({
    user: null,
    isAuthenticated: false,
  }),

  actions: {
    async login(identifier: string, password: string) {
      try {
        // Call the authentication service to check credentials
        const response = await userService.authenticate(identifier, password);

        // Check if the user is found
        if (response && response.id) {
          // Store user info in the state
          this.user = { id: response.id, name: response.name, email: response.email, password: response.password };
          this.isAuthenticated = true;

          // Store user info in localStorage
          localStorage.setItem('user', JSON.stringify(this.user));
        } else {
          throw new Error('Invalid login credentials');
        }
      } catch (error: any) {
        throw new Error(error.message || 'Login failed');
      }
    },

    logout() {
      // Clear user state and localStorage on logout
      this.user = null;
      this.isAuthenticated = false;
      localStorage.removeItem('user');
    },

    loadUserFromStorage() {
      // Check if the user is stored in localStorage
      const storedUser = localStorage.getItem('user');
      if (storedUser) {
        this.user = JSON.parse(storedUser);
        this.isAuthenticated = true;
      }
    },
  },
});

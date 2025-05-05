// src/services/UserService.ts

import axios from 'axios';
import type { User } from '../models/user'; 
const API_URL = 'http://localhost:8000/users/';



class UserService {
  // Register a new user
  async register(user: User): Promise<any> {
    try {
      const response = await axios.post(API_URL, user);
      return response.data;
    } catch (error) {
      throw new Error(`Failed to register user: ${error}`);
    }
  }

  // Authenticate a user and return user data
  async authenticate(identifier: string, password: string): Promise<any> {
    try {
      const response = await axios.post(`${API_URL}authenticate`, { identifier, password });
      return {
        id: response.data.user_id, // Map user_id to id
        name: response.data.name,
        email: identifier, // Since backend doesn't return email, use the input identifier
      };
    } catch (error) {
      throw new Error(`Failed to authenticate user: ${error}`);
    }
  }

  // Get all users (for admin or display purposes)
  async getUsers(): Promise<User[]> {
    try {
      const response = await axios.get(API_URL);
      return response.data;
    } catch (error) {
      throw new Error(`Failed to fetch users: ${error}`);
    }
  }

  // Update user information
  async updateUser(id: number, user: User): Promise<void> {
    try {
      await axios.put(`${API_URL}${id}`, user);
    } catch (error) {
      throw new Error(`Failed to update user: ${error}`);
    }
  }

  // Delete a user
  async deleteUser(id: number): Promise<void> {
    try {
      await axios.delete(`${API_URL}${id}`);
    } catch (error) {
      throw new Error(`Failed to delete user: ${error}`);
    }
  }
  
  async getUserById(id: number): Promise<User> {
    try {
      const response = await axios.get(`${API_URL}${id}`);
      return response.data;
    } catch (error) {
      throw new Error(`Failed to fetch user: ${error}`);
    }
  }
  
}


export const userService = new UserService();
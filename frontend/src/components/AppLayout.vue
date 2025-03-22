<template>
  <div :class="['layout', theme]">
    <!-- Sidebar -->
    <aside class="sidebar" :class="{ collapsed: isSidebarCollapsed }">
      <!-- Sidebar Header -->
      <div class="sidebar-header">
        <slot name="sidebar-header">
          <h2>{{ appName }}</h2>
        </slot>
      </div>

      <!-- Menu Links -->
      <nav class="menu">
        <ul>
          <li v-for="(item, index) in menuItems" :key="index">
            <CustomLink :to="item.link" :label="item.label" />
          </li>
        </ul>
      </nav>

      <!-- Sidebar toggle button -->
      <button class="toggle-btn" @click="toggleSidebar">
        <span class="material-icons">
          {{ isSidebarCollapsed ? "chevron_right" : "chevron_left" }}
        </span>
      </button>
    </aside>

    <!-- Main Content Area -->
    <div class="main-container">
      <!-- Header -->
      <header class="header">
        <slot name="header"></slot>

        <!-- Dark Mode Toggle -->
        <button class="toggle-mode-btn" @click="toggleTheme">
          <span class="material-icons">
            {{ theme === "dark" ? "light_mode" : "dark_mode" }}
          </span>
        </button>
      </header>

      <!-- Main Content -->
      <main class="content">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import CustomLink from "./material/CustomLink.vue";

defineProps<{
  menuItems: { link: string; label: string }[];
  appName: string;
}>();

const isSidebarCollapsed = ref(false);
const theme = ref(localStorage.getItem("theme") || "light");

const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value;
};

const toggleTheme = () => {
  theme.value = theme.value === "dark" ? "light" : "dark";
  localStorage.setItem("theme", theme.value);
};
</script>
<style scoped>
@import url("https://fonts.googleapis.com/icon?family=Material+Icons");

/* Main Layout */
.layout {
  display: flex;
  height: 100vh;
  font-family: "Roboto", sans-serif;
  transition: background-color 0.3s, color 0.3s;
  overflow: hidden;
}

/* Sidebar */
.sidebar {
  background: #37474f;
  color: white;
  padding: 1rem;
  width: 250px;  /* Adjusted width for better fit */
  display: flex;
  flex-direction: column;
  transition: width 0.3s ease, background-color 0.3s ease;
  box-shadow: 2px 0 10px rgba(0, 0, 0, 0.15);
  border-radius: 0 8px 8px 0;
}

.sidebar.collapsed {
  width: 70px;
}

/* Sidebar Header */
.sidebar-header {
  font-weight: bold;
  font-size: 1.6rem;
  text-transform: uppercase;
  letter-spacing: 1px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 1.5rem;  /* Added margin for spacing */
}

/* Menu */
.menu {
  flex-grow: 1;
  overflow-y: auto;
}

.menu ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.menu li {
  margin-bottom: 1rem;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 12px;
  transition: background-color 0.3s ease, padding-left 0.2s ease;
}

.menu li:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.menu li.selected {
  background-color: #0288d1;
  color: white;
}

/* Sidebar toggle button */
.toggle-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 12px;
  margin-top: auto;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  align-self: center;
  transition: background 0.3s ease;
}

.toggle-btn:hover {
  background-color: rgba(255, 255, 255, 0.2);
}

/* Main container */
.main-container {
  flex-grow: 1;
  background-color: #eceff1;
  border-radius: 8px;
  margin-left: 15px;
  overflow: hidden;
  transition: background-color 0.3s ease;
}

/* Header */
.header {
  background: #0288d1;
  color: white;
  padding: 1rem 1.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: bold;
  border-radius: 8px 8px 0 0;
  transition: background-color 0.3s ease;

  position: sticky;  /* Makes the header sticky */
  top: 0;  /* Sticks the header to the top */
  z-index: 1000;  
}

/* Toggle Mode Button */
.toggle-mode-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  transition: background 0.3s ease;
}

.toggle-mode-btn:hover {
  background-color: rgba(255, 255, 255, 0.2);
}

/* Dark Mode Styles */
.layout.dark {
  background-color: #121212;
  color: white;
}

.layout.dark .sidebar {
  background-color: #333;
}

.layout.dark .main-container {
  background-color: #1e1e1e;
}

.layout.dark .header {
  background-color: #0288d1;
}

.layout.dark .menu li:hover {
  background-color: #444;
}

.layout.dark .toggle-btn {
  background-color: #37474f;
}

.layout.dark .menu li.selected {
  background-color: #0288d1;
}

/* Ensure proper scroll handling */
.layout .main-container, .menu {
  overflow-y: auto;  /* Ensures scroll works in both main content and sidebar */
}

.layout.dark .menu li:hover {
  background-color: #444;
}

</style>

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
        {{ isSidebarCollapsed ? "▶" : "◀" }}
      </button>
    </aside>

    <!-- Main Content Area -->
    <div class="main-container">
      <!-- Header -->
      <header class="header">
        <slot name="header"></slot>

        <!-- Dark Mode Toggle -->
        <!-- <button class="toggle-mode-btn" @click="toggleTheme">
          {{ theme === 'dark' ? 'Light Mode' : 'Dark Mode' }}
        </button> -->
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
import CustomLink from "../components/CustomLink.vue"; 

// Define props with proper types
defineProps<{
  menuItems: { link: string; label: string }[];
  appName: string;
}>();

// Reactive state for sidebar collapse
const isSidebarCollapsed = ref(false);

// Toggle the sidebar collapse state
const toggleSidebar = () => {
  isSidebarCollapsed.value = !isSidebarCollapsed.value;
};

// Reactive state for theme (dark/light mode)
const theme = ref(localStorage.getItem("theme") || "light");

// Toggle theme between light and dark mode
const toggleTheme = () => {
  theme.value = theme.value === "dark" ? "light" : "dark";
  localStorage.setItem("theme", theme.value); // Store theme preference in localStorage
};
</script>

<style scoped>
/* Main layout container */
.layout {
  display: flex;
  height: 100vh; /* Full viewport height */
  overflow: hidden;
  font-family: "Roboto", sans-serif; /* Material font */
  transition: background-color 0.3s, color 0.3s; /* Smooth transition for theme change */
}

/* Sidebar styles */
.sidebar {
  background: #263238; /* Dark grey-blue for material feel */
  color: white;
  padding: 1rem;
  width: 250px;
  display: flex;
  flex-direction: column;
  transition: width 0.3s ease, box-shadow 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* Subtle shadow */
  border-radius: 8px; /* Rounded corners */
}

.sidebar.collapsed {
  width: 60px;
  overflow: hidden;
  box-shadow: none; /* No shadow on collapsed */
}

/* Sidebar Header */
.sidebar-header {
  font-weight: bold;
  padding-bottom: 1rem;
  font-size: 1.5rem;
  text-transform: uppercase;
  letter-spacing: 1px;
}

/* Menu Links */
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
  padding: 0.5rem;
  border-radius: 4px;
  transition: background-color 0.3s ease, padding-left 0.3s ease;
}

.menu li:hover {
  background-color: #37474f; /* Darker grey-blue on hover */
  padding-left: 1rem; /* Subtle left padding on hover */
}

.menu li.selected {
  background-color: #0288d1; /* Material Blue for selected items */
  color: white;
}

/* Sidebar toggle button */
.toggle-btn {
  background: #37474f;
  border: none;
  color: white;
  cursor: pointer;
  padding: 10px;
  margin-top: auto;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease;
}

.toggle-btn:hover {
  background-color: #0288d1; /* Hover effect on button */
}

/* Main container styles */
.main-container {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  overflow: hidden;
  background-color: #eceff1; /* Light grey background for main content */
  box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.05); /* Slight inner shadow */
  border-radius: 8px; /* Rounded corners */
  margin-left: 15px; /* Space between sidebar and main content */
}

/* Header styles */
.header {
  background: #0288d1; /* Material Blue */
  color: white;
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: bold;
  border-top-left-radius: 0; /* No rounded top left corner */
  border-bottom-left-radius: 0; /* No rounded bottom left corner */
  border-radius: 8px 8px 0 0; /* Only top right and bottom right corners are rounded */
}

/* Dark Mode Styles */
.layout.dark {
  background-color: #121212;
  color: white;
}

.layout.dark .sidebar {
  background-color: #333;
  color: white;
}

.layout.dark .main-container {
  background-color: #1e1e1e;
}

.layout.dark .header {
  background-color: #0288d1;
}

.layout.dark .menu li:hover {
  background-color: #444; /* Darker grey-blue on hover in dark mode */
}

.layout.dark .toggle-mode-btn {
  background-color: #0288d1; /* Hover effect on dark mode button */
}

.layout.dark .toggle-btn {
  background-color: #37474f;
}

.layout.dark .menu li.selected {
  background-color: #0288d1; /* Selected item color in dark mode */
  color: white;
}

/* Light Mode Styles */
.layout.light {
  background-color: #eceff1;
  color: #333;
}

.layout.light .sidebar {
  background-color: #263238;
  color: white;
}

.layout.light .main-container {
  background-color: #ffffff;
}

.layout.light .header {
  background-color: #0288d1;
}

.layout.light .toggle-mode-btn {
  background-color: #0288d1; /* Hover effect on light mode button */
}

.layout.light .toggle-btn {
  background-color: #37474f;
}

.layout.light .menu li.selected {
  background-color: #0288d1; /* Selected item color in light mode */
  color: white;
}

/* Toggle button styles for dark and light mode */
.toggle-mode-btn {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 8px 16px;
  font-size: 1rem;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.toggle-mode-btn:hover {
  background-color: #0288d1;
}
</style>

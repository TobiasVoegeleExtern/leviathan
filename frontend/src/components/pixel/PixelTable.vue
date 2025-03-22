<template>
    <div class="pixel-table">
      <table class="table">
        <!-- Table Header -->
        <thead>
          <tr>
            <th v-for="column in columns" :key="column.field">
              {{ column.label }}
            </th>
          </tr>
        </thead>
  
        <!-- Table Body -->
        <tbody>
          <tr v-for="(row, index) in data" :key="index">
            <td v-for="column in columns" :key="column.field">
              {{ row[column.field] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script setup lang="ts">
  import { defineProps } from "vue";
  
  interface Column {
    label: string;
    field: string;
  }
  
  interface Row {
    [key: string]: any;
  }
  
  defineProps({
    columns: {
      type: Array as () => Column[],
      required: true,
    },
    data: {
      type: Array as () => Row[],
      required: true,
    },
  });
  </script>
  
  <style scoped>
  /* Pixel Table Container */
  .pixel-table {
    font-family: "Roboto Mono", monospace; /* More robotic and digital */
    background-color: #111;
    padding: 10px;
    border: 8px solid #000;
    box-shadow: 5px 5px 0px #444;
    display: inline-block;
}

  /* Table Styling */
  .table {
    width: 100%;
    border-collapse: collapse;
    background-color: #222;
  }
  
  /* Table Headers */
  th {
    background: #ffcc00;
    color: #222;
    font-size: 14px;
    padding: 16px;
    border: 4px solid #000;
    text-align: left;
  }
  
  /* Table Rows */
  td {
    color: white;
    font-size: 12px;
    padding: 12px;
    text-align: left;
    border: 4px solid #000;
  }
  
  /* Hover Effect */
  tbody tr:hover {
    background-color: #ff4444;
    color: black;
  }
  
  /* Pixelated Effect */
  .pixel-table,
  th,
  td {
    image-rendering: pixelated;
  }
  </style>
  
  
<template>
    <div class="material-table">
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
  import { defineProps } from 'vue';
  
  interface Column {
    label: string;
    field: string;
  }
  
  interface Row {
    [key: string]: any;
  }
  

  const props = defineProps({
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
  .material-table {
    font-family: 'Roboto', sans-serif;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    border-radius: 8px;
  }
  
  .table {
    width: 100%;
    border-collapse: collapse;
    background-color: #fff;
  }
  
  th,
  td {
    padding: 16px;
    text-align: left;
    border-bottom: 1px solid #eee;
  }
  
  th {
    background-color: #6200ea;
    color: white;
    font-weight: 600;
  }
  
  tbody tr:hover {
    background-color: #f1f1f1;
  }
  
  td {
    color: #333;
  }
  </style>
  
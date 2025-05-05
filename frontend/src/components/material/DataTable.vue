<template>
  <div class="material-table">
    <table class="table">
      <!-- Table Header -->
      <thead>
        <tr>
          <th v-for="column in columns" :key="column.field" @click="sortByColumn(column.field)">
            {{ column.label }}
            <span v-if="sortedColumn === column.field">
              <i v-if="sortDirection === 'asc'" class="sort-icon">↑</i>
              <i v-if="sortDirection === 'desc'" class="sort-icon">↓</i>
            </span>
          </th>
          <th>Aktion</th> <!-- New column for delete action -->
        </tr>
      </thead>

      <!-- Table Body -->
      <tbody>
        <tr v-for="(row, index) in sortedData" :key="index">
          <td v-for="column in columns" :key="column.field">
            {{ row[column.field] }}
          </td>
          <td>
            <button class="delete-btn" @click="openDeleteDialog(row)">
              🗑️ 
            </button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Delete Confirmation Modal -->
    <Dialogue 
  ref="dialogueRef"
  title="Eintrag löschen" 
  confirmText="Löschen"
  @confirm="confirmDelete" 
  @close="isDeleteDialogOpen = false"
>
  <p>Bist du sicher, dass du diesen Eintrag löschen möchtest?</p>
</Dialogue>

  </div>
</template>

<script setup lang="ts">
import { defineProps, ref, computed } from 'vue';
import Dialogue from "./Modal.vue"

interface Column {
  label: string;
  field: string;
}

interface Row {
  id: number; 
  [key: string]: any;
}

// Props
const props = defineProps({
  columns: {
    type: Array as () => Column[],
    required: true,
  },
  data: {
    type: Array as () => Row[],
    required: true,
  },
  onDelete: {
    type: Function,
    required: false, // Ensure a delete method is passed as a prop
  },
});
const dialogueRef = ref<InstanceType<typeof Dialogue> | null>(null);
// Sorting state
const sortedColumn = ref<string | null>(null);
const sortDirection = ref<'asc' | 'desc'>('asc');

// Computed property for sorted data
const sortedData = computed(() => {
  if (!sortedColumn.value) return props.data;

  return [...props.data].sort((a, b) => {
    const aValue = a[sortedColumn.value!];
    const bValue = b[sortedColumn.value!];

    if (typeof aValue === 'string' && typeof bValue === 'string') {
      return sortDirection.value === 'asc' 
        ? aValue.localeCompare(bValue) 
        : bValue.localeCompare(aValue);
    }
    if (typeof aValue === 'number' && typeof bValue === 'number') {
      return sortDirection.value === 'asc' ? aValue - bValue : bValue - aValue;
    }

    return 0;
  });
});

// Function to handle sorting
const sortByColumn = (columnField: string) => {
  if (sortedColumn.value === columnField) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortedColumn.value = columnField;
    sortDirection.value = 'asc';
  }
};

// Modal State
const isDeleteDialogOpen = ref(false);
const selectedRowId = ref<number | null>(null);

const openDeleteDialog = (row: Row) => {
  console.log("Row Data:", row); // Debugging
  selectedRowId.value = row.id ?? row.ID; // Support both `id` and `ID`
  console.log("Opening delete dialog for ID:", selectedRowId.value);

  if (!selectedRowId.value) {
    console.error("Error: Row ID is undefined!", row);
    return;
  }

  // Ensure modal opens
  dialogueRef.value?.openDialog();
};

const confirmDelete = () => {
  if (selectedRowId.value !== null) {
    console.log("Confirmed delete for ID:", selectedRowId.value);
    props.onDelete(selectedRowId.value);
  }
  isDeleteDialogOpen.value = false;
};
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
  cursor: pointer;
  user-select: none;
}

th:hover {
  background-color: #3700b3;
}

tbody tr:hover {
  background-color: #f1f1f1;
}

td {
  color: #333;
}

.sort-icon {
  margin-left: 8px;
  font-size: 14px;
  color: #fff;
}

.delete-btn {
  background: none; 
  border: none; 
  padding: 0;
  cursor: pointer; 
  font-size: 14px; 
}

.delete-btn:hover {
  background-color: #b76f1c;
}
</style>

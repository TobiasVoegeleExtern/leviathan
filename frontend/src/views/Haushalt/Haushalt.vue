<template>
  <p></p>
  <Button label="Ausgabe hinzufügen" :execute="toggleFormVisibility" />
  <Card
    title="Haushaltsbuch"
    :onClick="handleClick"
    buttonText="Discover More"
    :fullWidth="true"
  >
    <template v-slot:description>
      <div class="home">
        <p>User: {{ userStore.user?.name }} (ID: {{ userStore.user?.id }})</p>
        <DatePicker @update:selectedDate="handleSelectedDate" />

        <p></p>
        <div v-if="!showForm">
          <DataTable :columns="columns" :data="rows" :onDelete="deleteRow" />
          <p><strong>Ausgaben gesamt diesen Monat: {{ totalMonthlyRate }} Euro</strong></p>
        </div>
        <Form v-if="showForm" @submit="fetchExpenses" />
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from "vue";
import Form from "./AusgabenForm.vue";
import Button from "../../components/material/Button.vue";
import DataTable from "../../components/material/DataTable.vue";
import Card from "../../components/material/Card.vue";
import HouseholdExpensesService from "../../services/HouseholdExpensesService";
import { useUserStore } from '../../stores/userStore';
import DatePicker from '../../components/material/DatePicker.vue';

const showForm = ref(false);
const userStore = useUserStore();
const rows = ref([]);
const selectedDate = ref<Date | null>(null);

const formattedDate = computed(() => {
  return selectedDate.value
    ? `${selectedDate.value.getFullYear()}-${(selectedDate.value.getMonth() + 1).toString().padStart(2, '0')}`
    : '';
});

const toggleFormVisibility = () => {
  showForm.value = !showForm.value;
};

const columns = ref([
  { label: "Bezeichnung", field: "Description" },
  { label: "Komplettbetrag", field: "ValueTotal" },
  { label: "Monatliche Rate", field: "ValueRate" },
  { label: "Kreditstart", field: "CreditStart" },
  { label: "Kreditende", field: "CreditEnd" },
  { label: "Faelligkeitstag", field: "Faelligkeitstag" },
  { label: "Rechnungsdatum", field: "Zahldatum" },
  { label: "Typ", field: "Type" },
]);

const formatValue = (value: any, field: string, row?: any) => {
  if (!value || value === '0001-01-01' || value === '1.1.1' || value === '0001-01-01T00:00:00Z') {
    return '';
  }

  if (typeof value === "string" && value.includes("T")) {
    const [year, month, day] = value.split("T")[0].split("-");
    return `${day}.${month}.${year}`;
  }

  if (field === "Faelligkeitstag") {
    if (row?.Zahldatum && row.Zahldatum !== '0001-01-01T00:00:00Z') {
      return formatValue(row.Zahldatum, "Zahldatum");
    }

    if (!isNaN(parseInt(value))) {
      const currentDate = new Date();
      const currentYear = currentDate.getFullYear();
      const currentMonth = (currentDate.getMonth() + 1).toString().padStart(2, '0');
      const day = parseInt(value).toString().padStart(2, '0');

      return `${day}.${currentMonth}.${currentYear}`;
    }
  }

  return value;
};

const fetchExpenses = async () => {
  try {
    const userId = userStore.user?.id;
    if (!userId) return;

    const formatted = formattedDate.value || new Date().toISOString().slice(0, 7);
    const data = await HouseholdExpensesService.getExpensesByUserAndMonth(userId, formatted);
    console.log("API Response:", data);

    rows.value = data.map((row: any) => ({
      ID: row.ID ?? row.id ?? null,
      Description: row.Description,
      ValueTotal: formatValue(row.ValueTotal, "ValueTotal"),
      ValueRate: formatValue(row.ValueRate, "ValueRate"),
      CreditStart: formatValue(row.CreditStart, "CreditStart"),
      CreditEnd: formatValue(row.CreditEnd, "CreditEnd"),
      Type: formatValue(row.Type, "Type"),
      Zahldatum: formatValue(row.Zahldatum, "Zahldatum"),
      Faelligkeitstag: formatValue(row.Faelligkeitstag, "Faelligkeitstag", row), 
    }));
  } catch (error) {
    console.error("Error fetching user expenses:", error);
  }
};

const handleSelectedDate = (date: Date) => {
  selectedDate.value = date;
  console.log('Selected Date:', formattedDate.value);
};

watch(selectedDate, fetchExpenses);
onMounted(fetchExpenses);

const handleClick = () => {
  alert("hi");
};

const totalMonthlyRate = computed(() => {
  return rows.value.reduce((sum, row) => {
    const valueRate = parseFloat(row.ValueRate);
    return !isNaN(valueRate) ? sum + valueRate : sum;
  }, 0).toFixed(2);
});

const deleteRow = async (ID: number) => {
  try {
    if (!ID) {
      console.error("Invalid ID:", ID);
      return;
    }

    await HouseholdExpensesService.deleteExpense(ID);
    rows.value = rows.value.filter(row => row.ID !== ID);
    console.log(`Deleted row with ID: ${ID}`);
  } catch (error) {
    console.error("Error deleting row:", error);
  }
};
</script>

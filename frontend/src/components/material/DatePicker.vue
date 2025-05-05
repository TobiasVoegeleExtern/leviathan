<template>
    <div class="datepicker-container" ref="datepickerContainer">
      <!-- Button für die Datumswahl -->
      <button
        class="material-button"
        @click="toggleCalendar"
        :aria-label="'Monat und Jahr auswählen'"
      >
        {{ formattedMonthYear }}
      </button>
  
      <!-- Kalender-Popup -->
      <div v-if="showCalendar" class="material-calendar" :style="calendarPosition">
        <div class="calendar-header">
          <button @click="changeYear(-1)"><<</button>
          <button @click="changeMonth(-1)"><</button>
          <span class="month-year" @click="showYearList">{{ currentMonthYear }}</span>
          <button @click="changeMonth(1)">></button>
          <button @click="changeYear(1)">>></button>
        </div>
  
        <div class="calendar-body">
          <div class="month-grid">
            <div
              v-for="(month, index) in months"
              :key="index"
              @click="selectMonth(index)"
              :class="{ selected: isSelectedMonth(index) }"
            >
              {{ month }}
            </div>
          </div>
  
          <!-- Jahr-Auswahl -->
          <div v-if="showYearSelection" class="year-selection">
            <div
              v-for="year in yearRange"
              :key="year"
              @click="selectYear(year)"
              :class="{ selected: year === currentDate.getFullYear() }"
            >
              {{ year }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue';
  
  /**
   * Definiert die Eingabe-Eigenschaften für den Datumswähler.
   * 
   * @prop {boolean} restrictToYearMonth - Beschränkt den Wähler auf Jahr und Monat. Standardmäßig auf true gesetzt.
   */
  const props = defineProps({
    restrictToYearMonth: {
      type: Boolean,
      default: true, // Standardmäßig auf true gesetzt, d.h., nur Jahr und Monat können ausgewählt werden.
    },
  });
  
  /**
   * Definiert das Emit für das Senden des ausgewählten Datums.
   * 
   * @event 'update:selectedDate' - Sendet das ausgewählte Datum an den übergeordneten Bestandteil.
   */
  const emit = defineEmits<{
    (event: 'update:selectedDate', date: Date): void;
  }>();
  
  const showCalendar = ref(false); // Steuerung der Sichtbarkeit des Kalenders
  const showYearSelection = ref(false); // Umschalten der Sichtbarkeit der Jahr-Auswahl
  const selectedDate = ref<Date | null>(null); // Speichert das ausgewählte Datum
  const currentDate = ref(new Date()); // Verfolgt das aktuelle Datum für die Monats/Jahresnavigation
  
  // Deutsche Monatsnamen
  const months = ref([
    'Januar', 'Februar', 'März', 'April', 'Mai', 'Juni',
    'Juli', 'August', 'September', 'Oktober', 'November', 'Dezember',
  ]);
  
  /**
   * Berechnet die Jahr-Spanne für die Jahr-Auswahl (5 Jahre vor und nach dem aktuellen Jahr).
   */
  const yearRange = computed(() => {
    const currentYear = currentDate.value.getFullYear();
    const startYear = currentYear - 5; // Zeigt 5 Jahre vor dem aktuellen Jahr an
    const endYear = currentYear + 5; // Zeigt 5 Jahre nach dem aktuellen Jahr an
    return Array.from({ length: endYear - startYear + 1 }, (_, i) => startYear + i);
  });
  
  /**
   * Gibt das ausgewählte Datum im Format "Monat Jahr" zurück (in deutscher Sprache).
   * 
   * @returns {string} Das formatierte Datum.
   */
  const formattedMonthYear = computed(() => {
    if (selectedDate.value) {
      const month = months.value[selectedDate.value.getMonth()];
      const year = selectedDate.value.getFullYear();
      return `${month} ${year}`;
    }
    return 'Monat wählen'; // Platzhaltertext, wenn kein Datum ausgewählt ist
  });
  
  /**
   * Gibt das aktuelle Monat und Jahr im Format "Monat Jahr" zurück (in deutscher Sprache).
   * 
   * @returns {string} Das aktuelle Monat und Jahr.
   */
  const currentMonthYear = computed(() => {
    const month = months.value[currentDate.value.getMonth()];
    const year = currentDate.value.getFullYear();
    return `${month} ${year}`;
  });
  
  /**
   * Überprüft, ob ein Monat ausgewählt ist.
   * 
   * @param {number} monthIndex - Der Index des Monats.
   * @returns {boolean} True, wenn der Monat ausgewählt ist, andernfalls false.
   */
  const isSelectedMonth = (monthIndex: number) => {
    return (
      selectedDate.value &&
      selectedDate.value.getFullYear() === currentDate.value.getFullYear() &&
      selectedDate.value.getMonth() === monthIndex
    );
  };
  
  /**
   * Umschaltet die Sichtbarkeit des Kalenders.
   */
  const toggleCalendar = () => {
    showCalendar.value = !showCalendar.value;
    if (!showCalendar.value) {
      showYearSelection.value = false; // Setzt die Jahr-Auswahl zurück, wenn der Kalender geschlossen wird
    }
  };
  
  /**
   * Ändert den aktuellen Monat um den angegebenen Offset (-1 für den vorherigen Monat, 1 für den nächsten Monat).
   * 
   * @param {number} offset - Der Offset, um den Monat zu ändern.
   */
  const changeMonth = (offset: number) => {
    const newDate = new Date(currentDate.value);
    newDate.setMonth(newDate.getMonth() + offset);
    currentDate.value = newDate; // Re-Assign wird verwendet, um die Reaktivität auszulösen
  };
  
  /**
   * Ändert das aktuelle Jahr um den angegebenen Offset (-1 für das vorherige Jahr, 1 für das nächste Jahr).
   * 
   * @param {number} offset - Der Offset, um das Jahr zu ändern.
   */
  const changeYear = (offset: number) => {
    const newDate = new Date(currentDate.value);
    newDate.setFullYear(newDate.getFullYear() + offset);
    currentDate.value = newDate; // Re-Assign wird verwendet, um die Reaktivität auszulösen
  };
  
  /**
   * Wählt einen Monat aus und setzt das ausgewählte Datum.
   * 
   * @param {number} monthIndex - Der Index des ausgewählten Monats.
   */
  const selectMonth = (monthIndex: number) => {
    selectedDate.value = new Date(currentDate.value.getFullYear(), monthIndex, 1);
    showCalendar.value = false; // Schließt den Kalender nach der Auswahl eines Monats
  
    // Sendet das ausgewählte Datum an den übergeordneten Bestandteil
    if (selectedDate.value) {
      emit('update:selectedDate', selectedDate.value);
    }
  };
  
  /**
   * Zeigt die Jahr-Auswahl an.
   */
  const showYearList = () => {
    showYearSelection.value = !showYearSelection.value;
  };
  
  /**
   * Wählt ein Jahr aus und aktualisiert das ausgewählte Datum.
   * 
   * @param {number} year - Das ausgewählte Jahr.
   */
  const selectYear = (year: number) => {
    const newDate = new Date(currentDate.value);
    newDate.setFullYear(year);
    currentDate.value = newDate; // Aktualisiert currentDate, um das ausgewählte Jahr widerzuspiegeln
    if (selectedDate.value) {
      selectedDate.value = new Date(year, selectedDate.value.getMonth(), 1); // Aktualisiert das ausgewählte Datum
    }
    showYearSelection.value = false; // Schließt die Jahr-Auswahl nach der Auswahl eines Jahres
  
    // Sendet das ausgewählte Datum an den übergeordneten Bestandteil
    if (selectedDate.value) {
      emit('update:selectedDate', selectedDate.value);
    }
  };
  
  /**
   * Berechnet die dynamische Position des Kalenders, damit er unter dem Button erscheint.
   * 
   * @returns {Object} Die Position des Kalenders mit 'top' und 'left' Eigenschaften.
   */
  const calendarPosition = computed(() => {
    const datepickerContainer = document.querySelector('.datepicker-container') as HTMLElement;
    if (!datepickerContainer) return {};
  
    const rect = datepickerContainer.getBoundingClientRect();
    return {
      top: `${rect.bottom + window.scrollY}px`,
      left: `${rect.left + window.scrollX}px`,
    };
  });
  </script>
  
  <style scoped>
  .material-button {
    padding: 12px;
    border-radius: 4px;
    font-size: 1rem;
    border: 1px solid #6200ea;
    background-color: white;
    color: #6200ea;
    width: 100%;
    box-sizing: border-box;
    text-align: left;
    cursor: pointer;
    transition: background-color 0.3s ease;
    width :28%;
  }
  
  .material-button:hover {
    background-color: #f1f1f1;
  }
  
  .material-calendar {
    position: absolute;
    background-color: white;
    border-radius: 8px;
    box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.2);
    width: 350px; 
    padding: 10px;
    z-index: 9999;
  }
  
  .calendar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
  }
  
  .calendar-header button {
    background-color: transparent;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
  }
  
  .calendar-header .month-year {
    font-size: 1.2rem;
    font-weight: bold;
    color: #6200ea;
    cursor: pointer;
  }
  
  .month-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }
  
  .month-grid div {
    padding: 8px;
    text-align: center;
    background-color: #f1f1f1;
    cursor: pointer;
    border-radius: 4px;
    transition: background-color 0.3s ease;
  }
  
  .month-grid div:hover {
    background-color: #6200ea;
    color: white;
  }
  
  .month-grid div.selected {
    background-color: #6200ea;
    color: white;
  }
  
  .year-selection {
    margin-top: 10px;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;
  }
  
  .year-selection div {
    padding: 8px;
    text-align: center;
    background-color: #f1f1f1;
    cursor: pointer;
    border-radius: 4px;
    transition: background-color 0.3s ease;
  }
  
  .year-selection div:hover,
  .year-selection div.selected {
    background-color: #6200ea;
    color: white;
  }
  </style>
  
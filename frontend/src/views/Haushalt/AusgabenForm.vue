<template>
  <form @submit.prevent="submitForm" class="form-wrapper">
    <Dropdown
      :options="dropdownOptions"
      v-model="selected"
      @update:selected-option="handleSelection"
      class="dropdown-wrapper"
    />

    <p v-if="formError" class="error-message">{{ formError }}</p>

    <TextInput
      label="Bezeichnung"
      inputType="text"
      v-model="inputText"
      :showError="showValidationError && !isBezeichnungValid"
      errorMessage="Bezeichnung ist ein Pflichtfeld"
      class="input-with-margin"
    />

    <TextInput
      label="Betrag"
      inputType="text"
      v-model="betragInput"
      :showError="showValidationError && !isBetragValid"
      errorMessage="Bitte eine gültige positive Zahl eingeben"
    />

    <TextInput
      v-if="isFaelligkeitsdatumVisible"
      label="Tag der Faelligkeit (1-31)"
      v-model="inputFaelligkeit" inputType="number"
      min="1"
      max="31"
      :showError="showValidationError && !isFaelligkeitsdatumValid"
      errorMessage="Gültigen Tag (1-31) eingeben"
    />

    <TextInput
      v-if="isKrediteSelected"
      label="Kreditbeginn"
      inputType="date"
      v-model="abbezahltInput"
      :showError="showValidationError && !isAbbezahltValid"
      errorMessage="Gültiges Datum für Kreditbeginn eingeben"
    />

    <TextInput
      v-if="isKrediteSelected"
      label="Kreditende"
      inputType="date"
      v-model="zahltagInput"
      :showError="showValidationError && !isZahltagValid"
      errorMessage="Gültiges Datum für Kreditende eingeben"
    />

    <button type="submit" :disabled="!isFormValid || isSubmitting" class="submit-button-style">
      {{ isSubmitting ? 'Sende...' : 'Submit' }} </button>

  </form>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import TextInput from '../../components/material/TextInput.vue'; // Pfad anpassen
// Button Komponente nicht mehr zwingend nötig für Submit, wenn Standard-Button verwendet wird
// import Button from '../../components/material/Button.vue'; // Pfad anpassen (optional)
import Dropdown from "../../components/material/DropdownSelect.vue"; // Pfad anpassen
import { useUserStore } from '../../stores/userStore'; // Pfad anpassen
import householdExpensesService from '../../services/HouseholdExpensesService'; // Pfad anpassen
import type { HouseholdExpenseFormData } from '../../services/HouseholdExpensesService'; // Typ-Import

interface Option {
  label: string;
  value: string;
}

// --- State Refs ---
const formError = ref<string | null>(null); // Fehlermeldung für das Formular
const inputText = ref(''); // Für Bezeichnung
const betragInput = ref('');
const inputFaelligkeit = ref(''); // Separates Ref für Fälligkeitstag
const abbezahltInput = ref(''); // Für Kreditbeginn
const zahltagInput = ref('');   // Für Kreditende
const selected = ref<Option | null>(null); // Für Dropdown-Auswahl
const isSubmitting = ref(false); // NEU: Status, ob gerade gesendet wird
const showValidationError = ref(false); // NEU: Steuert, wann Validierungsfehler angezeigt werden

// User Store
const userStore = useUserStore();
const userId = computed(() => userStore.user?.id);

// Dropdown Optionen
const dropdownOptions: Option[] = [
  { label: "Sonstiges", value: "allelse" },
  { label: "Kredite", value: "credit" },
  { label: "Laufende Kosten", value: "monthlycosts" },
];

// --- Validation Computed Properties ---
const isBezeichnungValid = computed(() => inputText.value.trim() !== '');
const isBetragValid = computed(() => {
  const num = parseFloat(betragInput.value);
  return !isNaN(num) && betragInput.value.trim() !== '' && num > 0;
});
const isFaelligkeitsdatumValid = computed(() => {
  if (!isFaelligkeitsdatumVisible.value) return true;
  return /^[1-9]$|^[12][0-9]$|^3[01]$/.test(inputFaelligkeit.value.trim());
});
const isAbbezahltValid = computed(() => {
  if (!isKrediteSelected.value) return true;
  return abbezahltInput.value.trim() !== '';
});
const isZahltagValid = computed(() => {
  if (!isKrediteSelected.value) return true;
  return zahltagInput.value.trim() !== '';
});

// --- Visibility Computed Properties ---
const isFaelligkeitsdatumVisible = computed(() => selected.value?.value === "credit" || selected.value?.value === "monthlycosts");
const isKrediteSelected = computed(() => selected.value?.value === "credit");

// --- Gesamtvalidierung ---
const isFormValid = computed(() => {
  let commonValid = isBezeichnungValid.value &&
                    isBetragValid.value &&
                    selected.value !== null &&
                    userId.value !== null && userId.value !== undefined;
  if (!commonValid) return false;
  const selectedType = selected.value?.value;
  if (selectedType === 'credit') {
    return isFaelligkeitsdatumValid.value && isAbbezahltValid.value && isZahltagValid.value;
  } else if (selectedType === 'monthlycosts') {
    return isFaelligkeitsdatumValid.value;
  } else if (selectedType === 'allelse') {
    return true;
  }
  return false;
});

// --- Event Handlers ---

// Wird vom @submit.prevent des Formulars aufgerufen
const submitForm = async () => {
  const submitId = Date.now();
  console.log(`submitForm: Funktion aufgerufen! (ID: ${submitId})`);
  showValidationError.value = true; // Aktiviere Anzeige von Validierungsfehlern

  // Verhindere erneute Ausführung, wenn schon gesendet wird
  if (isSubmitting.value) {
    console.log(`submitForm: Bereits am Senden (ID: ${submitId}), Abbruch.`);
    return;
  }

  // Validierung prüfen
  if (!isFormValid.value) {
    formError.value = "Bitte überprüfe die markierten Eingabefelder.";
    console.log(`submitForm: Formular ungültig (ID: ${submitId}), Abbruch.`);
    return; // Abbruch, wenn ungültig
  }

  // User ID prüfen (redundant wegen isFormValid, aber schadet nicht)
  if (!userId.value) {
      formError.value = "Benutzer-ID nicht gefunden. Bitte neu anmelden.";
      console.error("submitForm: User ID ist null oder undefined!");
      return;
  }

  // --- Senden beginnt ---
  isSubmitting.value = true; // Button deaktivieren, erneuten Aufruf verhindern
  formError.value = null;    // Alte Fehler löschen

  try {
    // Erstelle Payload mit korrekten Namen für den Service
    const payload: HouseholdExpenseFormData = {
      user_id: userId.value,
      description: inputText.value.trim() || undefined,
      valuetotal: parseFloat(betragInput.value),
      type: selected.value!.value, // Typ ist sicher gesetzt (geprüft in isFormValid)
      faelligkeitstag: (isFaelligkeitsdatumVisible.value && inputFaelligkeit.value.trim()) ? inputFaelligkeit.value.trim() : undefined,
      creditstart: (isKrediteSelected.value && abbezahltInput.value) ? abbezahltInput.value : undefined,
      creditend: (isKrediteSelected.value && zahltagInput.value) ? zahltagInput.value : undefined,
    };

    console.log(`submitForm: Sende Payload (ID: ${submitId})`, payload);

    // Service aufrufen
    const response = await householdExpensesService.submitExpense(payload);

    console.log(`submitForm: Antwort vom Backend (ID: ${submitId})`, response);
    alert('Eintrag erfolgreich gespeichert!'); // Oder bessere UI-Benachrichtigung
    resetForm(); // Formular zurücksetzen

  } catch (error: any) {
    console.error(`submitForm: Fehler beim Absenden! (ID: ${submitId})`, error);
    formError.value = error.message || 'Es gab einen Fehler beim Absenden des Formulars.';
    showValidationError.value = true; // Stelle sicher, dass Fehler angezeigt werden
  } finally {
    // --- Senden beendet (Erfolg oder Fehler) ---
    isSubmitting.value = false; // Button wieder aktivieren
    console.log(`submitForm: Verarbeitung beendet (ID: ${submitId})`);
  }
};

// Formular zurücksetzen
const resetForm = () => {
    inputText.value = '';
    betragInput.value = '';
    inputFaelligkeit.value = ''; // Zurücksetzen hinzugefügt
    abbezahltInput.value = '';
    zahltagInput.value = '';
    selected.value = null;
    formError.value = null;
    isSubmitting.value = false; // Sicherstellen, dass Senden-Status zurückgesetzt wird
    showValidationError.value = false; // Validierungsfehler ausblenden
};

// Handle selection in the dropdown
const handleSelection = (option: Option) => {
  selected.value = option;
  console.log('Selected option:', option);
  formError.value = null; // Fehler bei Typwechsel löschen
  showValidationError.value = false; // Validierungsfehler ausblenden
  // Optional: Setze typspezifische Felder zurück
  if (!isFaelligkeitsdatumVisible.value) inputFaelligkeit.value = '';
  if (!isKrediteSelected.value) {
    abbezahltInput.value = '';
    zahltagInput.value = '';
  }
};

</script>

<style scoped>
/* Dein CSS von vorher, plus optional Styles für Standard-Button */
.form-wrapper {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  max-width: 500px;
  margin: 0 auto;
}

.error-message {
  color: #e74c3c;
  font-size: 0.875rem;
  margin-top: 0.25rem;
  text-align: center;
  min-height: 1.2em;
}

.input-with-margin {
  margin-top: 1rem;
}

/* Beispiel-Styles für Standard-Button */
.submit-button-style {
    padding: 0.8rem 1.5rem;
    border: none;
    background-color: #3498db; /* Blau */
    color: white;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    transition: background-color 0.3s ease, opacity 0.3s ease;
    opacity: 1;
}

.submit-button-style:hover:not(:disabled) {
    background-color: #2980b9; /* Dunkleres Blau */
}

.submit-button-style:disabled {
    background-color: #bdc3c7; /* Grau */
    cursor: not-allowed;
    opacity: 0.7;
}

/* Styles für Dropdown etc. */
.dropdown-wrapper {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  width: 100%;
  position: relative;
  z-index: 1000;
}

.dropdown-wrapper select { /* Falls deine Dropdown-Komponente ein Select rendert */
  width: 100%;
  padding: 0.5rem;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.dropdown-wrapper option {
  width: 100%;
  padding: 0.5rem;
}
</style>
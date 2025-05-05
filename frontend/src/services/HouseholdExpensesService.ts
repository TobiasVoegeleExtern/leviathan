import axios from 'axios';

// Interface definiert die erwarteten Daten (muss mit Vue Payload übereinstimmen)
export interface HouseholdExpenseFormData {
  user_id: number;          // Wird im Backend zu 'userid' gemappt (mit json-Tag)
  description?: string;
  valuetotal: number;
  type: string;
  faelligkeitstag?: string; // Name muss mit Backend-Model/JSON übereinstimmen
  creditstart?: string;    // Name muss mit Backend-Model/JSON übereinstimmen
  creditend?: string;      // Name muss mit Backend-Model/JSON übereinstimmen
}

class HouseholdExpensesService {
  private apiUrl: string;

  constructor(apiUrl: string) {
    this.apiUrl = apiUrl;
  }

  /**
   * Zentrale Methode zum Einreichen von Ausgaben.
   * Leitet basierend auf dem Typ an spezifische interne Methoden weiter.
   */
  async submitExpense(formData: HouseholdExpenseFormData) {
    // Grundlegende Validierung im Service (optional, da schon im Frontend)
    if (!formData.user_id || formData.valuetotal === undefined || formData.valuetotal === null || !formData.type) {
      console.error("Service Validation Failed:", formData);
      throw new Error('User ID, Betrag, and Type are required in Service.');
    }

    console.log("Received formData in service:", formData);

    try {
      // Sende direkt das aufbereitete formData-Objekt
      // Die Typ-spezifische Logik ist nun im Frontend (payload-Erstellung)
      // oder könnte hier zur doppelten Prüfung/Aufbereitung dienen.
      // Einfachere Variante: Direkt posten, was vom Frontend kommt.
      return this.postExpenseData(formData);

      // Alte Variante mit interner Aufteilung (kann entfernt werden, wenn Payload im Frontend passt):
      /*
      switch (formData.type) {
        case 'credit':
          return this.submitCreditExpense(formData);
        case 'monthlycosts':
          return this.submitMonthlyCostsExpense(formData);
        case 'allelse':
          return this.submitAllelseExpense(formData);
        default:
          console.warn(`Unknown expense type encountered: ${formData.type}. Submitting as 'allelse'.`);
          return this.submitAllelseExpense(formData);
      }
      */
    } catch (error: any) {
      console.error('Error submitting household expense:', error);
      const message = error.response?.data?.message || error.message || 'Failed to submit household expense';
      throw new Error(message);
    }
  }

  // --- Interne Submit-Methoden sind jetzt weniger notwendig, ---
  // --- wenn das Payload-Objekt im Frontend korrekt erstellt wird ---
  // --- Man könnte sie als private Hilfsfunktionen behalten oder entfernen ---

  // Beispiel: Behalte sie zur Verdeutlichung oder für spezielle Logik
  private async submitCreditExpense(formData: HouseholdExpenseFormData) {
    const payload: any = {
        // Verwende Kleinbuchstaben für JSON Keys, wenn Backend das erwartet (mit json Tag im Go Model!)
        userid: formData.user_id, // Map zu UserID mit `json:"userid"`
        valuetotal: formData.valuetotal,
        type: 'credit'
    };
    if (formData.description) payload.description = formData.description;
    if (formData.faelligkeitstag) payload.faelligkeitstag = formData.faelligkeitstag; // Name konsistent halten
    if (formData.creditstart) payload.creditstart = new Date(formData.creditstart).toISOString();
    if (formData.creditend) payload.creditend = new Date(formData.creditend).toISOString();

    console.log("Sending 'credit' payload (JSON):", payload);
    return this.postExpenseData(payload);
  }

  private async submitMonthlyCostsExpense(formData: HouseholdExpenseFormData) {
     const payload: any = {
        userid: formData.user_id,
        valuetotal: formData.valuetotal,
        type: 'monthlycosts'
    };
    if (formData.description) payload.description = formData.description;
    if (formData.faelligkeitstag) payload.faelligkeitstag = formData.faelligkeitstag; // Name konsistent halten

    console.log("Sending 'monthlycosts' payload (JSON):", payload);
    return this.postExpenseData(payload);
  }

   private async submitAllelseExpense(formData: HouseholdExpenseFormData) {
     const payload: any = {
        userid: formData.user_id,
        valuetotal: formData.valuetotal,
        type: 'allelse'
    };
    if (formData.description) payload.description = formData.description;

    console.log("Sending 'allelse' payload (JSON):", payload);
    return this.postExpenseData(payload);
  }


  /**
   * Hilfsmethode zum tatsächlichen Senden der Daten per POST
   * Nimmt jetzt direkt das payload-Objekt entgegen.
   */
  private async postExpenseData(payload: HouseholdExpenseFormData) {
     // Stelle sicher, dass user_id zu userid wird, falls Backend das erwartet
     const backendPayload: any = { ...payload, userid: payload.user_id };
     delete backendPayload.user_id; // Entferne das Originalfeld

     // Konvertiere Daten im ISO Format, wenn nötig (bereits im Frontend passiert?)
     if (backendPayload.creditstart && typeof backendPayload.creditstart === 'string') {
         backendPayload.creditstart = new Date(backendPayload.creditstart).toISOString();
     }
      if (backendPayload.creditend && typeof backendPayload.creditend === 'string') {
         backendPayload.creditend = new Date(backendPayload.creditend).toISOString();
     }


     console.log("Final payload being sent:", backendPayload);

     try {
        // Sende das aufbereitete Backend-Payload
        const response = await axios.post(`${this.apiUrl}/haushaltsausgaben/`, backendPayload, {
            headers: { 'Content-Type': 'application/json' }
        });
        return response.data;
     } catch (error: any) {
        console.error(`Fehler beim Senden der Daten an ${this.apiUrl}/haushaltsausgaben/`, error);
        // Detailliertere Fehlermeldung versuchen zu extrahieren
        let errorMessage = 'Senden der Ausgabe fehlgeschlagen.';
        if (error.response) {
          // Server hat geantwortet, aber mit Fehlerstatus (4xx, 5xx)
          console.error('Error Response Data:', error.response.data);
          console.error('Error Response Status:', error.response.status);
          console.error('Error Response Headers:', error.response.headers);
          // Versuche, eine Fehlermeldung aus der Antwort zu bekommen
          if (error.response.data && error.response.data.error) {
            errorMessage = error.response.data.error;
          } else if (typeof error.response.data === 'string' && error.response.data.length < 200) {
            // Manchmal ist der Fehlertext direkt im Body
             errorMessage = error.response.data;
          } else {
            errorMessage = `Server responded with status ${error.response.status}`;
          }
        } else if (error.request) {
          // Request wurde gemacht, aber keine Antwort erhalten
          console.error('Error Request:', error.request);
          errorMessage = 'Keine Antwort vom Server erhalten.';
        } else {
          // Fehler beim Erstellen des Requests
          console.error('Error Message:', error.message);
          errorMessage = error.message;
        }
        throw new Error(errorMessage); // Gib die (hoffentlich) spezifischere Fehlermeldung weiter
     }
  }


  // --- Methoden für GET, UPDATE, DELETE bleiben unverändert ---

  async getExpenses() {
    try {
      const response = await axios.get(`${this.apiUrl}/haushaltsausgaben/`);
      return response.data;
    } catch (error: any) {
      console.error('Error fetching expenses:', error);
      const message = error.response?.data?.message || error.message || 'Failed to fetch expenses';
      throw new Error(message);
    }
  }

  async getExpensesByUserAndMonth(userid: number, month: string) {
     try {
      const response = await axios.get(`${this.apiUrl}/haushaltsausgaben/${userid}/${month}`);
      const expenses = response.data.map((expense: any) => {
        if (expense.CreditStart) expense.CreditStart = this.formatDate(expense.CreditStart);
        if (expense.CreditEnd) expense.CreditEnd = this.formatDate(expense.CreditEnd);
        // Passe ggf. an, wie das Backend das Datumsfeld zurückgibt
        if (expense.Zahltag) expense.Zahltag = this.formatDate(expense.Zahltag);
        else if (expense.zahldatum) expense.Zahltag = this.formatDate(expense.zahldatum);
        return expense;
      });
      return expenses;
    } catch (error: any) {
      console.error('Error fetching user expenses:', error);
      const message = error.response?.data?.message || error.message || 'Failed to fetch user expenses';
      throw new Error(message);
    }
  }

  formatDate(dateString: string | null | undefined): string | null {
     if (!dateString) return null;
     try {
        const date = new Date(dateString);
        if (isNaN(date.getTime())) return dateString; // Kein gültiges Datum, gib Original zurück
        const formattedDate = new Intl.DateTimeFormat('de-DE', { year: 'numeric', month: '2-digit', day: '2-digit' }).format(date);
        return formattedDate;
     } catch (e) {
        console.error("Error formatting date:", dateString, e);
        return dateString; // Gib Original im Fehlerfall zurück
     }
  }

  async updateExpense(id: number, formData: Partial<HouseholdExpenseFormData>) { // Erlaube partielle Updates
     // Sende nur die Felder, die im formData vorhanden sind
     const payload: any = { };
     if (formData.description !== undefined) payload.description = formData.description || null; // Leeren String als null?
     if (formData.valuetotal !== undefined) payload.valuetotal = formData.valuetotal;
     // Füge hier weitere update-bare Felder hinzu, falls nötig und vom Backend unterstützt
     // if (formData.faelligkeitstag !== undefined) payload.faelligkeitstag = formData.faelligkeitstag;
     // ...

    try {
      const response = await axios.put(`${this.apiUrl}/haushaltsausgaben/${id}`, payload, {
        headers: { 'Content-Type': 'application/json' }
      });
      return response.data;
    } catch (error: any) {
      console.error('Error updating expense:', error);
      const message = error.response?.data?.message || error.message || 'Failed to update expense';
      throw new Error(message);
    }
  }

  async deleteExpense(id: number) {
    try {
      const response = await axios.delete(`${this.apiUrl}/haushaltsausgaben/${id}`);
      return response.data;
    } catch (error: any) {
      console.error('Error deleting expense:', error);
      const message = error.response?.data?.message || error.message || 'Failed to delete expense';
      throw new Error(message);
    }
  }
}

// Stelle sicher, dass die URL korrekt ist und dein Backend dort läuft
export default new HouseholdExpensesService('http://localhost:8000');
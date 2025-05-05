package rest

import (
	"backend_go/db/dao"
	"backend_go/db/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterHaushaltsausgabenRoutes(r *gin.Engine, expenseDAO *dao.HaushaltsausgabenDAO) {
	expenseRoutes := r.Group("/haushaltsausgaben")
	{
		expenseRoutes.POST("/", createExpense(expenseDAO))
		expenseRoutes.GET("/", getExpenses(expenseDAO))
		expenseRoutes.PUT("/:id", updateExpense(expenseDAO))
		expenseRoutes.DELETE("/:id", deleteExpense(expenseDAO))
		expenseRoutes.GET("/:userid/:month", getExpensesByUserAndMonth(expenseDAO))
	}
}

func getExpenses(expenseDAO *dao.HaushaltsausgabenDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.DefaultQuery("user_id", "")
		var userID int
		var err error
		if userIDStr != "" {
			userID, err = strconv.Atoi(userIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
				return
			}
		}

		var expenses []models.Haushaltsausgaben
		if userID > 0 {
			expenses, err = expenseDAO.GetByUserID(userID)
		} else {
			expenses, err = expenseDAO.GetAll()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
			return
		}
		c.JSON(http.StatusOK, expenses)
	}
}

func createExpense(expenseDAO *dao.HaushaltsausgabenDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Definiere eine Variable für die Eingabedaten (kann das DB-Modell sein)
		var input models.Haushaltsausgaben // Verwende dein GORM-Modell

		// 2. Binde den JSON-Body an die 'input'-Struktur
		log.Println("[Handler.createExpense] Attempting to bind JSON body...") // Logging hinzugefügt
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("[Handler.createExpense] ERROR binding JSON: %v", err) // Logging hinzugefügt
			// Gib einen detaillierteren Fehler zurück, falls möglich
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
			return
		}

		// 3. JSON erfolgreich gebunden - Logge die empfangenen Daten (optional, auf sensible Daten achten)
		log.Printf("[Handler.createExpense] Successfully bound JSON: UserID=%d, Type=%s, ValueTotal=%.2f, Description=%s, Faelligkeitstag=%s, CreditStart=%v, CreditEnd=%v",
			input.UserID, input.Type, input.ValueTotal, input.Description, input.Faelligkeitstag, input.CreditStart, input.CreditEnd)

		// 4. Validierung der gebundenen Daten (Beispiele)
		if input.ValueTotal <= 0 {
			log.Println("[Handler.createExpense] Validation failed: ValueTotal is not positive.")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid valuetotal, must be positive"})
			return
		}
		// Achte auf den Feldnamen: Das Frontend sendet 'user_id' oder 'userid'?
		// Wenn 'userid' gesendet wird, füge `json:"userid"` zum UserID Feld im Model hinzu.
		if input.UserID <= 0 {
			log.Println("[Handler.createExpense] Validation failed: UserID is missing or invalid.")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing userid"})
			return
		}
		if input.Type == "" {
			log.Println("[Handler.createExpense] Validation failed: Type is missing.")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing expense type"})
			return
		}
		// Weitere Validierungen nach Bedarf (z.B. für Datumsformate, wenn sie nicht time.Time wären)

		// 5. Rufe die DAO Create-Methode mit den Daten aus der 'input'-Struktur auf
		//    Die Felder ValueRate und Zahldatum werden aus dem input übernommen (sind ggf. Nullwerte)
		log.Println("[Handler.createExpense] Calling DAO.Create...")
		expense, err := expenseDAO.Create(
			input.Description,
			input.ValueTotal,
			input.ValueRate,   // Kommt aus dem gebundenen JSON (ggf. 0)
			input.CreditStart, // Kommt aus dem gebundenen JSON (ggf. time.Time{})
			input.CreditEnd,   // Kommt aus dem gebundenen JSON (ggf. time.Time{})
			input.Type,
			input.UserID,
			input.Faelligkeitstag,
			input.Zahldatum, // Kommt aus dem gebundenen JSON (ggf. time.Time{})
		)
		if err != nil {
			// Fehler wurde bereits im DAO geloggt, hier nur Antwort senden
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"}) // Ggf. spezifischere Fehlermeldung
			return
		}

		// 6. Gebe das erstellte Objekt zurück (Erfolg wurde im DAO geloggt)
		c.JSON(http.StatusCreated, expense)
	}
}

// --- updateExpense, deleteExpense etc. ---
// Stelle sicher, dass updateExpense auch JSON bindet (was es laut deinem Code bereits tut)
// und dass die Logik zur Handhabung der Update-Parameter korrekt ist (siehe vorherige Antwort)
func updateExpense(expenseDAO *dao.HaushaltsausgabenDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Haushaltsausgaben
		log.Println("[Handler.updateExpense] Attempting to bind JSON body...")
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("[Handler.updateExpense] ERROR binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
			return
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("[Handler.updateExpense] ERROR parsing ID '%s': %v", idStr, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
			return
		}
		log.Printf("[Handler.updateExpense] Preparing update for ID: %d with data: %+v", id, input)

		// HIER: Wichtige Überlegung aus vorheriger Antwort anwenden:
		// Entweder DAO.Update so anpassen, dass es nur die Felder aus 'input' nimmt,
		// ODER hier das bestehende Objekt laden und nur die relevanten Felder aus 'input'
		// vor dem Aufruf von DAO.Update setzen.
		// Der aktuelle DAO.Update erwartet ALLE Felder.
		// Einfachere Variante für jetzt (sendet ggf. Nullwerte für nicht im JSON vorhandene Felder):
		err = expenseDAO.Update(
			id,
			input.Description,
			input.ValueTotal,
			input.ValueRate,
			input.CreditStart,
			input.CreditEnd,
			input.Type,   // Typ sollte i.d.R. nicht geändert werden
			input.UserID, // UserID sollte i.d.R. nicht geändert werden
			input.Faelligkeitstag,
			input.Receipt, // Receipt kommt nicht aus JSON
		)

		if err != nil {
			// Fehler wurde im DAO geloggt
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
	}
}

func deleteExpense(expenseDAO *dao.HaushaltsausgabenDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
			return
		}

		err = expenseDAO.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
	}
}

func getExpensesByUserAndMonth(expenseDAO *dao.HaushaltsausgabenDAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Param("userid")
		month := c.Param("month")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		expenses, err := expenseDAO.GetByUserIDAndMonth(userID, month)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
			return
		}

		c.JSON(http.StatusOK, expenses)
	}
}

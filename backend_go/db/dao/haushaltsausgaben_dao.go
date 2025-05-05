package dao

import (
	"backend_go/db/models" // Stelle sicher, dass dieser Importpfad korrekt ist
	"log"                  // Importiere das log-Paket
	"time"

	"gorm.io/gorm"
)

// HaushaltsausgabenDAO Struktur
type HaushaltsausgabenDAO struct {
	db *gorm.DB
}

// NewHaushaltsausgabenDAO Konstruktor für das DAO
func NewHaushaltsausgabenDAO(db *gorm.DB) *HaushaltsausgabenDAO {
	// Optional: Logge die erfolgreiche DAO-Initialisierung
	log.Println("HaushaltsausgabenDAO initialized successfully.")
	return &HaushaltsausgabenDAO{db: db}
}

// --- Konstanten (bleiben unverändert, aber ungenutzt von GORM Methoden) ---
const (
	CreateHaushaltsausgabenQuery = `
    INSERT INTO haushaltsausgaben
    (description, valuetotal, valuerate, creditstart, creditend, type, user_id, created_at, changed_at, faelligkeitstag, zahldatum)
    VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?)
    RETURNING id`

	UpdateHaushaltsausgabenQuery = `
    UPDATE haushaltsausgaben
    SET description = ?, valuetotal = ?, valuerate = ?, creditstart = ?, creditend = ?, type = ?, user_id = ?, changed_at = CURRENT_TIMESTAMP, faelligkeitstag = ?, zahldatum = ?
    WHERE id = ?`

	ReadHaushaltsausgabenQuery = `SELECT id, description, valuetotal, valuerate, creditstart, creditend, type, user_id, created_at, changed_at, faelligkeitstag FROM haushaltsausgaben`

	DeleteHaushaltsausgabenQuery = `DELETE FROM haushaltsausgaben WHERE id = ?`

	ExpensesQuery = `
    SELECT *
    FROM haushaltsausgaben
    WHERE userid = $1
    AND (
        type = 'monthlycosts'
        OR
        (type = 'credit'
            AND TO_CHAR(creditend::date, 'YYYY-MM') >= $2
            AND TO_CHAR(creditstart::date, 'YYYY-MM') <= $3
        )
        OR
        (type = 'allelse'
            AND TO_CHAR(created_at::date, 'YYYY-MM') = $4
        )
        OR
        (type = 'invoice'
            AND TO_CHAR(zahldatum::date, 'YYYY-MM') = $4
        )
    )
`
)

// --- CRUD Methoden mit Logging ---

// Create erstellt einen neuen Haushaltsausgaben-Eintrag
func (dao *HaushaltsausgabenDAO) Create(description string, valuetotal, valuerate float64, creditstart, creditend time.Time, typ string, userID int, faelligkeitstag string, zahldatum time.Time) (*models.Haushaltsausgaben, error) {
	// --- Logging: Eingangsparameter ---
	log.Printf("[DAO.Create] Received parameters: UserID=%d, Type=%s, ValueTotal=%.2f, Description=%s, Faelligkeitstag=%s, CreditStart=%v, CreditEnd=%v, Zahldatum=%v, ValueRate=%.2f",
		userID, typ, valuetotal, description, faelligkeitstag, creditstart, creditend, zahldatum, valuerate)

	// Erstelle das GORM-Modellobjekt
	expense := models.Haushaltsausgaben{
		Description:     description,
		ValueTotal:      valuetotal,
		ValueRate:       valuerate,
		CreditStart:     creditstart,
		CreditEnd:       creditend,
		Type:            typ,
		UserID:          userID,
		CreatedAt:       time.Now(), // Explizit gesetzt, auch wenn autoCreateTime aktiv sein könnte
		ChangedAt:       time.Now(), // Explizit gesetzt, auch wenn autoUpdateTime aktiv sein könnte
		Faelligkeitstag: faelligkeitstag,
		Zahldatum:       zahldatum,
		// Receipt wird hier nicht übergeben, bleibt also leer/nil
	}

	// --- Logging: Objekt vor dem Speichern ---
	log.Printf("[DAO.Create] Attempting to create expense object: %+v", expense)

	// Create the entry in the database using GORM
	// --- Logging: Vor dem DB-Aufruf ---
	log.Println("[DAO.Create] Calling db.Create()...")
	if err := dao.db.Create(&expense).Error; err != nil {
		// --- Logging: Fehler beim Speichern ---
		log.Printf("[DAO.Create] ERROR creating expense in DB: %v", err)
		return nil, err // Fehler zurückgeben
	}

	// --- Logging: Erfolg ---
	// Die 'expense'-Variable enthält jetzt die vom DB generierte ID (wenn erfolgreich)
	log.Printf("[DAO.Create] Successfully created expense with ID: %d", expense.ID)

	// Return the created Haushaltsausgaben
	return &expense, nil
}

// GetAll holt alle Haushaltsausgaben
func (dao *HaushaltsausgabenDAO) GetAll() ([]models.Haushaltsausgaben, error) {
	log.Println("[DAO.GetAll] Fetching all expenses...")
	var expenses []models.Haushaltsausgaben
	if err := dao.db.Find(&expenses).Error; err != nil {
		log.Printf("[DAO.GetAll] ERROR fetching expenses: %v", err)
		return nil, err
	}
	log.Printf("[DAO.GetAll] Fetched %d expenses.", len(expenses))
	return expenses, nil
}

// GetByUserID holt Ausgaben für eine bestimmte UserID
func (dao *HaushaltsausgabenDAO) GetByUserID(userID int) ([]models.Haushaltsausgaben, error) {
	log.Printf("[DAO.GetByUserID] Fetching expenses for UserID: %d...", userID)
	var expenses []models.Haushaltsausgaben
	// Annahme: GORM mappt das Feld 'UserID' im Struct auf die Spalte 'userid' in der DB
	// dank des `gorm:"column:userid"` Tags im Model. Sicherer ist es, explizit zu sein:
	if err := dao.db.Where("userid = ?", userID).Find(&expenses).Error; err != nil {
		log.Printf("[DAO.GetByUserID] ERROR fetching expenses for UserID %d: %v", userID, err)
		return nil, err
	}
	log.Printf("[DAO.GetByUserID] Fetched %d expenses for UserID %d.", len(expenses), userID)
	return expenses, nil
}

// GetByID holt eine einzelne Ausgabe anhand ihrer ID
func (dao *HaushaltsausgabenDAO) GetByID(id int) (*models.Haushaltsausgaben, error) {
	log.Printf("[DAO.GetByID] Fetching expense by ID: %d...", id)
	var expense models.Haushaltsausgaben
	// First() findet den ersten passenden Eintrag oder gibt gorm.ErrRecordNotFound zurück
	if err := dao.db.First(&expense, id).Error; err != nil {
		log.Printf("[DAO.GetByID] ERROR fetching expense ID %d: %v", id, err)
		return nil, err
	}
	log.Printf("[DAO.GetByID] Found expense ID %d.", id)
	return &expense, nil
}

// Update modifiziert einen bestehenden Eintrag
func (dao *HaushaltsausgabenDAO) Update(id int, description string, valuetotal, valuerate float64, creditstart, creditend time.Time, typ string, userID int, faelligkeitstag string, receipt []byte) error {
	// --- Logging: Eingangsparameter für Update ---
	// Hinweis: userID sollte normalerweise nicht über ein Update geändert werden. Typ auch selten.
	log.Printf("[DAO.Update] Attempting to update expense ID: %d (Data provided: UserID=%d, Type=%s, ValueTotal=%.2f, ...)", id, userID, typ, valuetotal)

	// Verwende eine Map für Updates, um GORM explizit zu sagen, welche Spalten aktualisiert werden sollen.
	// Die Schlüssel der Map sollten den Spaltennamen in der Datenbank entsprechen.
	updates := map[string]interface{}{
		"description": description,
		"valuetotal":  valuetotal,
		"valuerate":   valuerate,
		"creditstart": creditstart,
		"creditend":   creditend,
		"type":        typ,
		// "userid":          userID, // UserID sollte normalerweise nicht geändert werden!
		"changed_at":      time.Now(), // Immer aktualisieren
		"faelligkeitstag": faelligkeitstag,
		// "receipt":         receipt, // Receipt-Handling ggf. separat/anders
	}
	// Optional: Entferne Zero Values aus der Map, falls diese Felder nicht explizit
	// auf ihren Nullwert gesetzt werden sollen, wenn sie im Input fehlen.
	// (Standardmäßig aktualisiert GORM mit Updates() auch auf Nullwerte, wenn sie in der Map sind)

	// --- Logging: Update-Daten ---
	log.Printf("[DAO.Update] Applying updates for ID %d: %+v", id, updates)

	// Wende Updates mit GORM an
	// --- Logging: Vor DB-Aufruf ---
	log.Println("[DAO.Update] Calling db.Model().Where().Updates()...")
	// Wichtig: Model(&models.Haushaltsausgaben{}) gibt GORM den Tabellenkontext
	tx := dao.db.Model(&models.Haushaltsausgaben{}).Where("id = ?", id).Updates(updates)
	if err := tx.Error; err != nil {
		// --- Logging: Fehler beim Update ---
		log.Printf("[DAO.Update] ERROR updating expense ID %d: %v", id, err)
		return err
	}

	// --- Logging: Erfolg (Anzahl betroffener Zeilen) ---
	log.Printf("[DAO.Update] Successfully updated expense ID %d. Rows affected: %d", id, tx.RowsAffected)

	// Prüfe, ob überhaupt eine Zeile betroffen war (optional, aber gut zu wissen)
	if tx.RowsAffected == 0 {
		log.Printf("[DAO.Update] WARN: Update for expense ID %d affected 0 rows. Does the ID exist?", id)
		// Optional: Fehler zurückgeben, wenn ID nicht gefunden wurde (oder einfach Warnung belassen)
		// return gorm.ErrRecordNotFound
	}

	return nil
}

// Delete entfernt einen Eintrag
func (dao *HaushaltsausgabenDAO) Delete(id int) error {
	log.Printf("[DAO.Delete] Attempting to delete expense ID: %d...", id)
	// --- Logging: Vor DB-Aufruf ---
	log.Println("[DAO.Delete] Calling db.Where().Delete()...")
	tx := dao.db.Where("id = ?", id).Delete(&models.Haushaltsausgaben{})
	if err := tx.Error; err != nil {
		// --- Logging: Fehler beim Löschen ---
		log.Printf("[DAO.Delete] ERROR deleting expense ID %d: %v", id, err)
		return err
	}

	// --- Logging: Erfolg (Anzahl betroffener Zeilen) ---
	log.Printf("[DAO.Delete] Successfully deleted expense ID %d. Rows affected: %d", id, tx.RowsAffected)
	if tx.RowsAffected == 0 {
		log.Printf("[DAO.Delete] WARN: Delete for expense ID %d affected 0 rows. Does the ID exist?", id)
		// Optional: Fehler zurückgeben, wenn ID nicht gefunden wurde
		// return gorm.ErrRecordNotFound
	}
	return nil
}

// GetByUserIDAndMonth holt Ausgaben für User und Monat via Raw SQL
func (dao *HaushaltsausgabenDAO) GetByUserIDAndMonth(userID int, month string) ([]models.Haushaltsausgaben, error) {
	log.Printf("[DAO.GetByUserIDAndMonth] Fetching expenses for UserID %d and Month %s...", userID, month)
	var expenses []models.Haushaltsausgaben
	// --- Logging: Vor DB-Aufruf ---
	// Bereinige das Query-Logging, um nicht das ganze Query zu loggen, falls es sensibel ist
	log.Printf("[DAO.GetByUserIDAndMonth] Executing Raw Query with params: UserID=%d, Month=%s", userID, month)
	if err := dao.db.Raw(ExpensesQuery, userID, month, month, month).Scan(&expenses).Error; err != nil {
		log.Printf("[DAO.GetByUserIDAndMonth] ERROR fetching expenses for UserID %d, Month %s: %v", userID, month, err)
		return nil, err
	}
	log.Printf("[DAO.GetByUserIDAndMonth] Fetched %d expenses for UserID %d, Month %s.", len(expenses), userID, month)
	return expenses, nil
}

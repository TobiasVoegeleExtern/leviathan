// database/mongodb.go
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref" // Benötigt für die Ping-Prüfung
)

// ClientInstanz hält eine Instanz des verbundenen MongoDB Clients.
// In vielen Anwendungen möchten Sie nur eine einzige Client-Instanz haben,
// die über die Lebensdauer der Anwendung hinweg geteilt wird.
// Dies ist ein einfaches Beispiel für die Connect-Funktion; ein Singleton-Muster
// wäre für eine globale Instanz besser geeignet, aber für die "Connect"-Logik selbst
// ist diese Funktion korrekt.
var ClientInstanz *mongo.Client

// ConnectDB stellt eine Verbindung zu MongoDB unter Verwendung der bereitgestellten URI her.
// Es gibt eine Instanz von *mongo.Client und einen Fehler zurück, falls die Verbindung fehlschlägt.
// Ein Kontext mit Timeout wird verwendet, um unbegrenztes Warten zu vermeiden.
func ConnectDB(uri string) (*mongo.Client, error) {
	// Definieren Sie einen Kontext mit einem Timeout für den Verbindungsversuch.
	// Die Verwendung eines Timeouts verhindert, dass die Anwendung unbegrenzt blockiert,
	// falls die Datenbank unerreichbar ist.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Stellen Sie sicher, dass der Kontext bei Funktionsende immer abgebrochen wird,
	// um Ressourcen freizugeben.
	defer cancel()

	// Erstellen Sie einen neuen Client mit der gegebenen URI und Standardoptionen.
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// Loggen Sie den Verbindungsfehler und geben Sie ihn zurück.
		log.Printf("Fehler bei der Verbindung zu MongoDB: %v", err)
		return nil, fmt.Errorf("fehler bei der Verbindung zu MongoDB: %w", err)
	}

	// Optional: Prüfen Sie die Verbindung durch Senden eines Ping-Befehls.
	// Dies verifiziert, dass die Verbindung tatsächlich aufgebaut und die Authentifizierung
	// (falls in der URI angegeben) erfolgreich war.
	// Verwenden Sie einen neuen, kurzen Kontext nur für den Ping-Vorgang.
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel() // Kontext für Ping abbrechen

	// Ping an den primären Server senden.
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		// Loggen Sie den Ping-Fehler.
		log.Printf("Fehler beim Ping an MongoDB nach erfolgreicher Verbindung: %v", err)
		// Wichtig: Wenn Ping fehlschlägt, aber Connect erfolgreich war,
		// versuchen Sie, den Client sauber zu trennen.
		disconnectCtx, disconnectCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer disconnectCancel() // Kontext für Trennung abbrechen
		if disconnectErr := client.Disconnect(disconnectCtx); disconnectErr != nil {
			log.Printf("Fehler beim Trennen des MongoDB-Clients nach Ping-Fehler: %v", disconnectErr)
		}
		// Geben Sie den ursprünglichen Ping-Fehler zurück.
		return nil, fmt.Errorf("fehler beim Ping an MongoDB: %w", err)
	}

	// Wenn alles erfolgreich war, loggen Sie den Erfolg.
	log.Println("Verbindung zu MongoDB erfolgreich hergestellt und Ping empfangen.")

	// Setzen Sie die globale Instanz (falls Sie dieses Muster verwenden möchten)
	// ClientInstanz = client

	// Geben Sie den verbundenen Client zurück.
	return client, nil
}

// GetCollection holt eine Referenz auf eine spezifische Collection.
// Dies ist eine Hilfsfunktion, die oft zusammen mit ConnectDB verwendet wird.
// Es wird erwartet, dass der Client bereits verbunden ist.
func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	if client == nil {
		log.Fatal("MongoDB Client ist nicht initialisiert. Rufen Sie zuerst ConnectDB auf.")
	}
	database := client.Database(dbName)
	collection := database.Collection(collectionName)
	return collection
}

// CloseDBClient trennt die Verbindung zum MongoDB Server.
// Dies sollte normalerweise beim Herunterfahren der Anwendung aufgerufen werden.
func CloseDBClient(client *mongo.Client) {
	if client == nil {
		log.Println("MongoDB Client ist bereits null oder nicht verbunden.")
		return
	}
	// Kontext mit Timeout für die Trennung
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Anwendung wird heruntergefahren, trenne MongoDB-Verbindung...")
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Fehler während der MongoDB-Client-Trennung: %v", err)
	}
	fmt.Println("MongoDB-Client getrennt.")
}

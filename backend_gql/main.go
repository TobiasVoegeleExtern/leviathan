package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend_gql/database"
	"backend_gql/router"
)

func main() {
	log.Println("GraphQL Backend Service startet...")

	// --- MongoDB Verbindung herstellen ---
	mongoURI := os.Getenv("MONGODB_URI")
	mongoDBName := os.Getenv("MONGODB_DATABASE")
	if mongoURI == "" || mongoDBName == "" {
		log.Fatal("MONGODB_URI oder MONGODB_DATABASE Umgebungsvariable nicht gesetzt")
	}
	mongoClient, err := database.ConnectDB(mongoURI)
	if err != nil {
		log.Fatalf("Fehler beim Verbinden mit MongoDB: %v", err)
	}
	mongoDatabase := mongoClient.Database(mongoDBName)
	// Wichtig: Client am Ende schließen!
	defer database.CloseDBClient(mongoClient) // Einfaches Defer hier für den Hauptclient

	// --- Router initialisieren ---
	// Übergib die DB-Instanz an den spezifischen GraphQL-Router Setup
	engine := router.SetupGraphQLRouter(mongoDatabase)
	log.Println("GraphQL Router erfolgreich eingerichtet.")

	// --- HTTP Server starten ---
	server := &http.Server{
		Addr:    ":8080", // Lauscht auf Port 8080 im Container
		Handler: engine,
	}

	// --- Graceful Shutdown Handling ---
	go func() {
		log.Printf("GraphQL Server lauscht auf %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("GraphQL Serverfehler: %v", err)
		}
	}()

	// Warten auf Interrupt-Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("GraphQL Server wird heruntergefahren...")

	// Kontext für Shutdown erstellen
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// HTTP Server herunterfahren
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("GraphQL Server Shutdown fehlgeschlagen: %v", err)
	}

	// MongoDB Client wird durch defer am Ende von main geschlossen

	log.Println("GraphQL Server erfolgreich heruntergefahren.")
}

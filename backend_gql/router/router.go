package router

import (
	"log"
	"time"

	"backend_gql/graphql" // Passe Pfad an, falls Modulname anders

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupGraphQLRouter konfiguriert den Gin-Router NUR für GraphQL Endpunkte.
func SetupGraphQLRouter(mongoDatabase *mongo.Database) *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false

	// --- CORS Konfiguration ---
	// Passe AllowOrigins an deine Frontend URLs an
	// Wichtig: Wenn dieses Backend von denselben Frontends genutzt wird, gleiche Regeln wie bei backend_go verwenden!
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // OPTIONS ist wichtig für GraphQL POST Preflights
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- GraphQL Routen ---
	// Baue das GraphQL Schema
	schema, err := graphql.BuildSchema(mongoDatabase)
	if err != nil {
		log.Fatalf("Fehler beim Erstellen des GraphQL Schemas: %v", err)
	}
	log.Println("GraphQL Schema erfolgreich erstellt.")

	// Registriere die GraphQL Endpunkte (/graphql)
	// Die Funktion RegisterGraphQLRoutes ist in router/graphql_routes.go definiert
	RegisterGraphQLRoutes(r, schema)

	log.Println("GraphQL Router Setup abgeschlossen.")
	return r
}

// graphql/schema.go
package graphql

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Für MongoDB ObjectID
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/graphql-go/graphql"
)

// --- Go Structs passend zu Ihren MongoDB Dokumenten ---

// User struct - Angepasst basierend auf dem 'tobi' Dokument und vorherigem Schema
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` // MongoDB _id

	// Mappe das 'username'-Feld aus MongoDB auf das 'Name'-Feld in Go (für GraphQL Konsistenz)
	Name string `bson:"username"` // <-- Angepasst

	// Mappe das 'password'-Feld aus MongoDB. WIRD NICHT PER GRAPHQL EXPORTIERT!
	Password string `bson:"password"` // <-- Hinzugefügt

	// Neue Felder aus Ihrem Dokument
	Birthday int    `bson:"birthday"` // <-- Hinzugefügt (Annahme: Jahr als Integer)
	Weight   string `bson:"weight"`   // <-- Hinzugefügt (Annahme: Gewicht als String)
	Height   string `bson:"height"`   // <-- Hinzugefügt (Annahme: Größe als String)

	// Vorhandene Felder aus dem vorherigen Schema (nehmen wir an, sie existieren in DB oder sind für GraphQL geplant)
	Age      int    `bson:"age,omitempty"`       // bson:"age" in DB, omitempty falls Feld nicht in allen Docs
	Size     int    `bson:"size,omitempty"`      // bson:"size" in DB
	EyeColor string `bson:"eye_color,omitempty"` // bson:"eye_color" in DB

}

// Nutrition struct (typisch für Zutaten)
type Nutrition struct {
	Fat     float64 `bson:"fat"`
	Protein float64 `bson:"protein"`
	Carbs   float64 `bson:"carbs"`
}

// Ingredient struct (typisch für Rezepte)
type Ingredient struct {
	Name      string    `bson:"name"`
	Nutrition Nutrition `bson:"nutrition"` // Eingebettete Nutrition Struktur
	Category  string    `bson:"category"`  // Z.B. "Protein", "Fett", "Kohlenhydrate"
}

// Recipe struct - Angepasst um Bildfeld hinzuzufügen
type Recipe struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"` // "_id" in MongoDB
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Ingredients  []Ingredient       `bson:"ingredients"` // Array von Ingredient Strukturen
	Instructions string             `bson:"instructions"`
	Image        []byte             `bson:"image,omitempty"` // <-- Feld für Bilddaten (BLOB)
	// Fügen Sie hier weitere Felder hinzu, die Sie in Ihrer Recipe Collection haben
}

// --- GraphQL Objekt Typ Definitionen ---

// NutritionType GraphQL Objekt
var NutritionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Nutrition", // Name des Typs im GraphQL Schema
	Fields: graphql.Fields{
		"fat":     &graphql.Field{Type: graphql.Float}, // Feldname 'fat', Typ Float
		"protein": &graphql.Field{Type: graphql.Float},
		"carbs":   &graphql.Field{Type: graphql.Float},
	},
})

// IngredientType GraphQL Objekt
var IngredientType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ingredient",
	Fields: graphql.Fields{
		"name":      &graphql.Field{Type: graphql.String},
		"nutrition": &graphql.Field{Type: NutritionType}, // Referenz auf den NutritionType
		"category":  &graphql.Field{Type: graphql.String},
	},
})

// UserType GraphQL Objekt - Angepasst um neue Felder hinzuzufügen (ohne Password)
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User", // Name des Typs im GraphQL Schema
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.ID},     // Mappt von Go User.ID (_id)
		"name":     &graphql.Field{Type: graphql.String}, // Mappt von Go User.Name (entspricht DB 'username')
		"age":      &graphql.Field{Type: graphql.Int},    // Mappt von Go User.Age
		"size":     &graphql.Field{Type: graphql.Int},    // Mappt von Go User.Size
		"eyeColor": &graphql.Field{Type: graphql.String}, // Mappt von Go User.EyeColor (entspricht DB 'eye_color')

		// Neue Felder aus der angepassten Go-Struktur, die exportiert werden sollen
		"birthday": &graphql.Field{Type: graphql.Int},    // <-- Hinzugefügt
		"weight":   &graphql.Field{Type: graphql.String}, // <-- Hinzugefügt
		"height":   &graphql.Field{Type: graphql.String}, // <-- Hinzugefügt

		// Das 'password'-Feld wird NICHT hier hinzugefügt/exportiert.
	},
})

// RecipeType GraphQL Objekt - Angepasst um Bildfeld hinzuzufügen (war im letzten Code schon da)
var RecipeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recipe",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.ID},
		"name":         &graphql.Field{Type: graphql.String},
		"description":  &graphql.Field{Type: graphql.String},
		"ingredients":  &graphql.Field{Type: graphql.NewList(IngredientType)},
		"instructions": &graphql.Field{Type: graphql.String},
		"image":        &graphql.Field{Type: graphql.String}, // <-- Feld für Bilddaten als String (Base64)
		// Fügen Sie hier weitere Felder hinzu
	},
})

// --- Root Query Definition ---
// Definiert den Einstiegspunkt für Leseoperationen im GraphQL Schema.

// BuildSchema erstellt das GraphQL Schema und verbindet Resolver mit MongoDB Collections.
// Es benötigt eine verbundene MongoDB Datenbank Instanz.
func BuildSchema(mongoDatabase *mongo.Database) (graphql.Schema, error) {
	// Holen Sie Referenzen auf die spezifischen Collections, die von den Resolvern benötigt werden.
	// Die Namen der Collections werden aus Umgebungsvariablen gelesen, wie in docker-compose.yml definiert.
	usersCollectionName := os.Getenv("MONGODB_COLLECTION_USERS")
	if usersCollectionName == "" {
		log.Fatal("Umgebungsvariable MONGODB_COLLECTION_USERS nicht gesetzt")
	}
	usersCollection := mongoDatabase.Collection(usersCollectionName)

	recipesCollectionName := os.Getenv("MONGODB_COLLECTION_RECIPES")
	if recipesCollectionName == "" {
		log.Fatal("Umgebungsvariable MONGODB_COLLECTION_RECIPES nicht gesetzt")
	}
	recipesCollection := mongoDatabase.Collection(recipesCollectionName)

	// Definition des Root Query Typs
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery", // Konventioneller Name
		Fields: graphql.Fields{
			// === Benutzer Abfrage ===
			"users": &graphql.Field{
				Type:        graphql.NewList(UserType),                  // Dieser Query gibt eine Liste von UserType zurück
				Description: "Get a list of users, optionally filtered", // Beschreibung im Schema
				Args: graphql.FieldConfigArgument{ // Argumente, die der Query akzeptiert
					"id":       &graphql.ArgumentConfig{Type: graphql.ID},     // Optionaler Filter nach ID
					"name":     &graphql.ArgumentConfig{Type: graphql.String}, // Optionaler Filter nach Name
					"age":      &graphql.ArgumentConfig{Type: graphql.Int},    // Optionaler Filter nach Alter (exakt)
					"minAge":   &graphql.ArgumentConfig{Type: graphql.Int},    // Optionaler Filter nach Mindestalter
					"maxAge":   &graphql.ArgumentConfig{Type: graphql.Int},    // Optionaler Filter nach Maximalalter
					"eyeColor": &graphql.ArgumentConfig{Type: graphql.String}, // Optionaler Filter nach Augenfarbe
					// Fügen Sie hier weitere Filter-Argumente hinzu, falls benötigt
				},
				// Der Resolver: Diese Funktion wird ausgeführt, wenn der 'users' Query aufgerufen wird.
				// 'p' enthält die Argumente, den Kontext etc.
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Greifen Sie auf die MongoDB Collection Referenz zu.
					collection := usersCollection

					// Bauen Sie das MongoDB Filter-Dokument basierend auf den GraphQL-Argumenten.
					filter := bson.D{}

					if id, ok := p.Args["id"].(string); ok && id != "" {
						objectID, err := primitive.ObjectIDFromHex(id)
						if err != nil {
							log.Printf("Ungültiges ObjectID-Format für ID: %v", err)
							return []User{}, nil // Bei ungültiger ID im Query leere Liste zurückgeben
						}
						filter = append(filter, bson.E{Key: "_id", Value: objectID})
					}

					if name, ok := p.Args["name"].(string); ok && name != "" {
						// Filtere auf das 'username' Feld in MongoDB
						filter = append(filter, bson.E{Key: "username", Value: primitive.Regex{Pattern: name, Options: "i"}})
					}

					if age, ok := p.Args["age"].(int); ok {
						filter = append(filter, bson.E{Key: "age", Value: age})
					} else {
						ageFilterDoc := bson.D{}
						if minAge, ok := p.Args["minAge"].(int); ok {
							ageFilterDoc = append(ageFilterDoc, bson.E{Key: "$gte", Value: minAge})
						}
						if maxAge, ok := p.Args["maxAge"].(int); ok {
							ageFilterDoc = append(ageFilterDoc, bson.E{Key: "$lte", Value: maxAge})
						}
						if len(ageFilterDoc) > 0 {
							filter = append(filter, bson.E{Key: "age", Value: ageFilterDoc})
						}
					}

					if eyeColor, ok := p.Args["eyeColor"].(string); ok && eyeColor != "" {
						filter = append(filter, bson.E{Key: "eye_color", Value: primitive.Regex{Pattern: eyeColor, Options: "i"}})
					}

					// --- Abfrage auf MongoDB ausführen ---
					queryCtx, queryCancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer queryCancel()

					cursor, err := collection.Find(queryCtx, filter)
					if err != nil {
						log.Printf("Fehler beim Finden von Benutzern: %v", err)
						return nil, err
					}
					defer cursor.Close(queryCtx)

					// Dekodieren Sie die Ergebnisse in eine Slice von User Structs
					var users []User
					if err = cursor.All(queryCtx, &users); err != nil {
						log.Printf("Fehler beim Dekodieren von Benutzern: %v", err)
						return nil, err
					}

					// Geben Sie die Liste der dekodierten Benutzer zurück.
					return users, nil
				},
			},

			// === Rezepte Abfrage ===
			"recipes": &graphql.Field{
				Type:        graphql.NewList(RecipeType), // Gibt eine Liste von RecipeType zurück
				Description: "Get a list of recipes, optionally filtered",
				Args: graphql.FieldConfigArgument{ // Argumente für die Rezept-Abfrage
					"id":   &graphql.ArgumentConfig{Type: graphql.ID},
					"name": &graphql.ArgumentConfig{Type: graphql.String}, // Filter nach Rezeptname
					// Fügen Sie hier weitere Filter-Argumente hinzu
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Greifen Sie auf die MongoDB Collection Referenz für Rezepte zu.
					collection := recipesCollection

					// Bauen Sie das MongoDB Filter-Dokument.
					filter := bson.D{}

					if id, ok := p.Args["id"].(string); ok && id != "" {
						objectID, err := primitive.ObjectIDFromHex(id)
						if err != nil {
							log.Printf("Ungültiges ObjectID-Format für ID: %v", err)
							return []Recipe{}, nil // Bei ungültiger ID leere Liste zurückgeben
						}
						filter = append(filter, bson.E{Key: "_id", Value: objectID})
					}

					if name, ok := p.Args["name"].(string); ok && name != "" {
						filter = append(filter, bson.E{Key: "name", Value: primitive.Regex{Pattern: name, Options: "i"}})
					}

					// --- Abfrage auf MongoDB ausführen ---
					queryCtx, queryCancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer queryCancel()

					cursor, err := collection.Find(queryCtx, filter)
					if err != nil {
						log.Printf("Fehler beim Finden von Rezepten: %v", err)
						return nil, err
					}
					defer cursor.Close(queryCtx)

					// Dekodieren Sie die Ergebnisse in eine Slice von Recipe Structs
					var recipes []Recipe
					if err = cursor.All(queryCtx, &recipes); err != nil {
						log.Printf("Fehler beim Dekodieren von Rezepten: %v", err)
						return nil, err
					}

					return recipes, nil
				},
			},
			// Fügen Sie hier weitere Query-Felder hinzu (z.B. Mutationen in einem separaten Typ)
		},
	})

	// Erstellen Sie das finale GraphQL Schema.
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation, // Fügen Sie hier Mutationen hinzu, falls vorhanden
		// Subscription: rootSubscription, // Fügen Sie hier Subscriptions hinzu, falls vorhanden
	})
	if err != nil {
		log.Fatalf("Fehler beim Erstellen des GraphQL Schemas: %v", err)
		return graphql.Schema{}, fmt.Errorf("fehler beim Erstellen des GraphQL Schemas: %w", err)
	}

	return schema, nil
}

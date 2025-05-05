package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gql "github.com/graphql-go/graphql"
)

// RegisterGraphQLRoutes sets up the GraphQL endpoint and GraphiQL IDE route on the provided Gin router.
// It requires a compiled GraphQL schema.
func RegisterGraphQLRoutes(r *gin.Engine, schema gql.Schema) {
	// --- GraphQL HTTP Handler ---
	graphqlHandler := func(c *gin.Context) {
		// Set Content-Type header to JSON
		c.Header("Content-Type", "application/json")

		// Only process POST requests
		if c.Request.Method != "POST" {
			c.Status(http.StatusMethodNotAllowed)
			return
		}

		// Decode the request body containing the GraphQL query
		var params struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültiger Anfrage-Body", "details": err.Error()})
			return
		}

		// Execute GraphQL query using graphql.Do
		result := gql.Do(gql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			VariableValues: params.Variables,
			OperationName:  params.OperationName,
			Context:        c.Request.Context(),
		})

		// Log any execution errors
		if len(result.Errors) > 0 {
			log.Printf("GraphQL Ausführungsfehler: %v", result.Errors)
		}

		// Return the result as JSON
		c.JSON(http.StatusOK, result)
	}

	// Register routes
	r.POST("/graphql", graphqlHandler)

	log.Println("GraphQL endpoint registriert unter /graphql (POST)")
	log.Println("GraphiQL IDE verfügbar unter /graphiql (GET)")
}

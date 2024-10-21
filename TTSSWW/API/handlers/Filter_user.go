// backend/API/handlers/filter_users.go
package handlers

import (
    "encoding/json"
    "net/http"
    "TTSSWW/API/models"
    "strings"

    "github.com/gin-gonic/gin"
)

func FilterUsers(c *gin.Context) {
    // Ejemplo de usuarios
    users := []models.User{
        {ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30, Location: "New York", Interests: []string{"music", "sports"}, Preferences: map[string]string{"food": "vegan"}},
        {ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25, Location: "Los Angeles", Interests: []string{"movies", "sports"}, Preferences: map[string]string{"food": "vegetarian"}},
        // Agrega más usuarios según sea necesario
    }

    // Obtener parámetros de filtrado de la solicitud
    location := c.Query("location")
    interests := c.QueryArray("interests")
    preferences := c.Request.URL.Query()

    // Filtrar usuarios
    var filteredUsers []models.User
    for _, user := range users {
        if location != "" && user.Location != location {
            continue
        }
        if len(interests) > 0 && !containsAll(user.Interests, interests) {
            continue
        }
        if !matchesPreferences(user.Preferences, preferences) {
            continue
        }
        filteredUsers = append(filteredUsers, user)
    }

    // Convertir a JSON y enviar la respuesta
    c.JSON(http.StatusOK, filteredUsers)
}

func containsAll(userInterests, filterInterests []string) bool {
    for _, interest := range filterInterests {
        if !contains(userInterests, interest) {
            return false
        }
    }
    return true
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func matchesPreferences(userPreferences map[string]string, filterPreferences map[string][]string) bool {
    for key, values := range filterPreferences {
        if key == "location" || key == "interests" {
            continue
        }
        if userValue, ok := userPreferences[key]; ok {
            if !contains(values, userValue) {
                return false
            }
        } else {
            return false
        }
    }
    return true
}
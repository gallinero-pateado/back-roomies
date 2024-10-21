// backend/API/handlers/filter_users.go
package handlers

import (
    "net/http"
    "TTSSWW/API/models"
    "strings"

    "github.com/gin-gonic/gin"
)

func FilterUsers(c *gin.Context) {
    // Ejemplo de usuarios
    users := []models.User{
        {ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30, Comuna: "Comuna1", Carrera: "Ingeniería", Intereses: []string{"música", "deportes"}, Preferencias: map[string]string{"comida": "vegana"}},
        {ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25, Comuna: "Comuna2", Carrera: "Medicina", Intereses: []string{"películas", "deportes"}, Preferencias: map[string]string{"comida": "vegetariana"}},
        // Agrega más usuarios según sea necesario
    }

    // Obtener parámetros de filtrado de la solicitud
    comuna := c.Query("comuna")
    carrera := c.Query("carrera")
    intereses := c.QueryArray("intereses")
    preferencias := c.Request.URL.Query()

    // Filtrar usuarios
    var filteredUsers []models.User
    for _, user := range users {
        if comuna != "" && user.Comuna != comuna {
            continue
        }
        if carrera != "" && user.Carrera != carrera {
            continue
        }
        if len(intereses) > 0 && !containsAll(user.Intereses, intereses) {
            continue
        }
        if !matchesPreferences(user.Preferencias, preferencias) {
            continue
        }
        filteredUsers = append(filteredUsers, user)
    }

    // Convertir a JSON y enviar la respuesta
    c.JSON(http.StatusOK, filteredUsers)
}

func containsAll(userIntereses, filterIntereses []string) bool {
    for _, interest := range filterIntereses {
        if !contains(userIntereses, interest) {
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

func matchesPreferences(userPreferencias map[string]string, filterPreferencias map[string][]string) bool {
    for key, values := range filterPreferencias {
        if key == "comuna" || key == "carrera" || key == "intereses" {
            continue
        }
        if userValue, ok := userPreferencias[key]; ok {
            if !contains(values, userValue) {
                return false
            }
        } else {
            return false
        }
    }
    return true
}
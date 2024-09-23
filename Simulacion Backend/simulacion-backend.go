//Correr un programa: go run nombre.extension
//Compilar programa: go  build nombre.extension

//Expresiones Explicita e implicitas

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Definición de estructuras
type User struct {
	ID       int
	Username string
	Password string
	Name     string
}

type Message struct {
	SenderID   int
	ReceiverID int
	Content    string
}

// Variables globales
var users []User
var messages []Message
var favorites = make(map[int]map[int]struct{})

// Usuario actualmente logueado
var currentUser *User

// Inicialización de usuarios de ejemplo
func init() {
	users = []User{
		{ID: 1, Username: "CLopez", Password: "clave1", Name: "Carlos Lopez"},
		{ID: 2, Username: "MTam", Password: "clave2", Name: "Monica Tampar"},
		{ID: 3, Username: "ARodrgiuez", Password: "clave3", Name: "Adolfo Rodriguez"},
		{ID: 4, Username: "PMardo", Password: "clave4", Name: "Pablo Mardo"},
		{ID: 5, Username: "SLlach", Password: "clave5", Name: "Samuel Llach"},
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Bucle de inicio de sesión
	for {
		fmt.Println("Bienvenido. Por favor, inicie sesión.")
		fmt.Print("Usuario: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Contraseña: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		user := authenticate(username, password)
		if user != nil {
			currentUser = user
			fmt.Printf("Inicio de sesión exitoso. Bienvenido, %s!\n", currentUser.Username)
			break
		} else {
			fmt.Println("Credenciales incorrectas. Inténtelo de nuevo.\n")
		}
	}

	// Bucle del menú principal
	for {
		fmt.Println("\nMenú:")
		fmt.Println("1) Ver perfiles y marcar como favorito")
		fmt.Println("2) Enviar mensaje a un perfil")
		fmt.Println("3) Ver mensajes enviados")
		fmt.Println("4) Ver perfiles marcados como favoritos")
		fmt.Println("5) Salir")
		fmt.Print("Seleccione una opción: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Opción inválida. Inténtelo de nuevo.")
			continue
		}

		switch choice {
		case 1:
			viewProfilesAndMarkFavorite(reader)
		case 2:
			sendMessage(reader)
		case 3:
			viewSentMessages()
		case 4:
			viewFavorites()
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida. Inténtelo de nuevo.")
		}
	}
}

// Función para autenticar al usuario
func authenticate(username, password string) *User {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user
		}
	}
	return nil
}

// Función para ver perfiles y marcar como favorito
func viewProfilesAndMarkFavorite(reader *bufio.Reader) {
	fmt.Println("\nPerfiles disponibles:")
	for _, user := range users {
		if user.ID != currentUser.ID {
			fmt.Printf("ID: %d, %s\n", user.ID, user.Name)
		}
	}

	fmt.Print("Ingrese el ID del perfil que desea marcar como favorito (o presione Enter para volver): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	userID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	if userExists(userID) && userID != currentUser.ID {
		err := addFavorite(currentUser.ID, userID)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Perfil agregado a favoritos.")
		}
	} else {
		fmt.Println("Usuario no encontrado.")
	}
}

// Función para enviar un mensaje a un perfil
func sendMessage(reader *bufio.Reader) {
	fmt.Println("\nPerfiles disponibles para enviar mensaje:")
	for _, user := range users {
		if user.ID != currentUser.ID {
			fmt.Printf("ID: %d, %s\n", user.ID, user.Name)
		}
	}

	fmt.Print("Ingrese el ID del perfil al que desea enviar un mensaje (o presione Enter para volver): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	userID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}

	if userExists(userID) && userID != currentUser.ID {
		fmt.Print("Escriba el mensaje: ")
		content, _ := reader.ReadString('\n')
		content = strings.TrimSpace(content)
		sendMessageToUser(currentUser.ID, userID, content)
		fmt.Println("Mensaje enviado.")
	} else {
		fmt.Println("Usuario no encontrado.")
	}
}

// Función para ver mensajes enviados (modificada)
func viewSentMessages() {
	fmt.Println("\nMensajes enviados:")
	found := false
	for _, msg := range messages {
		if msg.SenderID == currentUser.ID {
			sender := getUserByID(msg.SenderID)
			receiver := getUserByID(msg.ReceiverID)
			fmt.Printf("De %s a %s: %s\n", sender.Name, receiver.Name, msg.Content)
			found = true
		}
	}
	if !found {
		fmt.Println("No has enviado mensajes.")
	}
}

// Función para ver perfiles marcados como favoritos
func viewFavorites() {
	fmt.Println("\nPerfiles marcados como favoritos:")
	favs, exists := favorites[currentUser.ID]
	if !exists || len(favs) == 0 {
		fmt.Println("No tienes perfiles favoritos.")
		return
	}
	for favID := range favs {
		user := getUserByID(favID)
		fmt.Printf("%s\n", user.Name)
	}
}

// Función para agregar un perfil a favoritos
func addFavorite(userID, favoriteUserID int) error {
	if _, exists := favorites[userID]; !exists {
		favorites[userID] = make(map[int]struct{})
	}

	if _, exists := favorites[userID][favoriteUserID]; exists {
		return fmt.Errorf("El usuario ya está en favoritos")
	}

	favorites[userID][favoriteUserID] = struct{}{}
	return nil
}

// Función para enviar un mensaje a un usuario
func sendMessageToUser(senderID, receiverID int, content string) {
	message := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}
	messages = append(messages, message)
}

// Función para verificar si un usuario existe
func userExists(userID int) bool {
	for _, user := range users {
		if user.ID == userID {
			return true
		}
	}
	return false
}

// Función para obtener un usuario por ID
func getUserByID(userID int) *User {
	for _, user := range users {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selecciona una opción:")
	fmt.Println("1. Generar Basic Auth")
	fmt.Println("2. Decifrar Basic Auth")
	fmt.Println("3. Decifrar Base64 normal")
	fmt.Print("Opción (1/2/3): ")

	opcionStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la opción:", err)
		os.Exit(1)
	}
	opcionStr = strings.TrimSpace(opcionStr)

	var opcion int
	if opcionStr == "1" {
		opcion = 1
	} else if opcionStr == "2" {
		opcion = 2
	} else if opcionStr == "3" {
		opcion = 3
	} else {
		fmt.Println("Opción inválida. Por favor ingresa 1, 2 o 3.")
		os.Exit(1)
	}

	switch opcion {
	case 1:
		// Generar Basic Auth
		fmt.Print("Ingresa el nombre de usuario: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer el usuario:", err)
			os.Exit(1)
		}
		username = strings.TrimSpace(username)

		fmt.Print("Ingresa la contraseña: ")
		password, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la contraseña:", err)
			os.Exit(1)
		}
		password = strings.TrimSpace(password)

		credentials := username + ":" + password
		encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
		authHeader := "Basic " + encodedCredentials

		fmt.Println("\n--- Resultado ---")
		fmt.Printf("Credenciales originales: %s\n", credentials)
		fmt.Printf("Cabecera de Autorización HTTP: %s\n", authHeader)
		fmt.Println("-----------------")
	case 2:
		// Decifrar Basic Auth
		fmt.Print("Ingresa la cabecera Basic Auth: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la cabecera:", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		// Eliminar "Basic " si está presente
		input = strings.TrimPrefix(input, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error al decodificar:", err)
			os.Exit(1)
		}
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			fmt.Println("Formato inválido de credenciales decodificadas.")
			os.Exit(1)
		}
		fmt.Println("\n--- Resultado ---")
		fmt.Printf("Usuario: %s\n", parts[0])
		fmt.Printf("Contraseña: %s\n", parts[1])
		fmt.Println("-----------------")
	case 3:
		// Decifrar Base64 normal
		fmt.Print("Ingresa la cadena Base64 a decodificar: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la cadena:", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)

		decoded, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error al decodificar Base64:", err)
			os.Exit(1)
		}

		fmt.Println("\n--- Resultado ---")
		fmt.Printf("Cadena original: %s\n", string(decoded))
		fmt.Println("-----------------")
	default:
		fmt.Println("Opción inválida.")
		os.Exit(1)
	}
}

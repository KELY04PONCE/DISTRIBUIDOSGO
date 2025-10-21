package main

//CLIENTE MENSAJES
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Bienvenido al sistemas de mensajes - Elija una opcion:")
		fmt.Println("1. Enviar mensaje")
		fmt.Println("2. Listar mensajes")
		fmt.Println("3. Salir")
		fmt.Print("Opcion:")

		if !scanner.Scan() {
			break
		}

		opcion := strings.TrimSpace(scanner.Text())
		switch opcion {
		case "1":
			fmt.Print("Ingrese mensaje:")
			if scanner.Scan() {
				mensaje := scanner.Text()
				urlCompleta := "http://localhost:8081/guardar?mensaje=" + url.QueryEscape(mensaje)
				resp, err := http.Get(urlCompleta)
				if err != nil {
					fmt.Println("ERROR SERVIDOR NO DISPONIBLE")
					continue
				}
				body, _ := io.ReadAll(resp.Body)
				resp.Body.Close()
				fmt.Println(string(body))

			}
		case "2":
			resp, err := http.Get("http://localhost:8081/listarmensajes")
			if err != nil {
				fmt.Println("ERROR SERVIDOR NO DISPONIBLE")
				continue
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			fmt.Println("\n" + string(body))

		case "3":
			fmt.Println("Nos vemos!...")
			return
		default:
			fmt.Println("Opcion invalida")

		}
	}
}

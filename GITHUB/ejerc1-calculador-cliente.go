CLIENTE
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//LEER MENSAJE
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//LEER OPERACIONES
	entrada := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese operaciones en el formato: OPERACION NUM1 NUM2")
	for {
		fmt.Print("> ")
		if !entrada.Scan() {
			break
		}
		operacion := entrada.Text()

		//ENVIO AL SERVIDOR
		fmt.Fprintf(conn, "%s\n", operacion)

		if scanner.Scan() {
			respuesta := scanner.Text()
			fmt.Println("Respuesta:", respuesta)
		}
	}
}
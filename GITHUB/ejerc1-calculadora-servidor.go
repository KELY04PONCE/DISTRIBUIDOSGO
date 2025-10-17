SERVIDOR
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Servidor calculadora escuchando ")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		//MANEJA CADA CLIENTE CON UNA GOROUTINE POR SEPARADO
		go manejarCliente(conn)
	}
}
func manejarCliente(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	fmt.Fprintf(conn, "Calculadora lista, El formato es: OPERACION NUM1 NUM2 \n")
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Fields(linea) //SEPARA POR ESPACIOS
		if len(partes) != 3 {
			fmt.Fprintf(conn, "ERROR, FORMATO INCORRECTO \n")
			continue
		}
		operacion := partes[0]
		num1, err1 := strconv.ParseFloat(partes[1], 64)
		num2, err2 := strconv.ParseFloat(partes[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Fprintf(conn, "ERROR, NUMEROS INCORRECTOS \n")
			continue
		}
		var resultado float64
		switch operacion {
		case "SUMA":
			resultado = num1 + num2
		case "RESTA":
			resultado = num1 - num2
		case "MULT":
			resultado = num1 * num2
		case "DIV":
			if num2 == 0 {
				fmt.Fprintf(conn, "ERROR, NO SE PUEDE DIVIDIR ENTRE CERO \n")
				continue
			}
			resultado = num1 / num2
		default:
			fmt.Fprintf(conn, "ERROR, OPERACION DESCONOCIDA \n")
			continue
		}
		fmt.Fprintf(conn, "EL RESULTADO ES: %.2f \n", resultado)
		log.Printf("Se calculo: %s %f %f = %f \n", operacion, num1, num2, resultado)
	}
}
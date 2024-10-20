package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	sustantivos = []string{"proyecto", "desarrollo", "aplicación", "código", "programa", "sistema", "interfaz", "usuario", "cliente", "servidor"}
	adjetivos   = []string{"innovador", "eficiente", "robusto", "escalable", "modular", "flexible", "intuitivo", "responsive", "dinámico", "adaptable"}
	verbos      = []string{"desarrollar", "implementar", "optimizar", "refactorizar", "diseñar", "programar", "testear", "depurar", "mantener", "actualizar"}
)

func generarPalabra(palabras []string) string {
	return palabras[rand.Intn(len(palabras))]
}

func generarFrase() string {
	return fmt.Sprintf("El %s %s permite %s el sistema de manera efectiva.",
		generarPalabra(sustantivos),
		generarPalabra(adjetivos),
		generarPalabra(verbos))
}

func generarParrafo(numFrases int) string {
	frases := make([]string, numFrases)
	for i := 0; i < numFrases; i++ {
		frases[i] = generarFrase()
	}
	return strings.Join(frases, " ")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Generador de Texto Aleatorio")
	fmt.Println("============================")

	fmt.Print("Ingrese el número de párrafos que desea generar: ")
	var numParrafos int
	fmt.Scanln(&numParrafos)

	fmt.Print("Ingrese el número de frases por párrafo: ")
	var frasesPorParrafo int
	fmt.Scanln(&frasesPorParrafo)

	for i := 0; i < numParrafos; i++ {
		fmt.Printf("\nPárrafo %d:\n%s\n", i+1, generarParrafo(frasesPorParrafo))
	}
}

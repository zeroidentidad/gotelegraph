package telegraph

import (
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// DecodeBits2Morse decodifica de bits a morse
func DecodeBits2Morse(messageBits string) (messageMorse string) {
	// Declaramos la exprecion regular que busca patrones repetidos de 0 o 1
	re := regexp.MustCompile(`(0)+|(1)+`)
	// Buscamos la exprecion en el mensaje
	matchs := re.FindAllString(messageBits, -1)
	// Decodificamos el patron de velocidad del usuario
	average, max := decodePulses(matchs)

	// Recorremos las coincidencias
	for _, match := range matchs {
		// Casteamos a float para comparar con el promedio
		count := float64(len(match))
		// Si la longitud de la coincidencia es mayor al maximo del patron de velocidad retornamos
		if count > max {
			return
		}

		switch match[0] {
		// Evaluamos los 0 como pausas
		case 48:
			// Si la pausa es menor o igual al promedio se evalua como una misma letra
			if count <= average {
				messageMorse += ""
				// Si no como otra letra
			} else {
				messageMorse += " "
			}
		// Evaluamos los 1 como pulsos
		case 49:
			// Si un pulso es menor o igual al promedio se evalua como un punto
			if count <= average {
				messageMorse += "."
				// Si no como un guion
			} else {
				messageMorse += "-"
			}
		}

	}

	return
}

// DecodeMorse2Bits convierte de morse a bits intentando emular la variacion humana
func DecodeMorse2Bits(messageMorse string) (messageBits string) {
	// Inicilizamos la semilla para el numero random
	rand.Seed(time.Now().UnixNano())
	// Min y maximo para el numero aleatorio
	var min, max int
	// Recorremos el mensaje letra por letra
	for i, message := range messageMorse {
		switch message {
		// Punto
		case 46:
			min = 1
			max = 4
			randomN := rand.Intn(max-min) + min
			messageBits += strings.Repeat("1", randomN)
		// Guion
		case 45:
			min = 4
			max = 8
			randomN := rand.Intn(max-min) + min
			messageBits += strings.Repeat("1", randomN)
		// Espacio
		case 32:
			if i > 0 && (messageMorse[i-1] == 45 || messageMorse[i-1] == 46) {
				min = 3
				max = 6
				randomN := rand.Intn(max-min) + min
				messageBits += strings.Repeat("0", randomN)
			}
		}
		// Espacios entre una misma letra
		if message != 32 {

			min = 1
			max = 3
			randomN := rand.Intn(max-min) + min
			messageBits += strings.Repeat("0", randomN)
		}

	}
	// Se añade una pausa larga para finalizar
	messageBits += strings.Repeat("0", 9)
	return
}

// Translate2Human decodifica de Morse a Humano
func Translate2Human(morseMessage string) (humanMessage string) {
	// Separamos por palabras el string en morse
	words := strings.Split(morseMessage, "  ")
	// Recorremos todas las palabras
	for _, word := range words {

		// Separamos por letras el string en morse
		letters := strings.Split(word, " ")
		// Recorremos todas las letras
		for _, letter := range letters {
			if fullStop == letter {
				return
			}
			// Hacemos una busqueda por indexacion
			if transLetter, ok := translateText[letter]; ok {
				humanMessage += transLetter
			}
		}
		humanMessage += " "
	}

	return
}

// Translate2Morse decodifica de Humano a Morse
func Translate2Morse(humanMessage string) (morseMessage string) {
	// Separamos en letras el string en texto
	letters := strings.Split(humanMessage, "")
	// Recorremos todas las letras
	for i, letter := range letters {
		// Hacemos una busqueda por indexacion
		if transMorse, ok := translateMorse[strings.ToLower(letter)]; ok {
			morseMessage += transMorse
		}
		// Separacion entre palabras
		if i != len(letters)-1 {
			morseMessage += " "
		}
	}

	return
}

func decodePulses(matchs []string) (avg, max float64) {
	// Array para tener control sobre la cantidad de pulsos enviados
	var pulses []int

	// Recorremos las cadenas de 0 y 1
	for i, match := range matchs {
		// Buscamos los 1 para los pulsos
		if match[0] == 49 {
			// Añadimos al array de pulsos
			pulses = append(pulses, i)
			// Tomamos la longitud para buscar un maximo y el promedio
			count := float64(len(match))
			if count > max {
				max = count
			}
			// Sumamos al promedio
			avg += count
		}
	}
	// Obtenemos la cantidad de pausas realizadas desde el primer pulso encontrado
	sumPauses, lenPauses := decodePauses(pulses, matchs)
	// A la cantidad de pulsos sumamos la de pausas
	avg += sumPauses
	// Obtenemos el total de pausas y pulsos enviados
	lenght := float64(len(pulses)) + lenPauses
	// Obtenemos el promedio de pulsos y pausas
	avg = math.Round(avg / lenght)

	return
}

func decodePauses(pulses []int, matchs []string) (sum, lenght float64) {
	// Recorremos todas las pausas registradas a partir del primer y ultimo pulso registrado
	for i := pulses[0]; i < pulses[len(pulses)-1]; i++ {
		if matchs[i][0] == 48 {
			sum += float64(len(matchs[i]))
			lenght++
		}
	}

	return
}

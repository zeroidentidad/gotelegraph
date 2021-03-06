package telegraph

// Equivalencia texto morse
var translateMorse = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
}

// Equivalencia morse texto
var translateText = map[string]string{
	".-":    "a",
	"-...":  "b",
	"-.-.":  "c",
	"-..":   "d",
	".":     "e",
	"..-.":  "f",
	"--.":   "g",
	"....":  "h",
	"..":    "i",
	".---":  "j",
	"-.-":   "k",
	".-..":  "l",
	"--":    "m",
	"-.":    "n",
	"---":   "o",
	".--.":  "p",
	"--.-":  "q",
	".-.":   "r",
	"...":   "s",
	"-":     "t",
	"..-":   "u",
	"...-":  "v",
	".--":   "w",
	"-..-":  "x",
	"-.--":  "y",
	"--..":  "z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

// Constante full stop
const fullStop = ".-.-.-"

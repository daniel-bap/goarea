package goarea

import "math"

// Pi é uma proporção numérica definida pela relação
// entre o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsavel por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Essa não é uma função visivel porque foi usado underline no inicio
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}

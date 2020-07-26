package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

//En una interface se definen los metodos que se van a utilizar
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	CalsificacionVegetal() string
}

/*Genero Humano*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}
type mujer struct {
	hombre // heredamos las mismas propiedades de hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	}

	return "Mujer"

}

func (h *hombre) estaVivo() bool {
	return h.vivo
}

func (m *mujer) respirar() {
	m.respirando = true
}
func (m *mujer) comer() {
	m.comiendo = true
}
func (m *mujer) pensar() {
	m.pensando = true
}
func (m *mujer) sexo() string {
	return "Mujer"
}
func (m *mujer) estaVivo() bool {
	return m.vivo
}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando", hu.sexo())
}

/*Genero animal*/
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) esCarnivoro() bool {
	return p.carnivoro
}
func (p *perro) estaVivo() bool {
	return true
}

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")

}

func AnimalesCarnivoros(an animal) int {
	if an.esCarnivoro() == true {
		return 1
	}
	return 0

}

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoreos %d", totalCarnivoros)
	fmt.Printf("Estoy vivo %t", estoyVivo(Dogo))

}

package estruturas

const Pi = 3.14159265

type Circle struct {
	Raio float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Forma interface {
	Area() float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
	return (retangulo.Largura * retangulo.Altura)
}

func AreaCircle(circle Circle) float64 {
	return (Pi * (circle.Raio * circle.Raio))
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circle) Area() float64 {
	return Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

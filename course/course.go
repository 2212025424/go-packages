package course

import "fmt"

// Course struct: Es una abstracción de un curso
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	classes map[uint]string
}

func (c *course) SetClases(clases map[uint]string) {
	c.classes = clases
}

// New: funcion que se emplea para instanciar un curso
func New(Name string, Price float64, IsFre bool) *course {
	if Price == 0 {
		Price = 30
	}
	return &course{
		Name:   Name,
		Price:  Price,
		IsFree: IsFre,
	}
}

// PrintClasses: Es un método el cual imprime las clases
func (c *course) PrintClasses() {
	text := "------------------------------------------ \n"
	text += "-- Las clases son: \n"
	text += "------------------------------------------ \n"
	for _, class := range c.classes {
		text += "-> " + class + "\n"
	}
	text += "------------------------------------------ \n"
	fmt.Println(text)
}

// ChangePrice: Es un método el cual cambia el precio de un curs
func (c *course) ChangePrice(price float64) {
	c.Price = price
}

package weightconv

import "fmt"


type Lb float64
type Kilogram float64

func (lb Lb) String() string {
	return fmt.Sprintf("%glb", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func LbToKg(lb Lb) Kilogram  {
	return Kilogram(0.453592 * lb)
}

func KgToLb(kg Kilogram) Lb  {
	return Lb(2.20462 * kg)
}


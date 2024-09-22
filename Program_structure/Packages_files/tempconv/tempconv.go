package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	CeroKelvin Kelvin = -273.15
	FeetConst Feet = 3.281
	PoutConst Pounds = 2.205
)

// Methods
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C",c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F",f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m",m)
}
func (fe Feet) String() string {
	return fmt.Sprintf("%g f",fe)
}
func (p Pounds) String() string {
	return fmt.Sprintf("%g pounds",p)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g k",k)
}

func (kilo Kilograms) String() string {
	return fmt.Sprintf("%g kilograms",kilo)
}

// Functions
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func FToC(f Fahrenheit) Celsius  {
	return Celsius((f - 32 ) * 5/9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + CeroKelvin)
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k + CeroKelvin) * 9/5 + 32)
}

func FeToMe( f Feet) Meters {
	return Meters(f / FeetConst)
}

func MeToFe(m Meters) Feet {
	return Feet(m * Meters(FeetConst))
}

func PToKil(p Pounds) Kilograms {
	return Kilograms(p / PoutConst)
}

func KilToP(k Kilograms) Pounds {
	return Pounds( k * Kilograms(PoutConst))
}
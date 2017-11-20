package weightconv

import "fmt"

type Pond float64
type Kilo float64

const (
	pondForkilo = 2.205
	kiloForPond = 0.454
)

func KiloToPond(k Kilo) Pond {
	return Pond(k) * pondForkilo
}
func PondToKilo(p Pond) Kilo {
	return Kilo(p) * kiloForPond

}
func (p Pond) String() string    { return fmt.Sprintf("%gol", p) }
 func (k Kilo) String() string { return fmt.Sprintf("%gkg", k) }

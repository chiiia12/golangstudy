package weightconv

import "testing"

func TestKiloToPond(t *testing.T) {
	var arg Kilo = 10
	if KiloToPond(arg) != 22.05{
		t.Error(`KiloToPond(10) is not '22.05' actual is`, KiloToPond(arg))
	}
}

func TestPondToKilo(t *testing.T) {
	var arg Pond = 10
	if PondToKilo(arg) != 4.54 {
		t.Error(`PondToKilo(10) is not '4.54' actual is `, PondToKilo(arg))
	}

}

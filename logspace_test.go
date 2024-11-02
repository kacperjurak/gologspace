package gologspace_test

import (
	"github.com/kacperjurak/gologspace"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("should generate proper logspace", func(t *testing.T) {
		got := gologspace.Generate(1, 10, 10)
		want := []float64{1, 1.2589254117941673, 1.5848931924611136, 1.9952623149688797, 2.51188643150958, 3.162277660168379, 3.9810717055349727, 5.011872336272722, 6.309573444801931, 7.943282347242812}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

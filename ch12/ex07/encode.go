package sexpr

import (
	"reflect"
	"fmt"
	"io"
)

type Encoder struct {
	w          io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}
func (enc *Encoder) Encode(v interface{}) error {
	if err := encode(enc.w, reflect.ValueOf(v)); err != nil {
		return err
	}
	return nil
}

func encode(w io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(w, "nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(w, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%f", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(w, "#C(%f,%f)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		output := "nil"
		if v.Bool() {
			output = "t"
		}
		fmt.Fprintf(w, "%v", output)
	case reflect.Interface:
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "(")
		if v.Type().Name() == "" {
			fmt.Fprintf(w, "%q ", v.Elem().Type().String())
		} else {
			fmt.Fprintf(w, "%s.%s", v.Type().PkgPath(), v.Type().Name())
		}
		if err := encode(w, v.Elem()); err != nil {
			return err
		}
		fmt.Fprintf(w, ")")
	case reflect.String:
		fmt.Fprintf(w, "%q", v.String())
	case reflect.Ptr:
		return encode(w, v.Elem())
	case reflect.Array, reflect.Slice:
		fmt.Fprintf(w, "(")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(w, " ")
			}
			if err := encode(w, v.Index(i)); err != nil {
				return err
			}
			if i != v.Len()-1 {
				fmt.Fprintf(w, "\n")
				fmt.Fprintf(w, "\n")
				fmt.Fprintf(w, "\t")
				fmt.Fprintf(w, "\t")
				fmt.Fprintf(w, " ")
			}
		}
		fmt.Fprintf(w, ")")
	case reflect.Struct:
		fmt.Fprintf(w, "(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(w, " ")
			}
			if i != 0 {
				fmt.Fprintf(w, "\n (%s ", v.Type().Field(i).Name)
			} else {
				fmt.Fprintf(w, "(%s ", v.Type().Field(i).Name)
			}
			if err := encode(w, v.Field(i)); err != nil {
				return err
			}
			fmt.Fprintf(w, ")")
		}
		fmt.Fprintf(w, ")")
	case reflect.Map:
		fmt.Fprintf(w, "(")
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprintf(w, " ")
			}
			if i != 0 {
				fmt.Fprintf(w, "\n")
				fmt.Fprintf(w, "\t")
				fmt.Fprintf(w, "\t")
				fmt.Fprintf(w, " ")
			}

			fmt.Fprintf(w, "(")
			if err := encode(w, key); err != nil {
				return err
			}
			fmt.Fprintf(w, " ")
			if err := encode(w, v.MapIndex(key)); err != nil {
				return err
			}
			fmt.Fprintf(w, ")")
		}

		fmt.Fprintf(w, ")")

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

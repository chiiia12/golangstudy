package sexpr

import (
	"bytes"
	"reflect"
	"fmt"
)

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%f,%f)", real(v.Complex()), imag(v.Complex()))
	case reflect.Bool:
		fmt.Fprintf(buf, "%v", v.Bool())
	case reflect.Interface:
		buf.WriteByte('{')
		if v.Type().Name() == "" {
			fmt.Fprintf(buf, "%q", v.Elem().Type().String())
		} else {
			fmt.Fprintf(buf, "%s.%s", v.Type().PkgPath(), v.Type().Name())
		}
		if err := encode(buf, v.Elem()); err != nil {
			return err
		}
		buf.WriteByte('}')
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			fmt.Fprintf(buf, "\"%s\":", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')
	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {

			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte(',')
			}
		}

		buf.WriteByte('}')

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

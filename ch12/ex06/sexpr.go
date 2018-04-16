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
		if v.Int() != 0 {
			fmt.Fprintf(buf, "%d", v.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() != 0 {
			fmt.Fprintf(buf, "%d", v.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() != 0 {
			fmt.Fprintf(buf, "%f", v.Float())
		}
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
		if v.String() != "" {
			fmt.Fprintf(buf, "%q", v.String())
		}
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		if v.Len() == 0 {
			break
		}
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
			if isZeroValue(v.Field(i)) {
				continue
			}
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
		if v.Len() == 0 {
			break
		}
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

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Int() == 0
	case reflect.Interface:
		return v.IsNil()
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr:
		return v.IsNil()
	case reflect.Map:
		return v.Len() == 0
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	}
	return false
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

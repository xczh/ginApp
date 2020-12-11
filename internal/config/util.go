package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// fillVar 利用字符串strVal表示的值填充变量refVal
func fillVar(refVal reflect.Value, strVal string) error {
	if !refVal.CanSet() {
		return fmt.Errorf("reflect value cannot set")
	}
	switch refVal.Kind() {
	case reflect.String:
		refVal.SetString(strVal)
	case reflect.Int:
		iv, err := strconv.ParseInt(strVal, 10, 32)
		if err != nil {
			return err
		}
		refVal.SetInt(iv)
	case reflect.Int64:
		if refVal.Type().String() == "time.Duration" {
			t, err := time.ParseDuration(strVal)
			if err != nil {
				return err
			}
			refVal.Set(reflect.ValueOf(t))
		} else {
			iv, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				return err
			}
			refVal.SetInt(iv)
		}
	case reflect.Uint:
		uiv, err := strconv.ParseUint(strVal, 10, 32)
		if err != nil {
			return err
		}
		refVal.SetUint(uiv)
	case reflect.Uint64:
		uiv, err := strconv.ParseUint(strVal, 10, 64)
		if err != nil {
			return err
		}
		refVal.SetUint(uiv)
	case reflect.Float32:
		f32, err := strconv.ParseFloat(strVal, 32)
		if err != nil {
			return err
		}
		refVal.SetFloat(f32)
	case reflect.Float64:
		f64, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return err
		}
		refVal.SetFloat(f64)
	case reflect.Bool:
		if strVal == "" {
			refVal.SetBool(false)
			break
		}
		b, err := strconv.ParseBool(strVal)
		if err != nil {
			return err
		}
		refVal.SetBool(b)
	case reflect.Slice:
		if strVal == "" {
			return nil
		}
		const sep = ","
		vals := strings.Split(strVal, sep)
		switch refVal.Type() {
		case reflect.TypeOf([]string{}):
			refVal.Set(reflect.ValueOf(vals))
		case reflect.TypeOf([]int{}):
			t := make([]int, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					return err
				}
				t[i] = int(val)
			}
		case reflect.TypeOf([]int64{}):
			t := make([]int64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				t[i] = val
			}
		case reflect.TypeOf([]uint{}):
			t := make([]uint, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return err
				}
				t[i] = uint(val)
			}
		case reflect.TypeOf([]uint64{}):
			t := make([]uint64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					return err
				}
				t[i] = val
			}
		case reflect.TypeOf([]float32{}):
			t := make([]float32, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				t[i] = float32(val)
			}
		case reflect.TypeOf([]float64{}):
			t := make([]float64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				t[i] = val
			}
		}
	}
	return nil
}

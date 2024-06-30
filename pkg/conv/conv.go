package conv

import (
	"fmt"
	"reflect"
	"strings"

	"k8s.io/client-go/util/jsonpath"
)

func ToPtr[T any](v T) *T {
	a := v
	return &a
}

func Raw[T any](v T) T {
	return v
}

func StrPtr(v interface{}) interface{} {
	return ToPtr(v.(string))
}

func PtrStr(v interface{}) interface{} {
	return *v.(*string)
}

func GetVal(data interface{}, template string) (interface{}, error) {
	jp := jsonpath.New("conv").AllowMissingKeys(true)
	jp.Parse(fmt.Sprintf("{%s}", template))
	values, err := jp.FindResults(data)
	if err != nil {
		return nil, err
	}
	if len(values) > 0 && len(values[0]) > 0 {
		return values[0][0].Interface(), nil
	}
	return nil, ErrNotFound
}

func SetVal(data interface{}, template string, val interface{}) error {
	jp := jsonpath.New("conv").AllowMissingKeys(true)
	jp.Parse(fmt.Sprintf("{%s}", template))
	values, err := jp.FindResults(data)
	if err != nil {
		return err
	}
	if len(values) > 0 && len(values[0]) > 0 {
		values[0][0].Set(reflect.ValueOf(val))
	}
	return nil
}

type Conv func(val any) any

func CopyField(srcdata interface{}, src string, dstdata interface{}, dst string, conv ...Conv) error {
	jp := jsonpath.New("conv").AllowMissingKeys(true)
	jp.Parse(fmt.Sprintf("{%s}", src))
	values, err := jp.FindResults(srcdata)
	if err != nil {
		return err
	}
	if len(values) > 0 && len(values[0]) > 0 {
		val := values[0][0].Interface()
		for _, c := range conv {
			if c != nil {
				val = c(val)
			}
		}
		jp.Parse(fmt.Sprintf("{%s}", dst))
		dstval, err := jp.FindResults(dstdata)
		if err != nil {
			return err
		}
		if len(dstval) > 0 && len(dstval[0]) > 0 {
			v := dstval[0][0]
			if v.CanSet() {
				v.Set(reflect.ValueOf(val))
				return nil
			}
			return ErrProtectedValue
		}
	}
	return ErrNotFound
}

// 同じ構造の構造体をコピー
func CopyStruct(srcdata interface{}, dstdata interface{}) error {
	var err error
	src, isSrcStructure := toReflectValue(srcdata)
	dst, isDstStructure := toReflectValue(dstdata)
	if isSrcStructure && isDstStructure {
		stv := src.Type()
		dtv := dst.Type()
		for i := 0; i < stv.NumField(); i++ {
			sfd := stv.Field(i)
			dfd, exist := dtv.FieldByName(sfd.Name)
			if !exist {
				continue
			}
			if sfd.Type.Kind() != dfd.Type.Kind() {
				return ErrDifferentStruct
			}
			svalue := src.FieldByName(sfd.Name)
			dvalue := dst.FieldByName(dfd.Name)
			switch sfd.Type.Kind() {
			case reflect.Ptr:
				if svalue.IsNil() {
					continue
				}
				var createElement func(val reflect.Value)
				createElement = func(val reflect.Value) {
					if val.Kind() == reflect.Ptr && val.IsNil() {
						n := reflect.New(val.Type().Elem())
						createElement(n.Elem())
						val.Set(n)
					}
				}
				createElement(dvalue)
				if dvalue.CanSet() {
					srckind := svalue.Elem().Kind()
					switch srckind {
					case reflect.Struct:
						err = CopyStruct(svalue, dvalue)
						if err != nil {
							return err
						}
					case reflect.Ptr:
						dvalue.Elem().Set(reflect.Indirect(svalue))
					default:
						dvalue.Elem().Set(reflect.Indirect(svalue))
					}
				} else {
					return ErrProtectedValue
				}
			case reflect.Struct:
				err = CopyStruct(svalue, dvalue)
				if err != nil {
					return err
				}
			default:
				if dvalue.CanSet() {
					dvalue.Set(svalue)
				}
			}
		}
		return nil
	}
	return ErrNotStruct
}

// reflect.Value を返す関数
func toReflectValue(i interface{}) (reflect.Value, bool) {
	rv, ok := i.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(i)
	}

	switch rv.Kind() {
	case reflect.Ptr:
		// ポインタ以外になるまで reflect.Indirect する
		return toReflectValue(reflect.Indirect(rv))
	case reflect.Struct:
		return rv, true
	}
	// 構造体以外
	return rv, false
}

func output(depth int, name string, value interface{}, kind reflect.Kind) {
	indent := strings.Repeat("  ", depth)
	if kind == reflect.Struct {
		fmt.Printf("%s- FieldName: %s, Type: %v\n", indent, name, kind)
	} else {
		fmt.Printf("%s- FieldName: %s, Value: %v, Type: %v\n", indent, name, value, kind)
	}
}

func recursive(depth int, rv reflect.Value) {
	typ := rv.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		kind := field.Type.Kind()
		value := rv.FieldByName(field.Name)

		if kind == reflect.Ptr {
			vv, isStructure := toReflectValue(value)
			output(depth, field.Name, vv.Interface(), vv.Kind())
			if isStructure {
				recursive(depth+1, vv)
			}
		} else {
			output(depth, field.Name, value, kind)
			if kind == reflect.Struct {
				recursive(depth+1, value)
			}
		}
	}
}

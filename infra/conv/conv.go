package conv

import "github.com/yamagame/school-api-gateway/entity"

func ToStrPtr(s interface{}) *string {
	r := s.(string)
	return &r
}

func ToStr(s interface{}) string {
	return s.(string)
}

func SetIfNotNil[T any](key string, field entity.FieldInterface, valptr *T) error {
	if valptr != nil {
		err := field.SetValue(key, *valptr)
		if err != nil {
			return err
		}
	}
	return nil
}

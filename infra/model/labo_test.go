package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLabo(t *testing.T) {
	rt := reflect.TypeOf(Labo{})
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		// フィールド名
		fmt.Println("---", f.Name)
		// フィールドの型
		fmt.Println(f.Type)
		// タグ
		fmt.Println(f.Tag)
		// gorm タグ
		tagKeyValue := reflect.StructTag(f.Tag)
		fmt.Println(tagKeyValue.Get("gorm"))
		if val, ok := f.Tag.Lookup("gorm"); ok {
			fmt.Println(val)
		}
	}
}

package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yamagame/school-api-gateway/pkg/snapshot"
)

func TestField(t *testing.T) {
	t.Run("set_and_get", func(t *testing.T) {
		rec := NewRecord()
		rec.SetValue("value1", 100)
		rec.SetValue("value2", "value2_name")
		rec.SetValue("value3", 1.234)

		value1, err := rec.GetValue("value1")
		assert.NoError(t, err)
		assert.Equal(t, 100, value1)
		value2, err := rec.GetValue("value2")
		assert.NoError(t, err)
		assert.Equal(t, "value2_name", value2)
		value3, err := rec.GetValue("value3")
		assert.NoError(t, err)
		assert.Equal(t, 1.234, value3)

		snapshot.Equal(t, rec.ValueMap(), "test1.json")
		// snapshot.Save(t, rec.ValueMap(), "test1.json")
	})
	t.Run("struct_conv", func(t *testing.T) {
		type Field struct {
			FieldValue1 string
			FieldValue2 string
		}
		a := &struct {
			Value1 string
			Value2 string
			Field1 Field
			Field2 *Field
		}{
			Value1: "struct_value1",
			Value2: "struct_value2",
			Field1: Field{
				FieldValue1: "field_value1",
				FieldValue2: "field_value2",
			},
			Field2: &Field{
				FieldValue1: "field_value3",
				FieldValue2: "field_value4",
			},
		}
		b := &struct {
			Value3 string
			Value4 string
			Value5 string
			Value6 *string
			Field3 *Field
		}{
			Value3: "struct_value3",
			Value4: "struct_value4",
			Field3: &Field{
				FieldValue1: "field_value5",
				FieldValue2: "field_value6",
			},
		}

		atob := []struct {
			apath string
			bpath string
			conv  Conv
		}{
			{".Value2", ".Value4", nil},
			{".Field1.FieldValue1", ".Value5", nil},
			{".Field1.FieldValue1", ".ValueSome", nil},
			{".Field2.FieldValue2", ".Value6", func(v any) any {
				return StrPtr(v)
			}},
			{".Field1.FieldValue3", ".Value5", nil},
			{".Value2", ".Field3.FieldValue1", nil},
		}

		for _, c := range atob {
			CopyField(a, c.apath, b, c.bpath, c.conv)
		}

		CopyField(b, ".Value3", a, ".Value1")
		assert.Equal(t, "struct_value3", a.Value1)
		assert.Equal(t, "struct_value2", b.Value4)
		assert.Equal(t, "field_value1", b.Value5)
		assert.Equal(t, "field_value4", *b.Value6)
		assert.Equal(t, "struct_value2", b.Field3.FieldValue1)
	})
}

func TestStructCopy(t *testing.T) {
	t.Run("struct_copy", func(t *testing.T) {
		pval := ToPtr(int32(30))
		type Fielda struct {
			Value3 string
			Value4 *string
			Value5 bool
			Value6 *string
		}
		type Fieldb struct {
			Value3 string
			Value4 *string
			Value5 bool
			Value6 *string
		}
		a := &struct {
			Value1  int32
			Value2  int64
			PtrVal1 *int32
			Field1  Fielda
			Field2  *Fielda
		}{
			Value1:  100,
			Value2:  200,
			PtrVal1: pval,
			Field1: Fielda{
				Value3: "hello world",
				Value5: true,
				Value6: ToPtr("value6"),
			},
			Field2: &Fielda{
				Value3: "ptr field",
			},
		}
		b := &struct {
			Value1  int32
			Value2  int64
			PtrVal1 *int32
			Field1  Fieldb
			Field2  *Fieldb
		}{
			Field1: Fieldb{
				Value3: "",
				Value4: ToPtr("hello"),
			},
			// Field2: &Fieldb{},
		}

		err := CopyStruct(a, b)
		assert.NoError(t, err)

		*pval = 10

		// tv := reflect.TypeOf(*a)
		// for i := 0; i < tv.NumField(); i++ {
		// 	f := tv.Field(i)
		// 	CopyField(a, "."+f.Name, b, "."+f.Name)
		// }

		assert.Equal(t, a.Value1, b.Value1)
		assert.Equal(t, a.Value2, b.Value2)
		assert.Equal(t, int32(30), *b.PtrVal1)
		assert.Equal(t, int32(10), *a.PtrVal1)
		assert.Equal(t, a.Field1.Value3, b.Field1.Value3)
		assert.Equal(t, a.Field1.Value5, b.Field1.Value5)
		assert.Equal(t, "hello", *b.Field1.Value4)
		assert.Equal(t, "value6", *b.Field1.Value6)
		assert.Equal(t, "ptr field", b.Field2.Value3)
	})
}
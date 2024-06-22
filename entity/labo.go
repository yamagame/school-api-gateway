package entity

type Labo struct {
	Field
}

func NewLabo() *Labo {
	v := &Labo{
		NewField(),
	}
	v.AddUintProp("id", 0)
	v.AddProp("name", "")
	v.AddProp("group", "")
	v.AddProp("program", "")
	v.AddProp("building", "")
	return v
}

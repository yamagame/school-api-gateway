package model

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Areas struct {
	*iconv.SliceWrapper[Area]
}

type Courses struct {
	*iconv.SliceWrapper[Course]
}

// type LaboRooms struct {
//   *irconv.Slice[LaboRoom]
// }

type Schools struct {
	*iconv.SliceWrapper[School]
}

type Addresses struct {
	*iconv.SliceWrapper[Address]
}

// type ClassRooms struct {
//   *irconv.Slice[ClassRoom]
// }

type Groups struct {
	*iconv.SliceWrapper[Group]
}

type Roles struct {
	*iconv.SliceWrapper[Role]
}

type Students struct {
	*iconv.SliceWrapper[Student]
}

// type LaboProfessors struct {
//   *irconv.Slice[LaboProfessor]
// }

type Labos struct {
	*iconv.SliceWrapper[Labo]
}

// type SchoolPersons struct {
//   *irconv.Slice[SchoolPerson]
// }

type Posts struct {
	*iconv.SliceWrapper[Post]
}

type Professors struct {
	*iconv.SliceWrapper[Professor]
}

type Chairs struct {
	*iconv.SliceWrapper[Chair]
}

type Aliases struct {
	*iconv.SliceWrapper[Alias]
}

type Buildings struct {
	*iconv.SliceWrapper[Building]
}

type Desks struct {
	*iconv.SliceWrapper[Desk]
}

type People struct {
	*iconv.SliceWrapper[Person]
}

type Employees struct {
	*iconv.SliceWrapper[Employee]
}

type Programs struct {
	*iconv.SliceWrapper[Program]
}

// type SchoolClasses struct {
//   *irconv.Slice[SchoolClass]
// }

// type BuildingAliases struct {
//   *irconv.Slice[BuildingAlias]
// }

type Classes struct {
	*iconv.SliceWrapper[Class]
}

// type LaboStudents struct {
//   *irconv.Slice[LaboStudent]
// }

type Rooms struct {
	*iconv.SliceWrapper[Room]
}

// type SchoolBuildings struct {
//   *irconv.Slice[SchoolBuilding]
// }

type Properties struct {
	*iconv.SliceWrapper[Property]
}

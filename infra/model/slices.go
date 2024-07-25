package model

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Areas struct {
	*iconv.SliceContainer[Area]
}

type Courses struct {
	*iconv.SliceContainer[Course]
}

// type LaboRooms struct {
//   *irconv.Slice[LaboRoom]
// }

type Schools struct {
	*iconv.SliceContainer[School]
}

type Addresses struct {
	*iconv.SliceContainer[Address]
}

// type ClassRooms struct {
//   *irconv.Slice[ClassRoom]
// }

type Groups struct {
	*iconv.SliceContainer[Group]
}

type Roles struct {
	*iconv.SliceContainer[Role]
}

type Students struct {
	*iconv.SliceContainer[Student]
}

// type LaboProfessors struct {
//   *irconv.Slice[LaboProfessor]
// }

type Labos struct {
	*iconv.SliceContainer[Labo]
}

// type SchoolPersons struct {
//   *irconv.Slice[SchoolPerson]
// }

type Posts struct {
	*iconv.SliceContainer[Post]
}

type Professors struct {
	*iconv.SliceContainer[Professor]
}

type Chairs struct {
	*iconv.SliceContainer[Chair]
}

type Aliases struct {
	*iconv.SliceContainer[Alias]
}

type Buildings struct {
	*iconv.SliceContainer[Building]
}

type Desks struct {
	*iconv.SliceContainer[Desk]
}

type People struct {
	*iconv.SliceContainer[Person]
}

type Employees struct {
	*iconv.SliceContainer[Employee]
}

type Programs struct {
	*iconv.SliceContainer[Program]
}

// type SchoolClasses struct {
//   *irconv.Slice[SchoolClass]
// }

// type BuildingAliases struct {
//   *irconv.Slice[BuildingAlias]
// }

type Classes struct {
	*iconv.SliceContainer[Class]
}

// type LaboStudents struct {
//   *irconv.Slice[LaboStudent]
// }

type Rooms struct {
	*iconv.SliceContainer[Room]
}

// type SchoolBuildings struct {
//   *irconv.Slice[SchoolBuilding]
// }

type Properties struct {
	*iconv.SliceContainer[Property]
}

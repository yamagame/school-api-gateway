package model

import "github.com/yamagame/school-api-gateway/pkg/iconv"

type Areas struct {
	*iconv.Slice[Area]
}

type Courses struct {
	*iconv.Slice[Course]
}

// type LaboRooms struct {
//   *irconv.Slice[LaboRoom]
// }

type Schools struct {
	*iconv.Slice[School]
}

type Addresses struct {
	*iconv.Slice[Address]
}

// type ClassRooms struct {
//   *irconv.Slice[ClassRoom]
// }

type Groups struct {
	*iconv.Slice[Group]
}

type Roles struct {
	*iconv.Slice[Role]
}

type Students struct {
	*iconv.Slice[Student]
}

// type LaboProfessors struct {
//   *irconv.Slice[LaboProfessor]
// }

type Labos struct {
	*iconv.Slice[Labo]
}

// type SchoolPersons struct {
//   *irconv.Slice[SchoolPerson]
// }

type Posts struct {
	*iconv.Slice[Post]
}

type Professors struct {
	*iconv.Slice[Professor]
}

type Chairs struct {
	*iconv.Slice[Chair]
}

type Aliases struct {
	*iconv.Slice[Alias]
}

type Buildings struct {
	*iconv.Slice[Building]
}

type Desks struct {
	*iconv.Slice[Desk]
}

type People struct {
	*iconv.Slice[Person]
}

type Employees struct {
	*iconv.Slice[Employee]
}

type Programs struct {
	*iconv.Slice[Program]
}

// type SchoolClasses struct {
//   *irconv.Slice[SchoolClass]
// }

// type BuildingAliases struct {
//   *irconv.Slice[BuildingAlias]
// }

type Classes struct {
	*iconv.Slice[Class]
}

// type LaboStudents struct {
//   *irconv.Slice[LaboStudent]
// }

type Rooms struct {
	*iconv.Slice[Room]
}

// type SchoolBuildings struct {
//   *irconv.Slice[SchoolBuilding]
// }

type Properties struct {
	*iconv.Slice[Property]
}

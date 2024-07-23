package model

import "github.com/yamagame/school-api-gateway/pkg/irconv"

type Areas struct {
	*irconv.Slice[Area]
}

type Courses struct {
	*irconv.Slice[Course]
}

// type LaboRooms struct {
//   *irconv.Slice[LaboRoom]
// }

type Schools struct {
	*irconv.Slice[School]
}

type Addresses struct {
	*irconv.Slice[Address]
}

// type ClassRooms struct {
//   *irconv.Slice[ClassRoom]
// }

type Groups struct {
	*irconv.Slice[Group]
}

type Roles struct {
	*irconv.Slice[Role]
}

type Students struct {
	*irconv.Slice[Student]
}

// type LaboProfessors struct {
//   *irconv.Slice[LaboProfessor]
// }

type Labos struct {
	*irconv.Slice[Labo]
}

// type SchoolPersons struct {
//   *irconv.Slice[SchoolPerson]
// }

type Posts struct {
	*irconv.Slice[Post]
}

type Professors struct {
	*irconv.Slice[Professor]
}

type Chairs struct {
	*irconv.Slice[Chair]
}

type Aliases struct {
	*irconv.Slice[Alias]
}

type Buildings struct {
	*irconv.Slice[Building]
}

type Desks struct {
	*irconv.Slice[Desk]
}

type People struct {
	*irconv.Slice[Person]
}

type Employees struct {
	*irconv.Slice[Employee]
}

type Programs struct {
	*irconv.Slice[Program]
}

// type SchoolClasses struct {
//   *irconv.Slice[SchoolClass]
// }

// type BuildingAliases struct {
//   *irconv.Slice[BuildingAlias]
// }

type Classes struct {
	*irconv.Slice[Class]
}

// type LaboStudents struct {
//   *irconv.Slice[LaboStudent]
// }

type Rooms struct {
	*irconv.Slice[Room]
}

// type SchoolBuildings struct {
//   *irconv.Slice[SchoolBuilding]
// }

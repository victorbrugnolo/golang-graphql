package graph

import "github.com/victorbrugnolo/golang-graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}

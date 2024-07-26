// Code generated by ent, DO NOT EDIT.

package category

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeStudyLogs holds the string denoting the study_logs edge name in mutations.
	EdgeStudyLogs = "study_logs"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "categories"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_owned_categories"
	// StudyLogsTable is the table that holds the study_logs relation/edge.
	StudyLogsTable = "study_logs"
	// StudyLogsInverseTable is the table name for the StudyLog entity.
	// It exists in this package in order to avoid circular dependency with the "studylog" package.
	StudyLogsInverseTable = "study_logs"
	// StudyLogsColumn is the table column denoting the study_logs relation/edge.
	StudyLogsColumn = "category_study_logs"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldColor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "categories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_owned_categories",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Category queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByStudyLogsCount orders the results by study_logs count.
func ByStudyLogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudyLogsStep(), opts...)
	}
}

// ByStudyLogs orders the results by study_logs terms.
func ByStudyLogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudyLogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newStudyLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudyLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudyLogsTable, StudyLogsColumn),
	)
}

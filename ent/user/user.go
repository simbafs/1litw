// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldTgid holds the string denoting the tgid field in the database.
	FieldTgid = "tgid"
	// EdgeRecords holds the string denoting the records edge name in mutations.
	EdgeRecords = "records"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RecordsTable is the table that holds the records relation/edge.
	RecordsTable = "records"
	// RecordsInverseTable is the table name for the Record entity.
	// It exists in this package in order to avoid circular dependency with the "record" package.
	RecordsInverseTable = "records"
	// RecordsColumn is the table column denoting the records relation/edge.
	RecordsColumn = "user_records"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldTgid,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByTgid orders the results by the tgid field.
func ByTgid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTgid, opts...).ToFunc()
}

// ByRecordsCount orders the results by records count.
func ByRecordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRecordsStep(), opts...)
	}
}

// ByRecords orders the results by records terms.
func ByRecords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRecordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RecordsTable, RecordsColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"1li/ent/record"
	"1li/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Record is the model entity for the Record schema.
type Record struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecordQuery when eager-loading is set.
	Edges        RecordEdges `json:"edges"`
	user_records *int
	selectValues sql.SelectValues
}

// RecordEdges holds the relations/edges for other nodes in the graph.
type RecordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			values[i] = new(sql.NullInt64)
		case record.FieldCode, record.FieldTarget:
			values[i] = new(sql.NullString)
		case record.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case record.ForeignKeys[0]: // user_records
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case record.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				r.Code = value.String
			}
		case record.FieldTarget:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value.Valid {
				r.Target = value.String
			}
		case record.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case record.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_records", value)
			} else if value.Valid {
				r.user_records = new(int)
				*r.user_records = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Record.
// This includes values selected through modifiers, order, etc.
func (r *Record) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Record entity.
func (r *Record) QueryUser() *UserQuery {
	return NewRecordClient(r.config).QueryUser(r)
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return NewRecordClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("code=")
	builder.WriteString(r.Code)
	builder.WriteString(", ")
	builder.WriteString("target=")
	builder.WriteString(r.Target)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Records is a parsable slice of Record.
type Records []*Record

// Code generated by ent, DO NOT EDIT.

package user

import (
	"1li/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Userid applies equality check predicate on the "userid" field. It's identical to UseridEQ.
func Userid(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserid, v))
}

// SuperAdmin applies equality check predicate on the "superAdmin" field. It's identical to SuperAdminEQ.
func SuperAdmin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSuperAdmin, v))
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAdmin, v))
}

// Create applies equality check predicate on the "create" field. It's identical to CreateEQ.
func Create(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreate, v))
}

// CustomCode applies equality check predicate on the "customCode" field. It's identical to CustomCodeEQ.
func CustomCode(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCustomCode, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// UseridEQ applies the EQ predicate on the "userid" field.
func UseridEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserid, v))
}

// UseridNEQ applies the NEQ predicate on the "userid" field.
func UseridNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserid, v))
}

// UseridIn applies the In predicate on the "userid" field.
func UseridIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserid, vs...))
}

// UseridNotIn applies the NotIn predicate on the "userid" field.
func UseridNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserid, vs...))
}

// UseridGT applies the GT predicate on the "userid" field.
func UseridGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldUserid, v))
}

// UseridGTE applies the GTE predicate on the "userid" field.
func UseridGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUserid, v))
}

// UseridLT applies the LT predicate on the "userid" field.
func UseridLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldUserid, v))
}

// UseridLTE applies the LTE predicate on the "userid" field.
func UseridLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUserid, v))
}

// SuperAdminEQ applies the EQ predicate on the "superAdmin" field.
func SuperAdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSuperAdmin, v))
}

// SuperAdminNEQ applies the NEQ predicate on the "superAdmin" field.
func SuperAdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSuperAdmin, v))
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAdmin, v))
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAdmin, v))
}

// CreateEQ applies the EQ predicate on the "create" field.
func CreateEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreate, v))
}

// CreateNEQ applies the NEQ predicate on the "create" field.
func CreateNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreate, v))
}

// CustomCodeEQ applies the EQ predicate on the "customCode" field.
func CustomCodeEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCustomCode, v))
}

// CustomCodeNEQ applies the NEQ predicate on the "customCode" field.
func CustomCodeNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCustomCode, v))
}

// HasRecords applies the HasEdge predicate on the "records" edge.
func HasRecords() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RecordsTable, RecordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordsWith applies the HasEdge predicate on the "records" edge with a given conditions (other predicates).
func HasRecordsWith(preds ...predicate.Record) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRecordsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}

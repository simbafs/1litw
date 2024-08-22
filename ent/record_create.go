// Code generated by ent, DO NOT EDIT.

package ent

import (
	"1li/ent/record"
	"1li/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordCreate is the builder for creating a Record entity.
type RecordCreate struct {
	config
	mutation *RecordMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (rc *RecordCreate) SetCode(s string) *RecordCreate {
	rc.mutation.SetCode(s)
	return rc
}

// SetTarget sets the "target" field.
func (rc *RecordCreate) SetTarget(s string) *RecordCreate {
	rc.mutation.SetTarget(s)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RecordCreate) SetCreatedAt(t time.Time) *RecordCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RecordCreate) SetNillableCreatedAt(t *time.Time) *RecordCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *RecordCreate) SetUserID(id int) *RecordCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rc *RecordCreate) SetNillableUserID(id *int) *RecordCreate {
	if id != nil {
		rc = rc.SetUserID(*id)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RecordCreate) SetUser(u *User) *RecordCreate {
	return rc.SetUserID(u.ID)
}

// Mutation returns the RecordMutation object of the builder.
func (rc *RecordCreate) Mutation() *RecordMutation {
	return rc.mutation
}

// Save creates the Record in the database.
func (rc *RecordCreate) Save(ctx context.Context) (*Record, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecordCreate) SaveX(ctx context.Context) *Record {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecordCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecordCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RecordCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := record.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecordCreate) check() error {
	if _, ok := rc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Record.code"`)}
	}
	if v, ok := rc.mutation.Code(); ok {
		if err := record.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Record.code": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Target(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required field "Record.target"`)}
	}
	if v, ok := rc.mutation.Target(); ok {
		if err := record.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`ent: validator failed for field "Record.target": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Record.created_at"`)}
	}
	return nil
}

func (rc *RecordCreate) sqlSave(ctx context.Context) (*Record, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RecordCreate) createSpec() (*Record, *sqlgraph.CreateSpec) {
	var (
		_node = &Record{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(record.Table, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Code(); ok {
		_spec.SetField(record.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rc.mutation.Target(); ok {
		_spec.SetField(record.FieldTarget, field.TypeString, value)
		_node.Target = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(record.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   record.UserTable,
			Columns: []string{record.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_records = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecordCreateBulk is the builder for creating many Record entities in bulk.
type RecordCreateBulk struct {
	config
	err      error
	builders []*RecordCreate
}

// Save creates the Record entities in the database.
func (rcb *RecordCreateBulk) Save(ctx context.Context) ([]*Record, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Record, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecordCreateBulk) SaveX(ctx context.Context) []*Record {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecordCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecordCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

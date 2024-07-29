// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
)

// TimerCreate is the builder for creating a Timer entity.
type TimerCreate struct {
	config
	mutation *TimerMutation
	hooks    []Hook
}

// SetStartAt sets the "start_at" field.
func (tc *TimerCreate) SetStartAt(t time.Time) *TimerCreate {
	tc.mutation.SetStartAt(t)
	return tc
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (tc *TimerCreate) SetNillableStartAt(t *time.Time) *TimerCreate {
	if t != nil {
		tc.SetStartAt(*t)
	}
	return tc
}

// SetContent sets the "content" field.
func (tc *TimerCreate) SetContent(s string) *TimerCreate {
	tc.mutation.SetContent(s)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TimerCreate) SetUserID(u uuid.UUID) *TimerCreate {
	tc.mutation.SetUserID(u)
	return tc
}

// SetID sets the "id" field.
func (tc *TimerCreate) SetID(u uuid.UUID) *TimerCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TimerCreate) SetNillableID(u *uuid.UUID) *TimerCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TimerCreate) SetUser(u *User) *TimerCreate {
	return tc.SetUserID(u.ID)
}

// SetSubjectID sets the "subject" edge to the Subject entity by ID.
func (tc *TimerCreate) SetSubjectID(id uuid.UUID) *TimerCreate {
	tc.mutation.SetSubjectID(id)
	return tc
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tc *TimerCreate) SetSubject(s *Subject) *TimerCreate {
	return tc.SetSubjectID(s.ID)
}

// AddSharedGroupIDs adds the "shared_group" edge to the Group entity by IDs.
func (tc *TimerCreate) AddSharedGroupIDs(ids ...uuid.UUID) *TimerCreate {
	tc.mutation.AddSharedGroupIDs(ids...)
	return tc
}

// AddSharedGroup adds the "shared_group" edges to the Group entity.
func (tc *TimerCreate) AddSharedGroup(g ...*Group) *TimerCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tc.AddSharedGroupIDs(ids...)
}

// Mutation returns the TimerMutation object of the builder.
func (tc *TimerCreate) Mutation() *TimerMutation {
	return tc.mutation
}

// Save creates the Timer in the database.
func (tc *TimerCreate) Save(ctx context.Context) (*Timer, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimerCreate) SaveX(ctx context.Context) *Timer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TimerCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TimerCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TimerCreate) defaults() {
	if _, ok := tc.mutation.StartAt(); !ok {
		v := timer.DefaultStartAt()
		tc.mutation.SetStartAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := timer.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TimerCreate) check() error {
	if _, ok := tc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "Timer.start_at"`)}
	}
	if _, ok := tc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Timer.content"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Timer.user_id"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Timer.user"`)}
	}
	if _, ok := tc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required edge "Timer.subject"`)}
	}
	return nil
}

func (tc *TimerCreate) sqlSave(ctx context.Context) (*Timer, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TimerCreate) createSpec() (*Timer, *sqlgraph.CreateSpec) {
	var (
		_node = &Timer{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(timer.Table, sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.StartAt(); ok {
		_spec.SetField(timer.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := tc.mutation.Content(); ok {
		_spec.SetField(timer.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   timer.UserTable,
			Columns: []string{timer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.SubjectTable,
			Columns: []string{timer.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.subject_timers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SharedGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timer.SharedGroupTable,
			Columns: timer.SharedGroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TimerCreateBulk is the builder for creating many Timer entities in bulk.
type TimerCreateBulk struct {
	config
	err      error
	builders []*TimerCreate
}

// Save creates the Timer entities in the database.
func (tcb *TimerCreateBulk) Save(ctx context.Context) ([]*Timer, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Timer, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimerMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TimerCreateBulk) SaveX(ctx context.Context) []*Timer {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TimerCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TimerCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

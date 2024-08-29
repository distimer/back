// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/apnstoken"
	"pentag.kr/distimer/ent/predicate"
	"pentag.kr/distimer/ent/session"
)

// APNsTokenUpdate is the builder for updating APNsToken entities.
type APNsTokenUpdate struct {
	config
	hooks    []Hook
	mutation *APNsTokenMutation
}

// Where appends a list predicates to the APNsTokenUpdate builder.
func (antu *APNsTokenUpdate) Where(ps ...predicate.APNsToken) *APNsTokenUpdate {
	antu.mutation.Where(ps...)
	return antu
}

// SetStartToken sets the "start_token" field.
func (antu *APNsTokenUpdate) SetStartToken(s string) *APNsTokenUpdate {
	antu.mutation.SetStartToken(s)
	return antu
}

// SetNillableStartToken sets the "start_token" field if the given value is not nil.
func (antu *APNsTokenUpdate) SetNillableStartToken(s *string) *APNsTokenUpdate {
	if s != nil {
		antu.SetStartToken(*s)
	}
	return antu
}

// SetUpdateToken sets the "update_token" field.
func (antu *APNsTokenUpdate) SetUpdateToken(s string) *APNsTokenUpdate {
	antu.mutation.SetUpdateToken(s)
	return antu
}

// SetNillableUpdateToken sets the "update_token" field if the given value is not nil.
func (antu *APNsTokenUpdate) SetNillableUpdateToken(s *string) *APNsTokenUpdate {
	if s != nil {
		antu.SetUpdateToken(*s)
	}
	return antu
}

// SetSessionID sets the "session" edge to the Session entity by ID.
func (antu *APNsTokenUpdate) SetSessionID(id uuid.UUID) *APNsTokenUpdate {
	antu.mutation.SetSessionID(id)
	return antu
}

// SetSession sets the "session" edge to the Session entity.
func (antu *APNsTokenUpdate) SetSession(s *Session) *APNsTokenUpdate {
	return antu.SetSessionID(s.ID)
}

// Mutation returns the APNsTokenMutation object of the builder.
func (antu *APNsTokenUpdate) Mutation() *APNsTokenMutation {
	return antu.mutation
}

// ClearSession clears the "session" edge to the Session entity.
func (antu *APNsTokenUpdate) ClearSession() *APNsTokenUpdate {
	antu.mutation.ClearSession()
	return antu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (antu *APNsTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, antu.sqlSave, antu.mutation, antu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (antu *APNsTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := antu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (antu *APNsTokenUpdate) Exec(ctx context.Context) error {
	_, err := antu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (antu *APNsTokenUpdate) ExecX(ctx context.Context) {
	if err := antu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (antu *APNsTokenUpdate) check() error {
	if _, ok := antu.mutation.SessionID(); antu.mutation.SessionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APNsToken.session"`)
	}
	return nil
}

func (antu *APNsTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := antu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apnstoken.Table, apnstoken.Columns, sqlgraph.NewFieldSpec(apnstoken.FieldID, field.TypeInt))
	if ps := antu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := antu.mutation.StartToken(); ok {
		_spec.SetField(apnstoken.FieldStartToken, field.TypeString, value)
	}
	if value, ok := antu.mutation.UpdateToken(); ok {
		_spec.SetField(apnstoken.FieldUpdateToken, field.TypeString, value)
	}
	if antu.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apnstoken.SessionTable,
			Columns: []string{apnstoken.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := antu.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apnstoken.SessionTable,
			Columns: []string{apnstoken.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, antu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apnstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	antu.mutation.done = true
	return n, nil
}

// APNsTokenUpdateOne is the builder for updating a single APNsToken entity.
type APNsTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APNsTokenMutation
}

// SetStartToken sets the "start_token" field.
func (antuo *APNsTokenUpdateOne) SetStartToken(s string) *APNsTokenUpdateOne {
	antuo.mutation.SetStartToken(s)
	return antuo
}

// SetNillableStartToken sets the "start_token" field if the given value is not nil.
func (antuo *APNsTokenUpdateOne) SetNillableStartToken(s *string) *APNsTokenUpdateOne {
	if s != nil {
		antuo.SetStartToken(*s)
	}
	return antuo
}

// SetUpdateToken sets the "update_token" field.
func (antuo *APNsTokenUpdateOne) SetUpdateToken(s string) *APNsTokenUpdateOne {
	antuo.mutation.SetUpdateToken(s)
	return antuo
}

// SetNillableUpdateToken sets the "update_token" field if the given value is not nil.
func (antuo *APNsTokenUpdateOne) SetNillableUpdateToken(s *string) *APNsTokenUpdateOne {
	if s != nil {
		antuo.SetUpdateToken(*s)
	}
	return antuo
}

// SetSessionID sets the "session" edge to the Session entity by ID.
func (antuo *APNsTokenUpdateOne) SetSessionID(id uuid.UUID) *APNsTokenUpdateOne {
	antuo.mutation.SetSessionID(id)
	return antuo
}

// SetSession sets the "session" edge to the Session entity.
func (antuo *APNsTokenUpdateOne) SetSession(s *Session) *APNsTokenUpdateOne {
	return antuo.SetSessionID(s.ID)
}

// Mutation returns the APNsTokenMutation object of the builder.
func (antuo *APNsTokenUpdateOne) Mutation() *APNsTokenMutation {
	return antuo.mutation
}

// ClearSession clears the "session" edge to the Session entity.
func (antuo *APNsTokenUpdateOne) ClearSession() *APNsTokenUpdateOne {
	antuo.mutation.ClearSession()
	return antuo
}

// Where appends a list predicates to the APNsTokenUpdate builder.
func (antuo *APNsTokenUpdateOne) Where(ps ...predicate.APNsToken) *APNsTokenUpdateOne {
	antuo.mutation.Where(ps...)
	return antuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (antuo *APNsTokenUpdateOne) Select(field string, fields ...string) *APNsTokenUpdateOne {
	antuo.fields = append([]string{field}, fields...)
	return antuo
}

// Save executes the query and returns the updated APNsToken entity.
func (antuo *APNsTokenUpdateOne) Save(ctx context.Context) (*APNsToken, error) {
	return withHooks(ctx, antuo.sqlSave, antuo.mutation, antuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (antuo *APNsTokenUpdateOne) SaveX(ctx context.Context) *APNsToken {
	node, err := antuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (antuo *APNsTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := antuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (antuo *APNsTokenUpdateOne) ExecX(ctx context.Context) {
	if err := antuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (antuo *APNsTokenUpdateOne) check() error {
	if _, ok := antuo.mutation.SessionID(); antuo.mutation.SessionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "APNsToken.session"`)
	}
	return nil
}

func (antuo *APNsTokenUpdateOne) sqlSave(ctx context.Context) (_node *APNsToken, err error) {
	if err := antuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apnstoken.Table, apnstoken.Columns, sqlgraph.NewFieldSpec(apnstoken.FieldID, field.TypeInt))
	id, ok := antuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "APNsToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := antuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apnstoken.FieldID)
		for _, f := range fields {
			if !apnstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apnstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := antuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := antuo.mutation.StartToken(); ok {
		_spec.SetField(apnstoken.FieldStartToken, field.TypeString, value)
	}
	if value, ok := antuo.mutation.UpdateToken(); ok {
		_spec.SetField(apnstoken.FieldUpdateToken, field.TypeString, value)
	}
	if antuo.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apnstoken.SessionTable,
			Columns: []string{apnstoken.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := antuo.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apnstoken.SessionTable,
			Columns: []string{apnstoken.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &APNsToken{config: antuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, antuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apnstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	antuo.mutation.done = true
	return _node, nil
}
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
	"pentag.kr/distimer/ent/apnstoken"
	"pentag.kr/distimer/ent/fcmtoken"
	"pentag.kr/distimer/ent/session"
	"pentag.kr/distimer/ent/user"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetRefreshToken sets the "refresh_token" field.
func (sc *SessionCreate) SetRefreshToken(u uuid.UUID) *SessionCreate {
	sc.mutation.SetRefreshToken(u)
	return sc
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (sc *SessionCreate) SetNillableRefreshToken(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetRefreshToken(*u)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SessionCreate) SetCreatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetDeviceType sets the "device_type" field.
func (sc *SessionCreate) SetDeviceType(i int8) *SessionCreate {
	sc.mutation.SetDeviceType(i)
	return sc
}

// SetLastActive sets the "last_active" field.
func (sc *SessionCreate) SetLastActive(t time.Time) *SessionCreate {
	sc.mutation.SetLastActive(t)
	return sc
}

// SetNillableLastActive sets the "last_active" field if the given value is not nil.
func (sc *SessionCreate) SetNillableLastActive(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetLastActive(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SessionCreate) SetNillableID(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (sc *SessionCreate) SetOwnerID(id uuid.UUID) *SessionCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetOwner sets the "owner" edge to the User entity.
func (sc *SessionCreate) SetOwner(u *User) *SessionCreate {
	return sc.SetOwnerID(u.ID)
}

// SetApnsTokenID sets the "apns_token" edge to the APNsToken entity by ID.
func (sc *SessionCreate) SetApnsTokenID(id int) *SessionCreate {
	sc.mutation.SetApnsTokenID(id)
	return sc
}

// SetNillableApnsTokenID sets the "apns_token" edge to the APNsToken entity by ID if the given value is not nil.
func (sc *SessionCreate) SetNillableApnsTokenID(id *int) *SessionCreate {
	if id != nil {
		sc = sc.SetApnsTokenID(*id)
	}
	return sc
}

// SetApnsToken sets the "apns_token" edge to the APNsToken entity.
func (sc *SessionCreate) SetApnsToken(a *APNsToken) *SessionCreate {
	return sc.SetApnsTokenID(a.ID)
}

// SetFcmTokenID sets the "fcm_token" edge to the FCMToken entity by ID.
func (sc *SessionCreate) SetFcmTokenID(id int) *SessionCreate {
	sc.mutation.SetFcmTokenID(id)
	return sc
}

// SetNillableFcmTokenID sets the "fcm_token" edge to the FCMToken entity by ID if the given value is not nil.
func (sc *SessionCreate) SetNillableFcmTokenID(id *int) *SessionCreate {
	if id != nil {
		sc = sc.SetFcmTokenID(*id)
	}
	return sc
}

// SetFcmToken sets the "fcm_token" edge to the FCMToken entity.
func (sc *SessionCreate) SetFcmToken(f *FCMToken) *SessionCreate {
	return sc.SetFcmTokenID(f.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.RefreshToken(); !ok {
		v := session.DefaultRefreshToken()
		sc.mutation.SetRefreshToken(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := session.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.LastActive(); !ok {
		v := session.DefaultLastActive()
		sc.mutation.SetLastActive(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := session.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.RefreshToken(); !ok {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required field "Session.refresh_token"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Session.created_at"`)}
	}
	if _, ok := sc.mutation.DeviceType(); !ok {
		return &ValidationError{Name: "device_type", err: errors.New(`ent: missing required field "Session.device_type"`)}
	}
	if _, ok := sc.mutation.LastActive(); !ok {
		return &ValidationError{Name: "last_active", err: errors.New(`ent: missing required field "Session.last_active"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Session.owner"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.RefreshToken(); ok {
		_spec.SetField(session.FieldRefreshToken, field.TypeUUID, value)
		_node.RefreshToken = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.DeviceType(); ok {
		_spec.SetField(session.FieldDeviceType, field.TypeInt8, value)
		_node.DeviceType = value
	}
	if value, ok := sc.mutation.LastActive(); ok {
		_spec.SetField(session.FieldLastActive, field.TypeTime, value)
		_node.LastActive = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.OwnerTable,
			Columns: []string{session.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_sessions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ApnsTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   session.ApnsTokenTable,
			Columns: []string{session.ApnsTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apnstoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.FcmTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   session.FcmTokenTable,
			Columns: []string{session.FcmTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fcmtoken.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	err      error
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
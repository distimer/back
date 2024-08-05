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
	"pentag.kr/distimer/ent/category"
	"pentag.kr/distimer/ent/predicate"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
)

// SubjectUpdate is the builder for updating Subject entities.
type SubjectUpdate struct {
	config
	hooks    []Hook
	mutation *SubjectMutation
}

// Where appends a list predicates to the SubjectUpdate builder.
func (su *SubjectUpdate) Where(ps ...predicate.Subject) *SubjectUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SubjectUpdate) SetName(s string) *SubjectUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SubjectUpdate) SetNillableName(s *string) *SubjectUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetColor sets the "color" field.
func (su *SubjectUpdate) SetColor(s string) *SubjectUpdate {
	su.mutation.SetColor(s)
	return su
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (su *SubjectUpdate) SetNillableColor(s *string) *SubjectUpdate {
	if s != nil {
		su.SetColor(*s)
	}
	return su
}

// SetOrder sets the "order" field.
func (su *SubjectUpdate) SetOrder(i int8) *SubjectUpdate {
	su.mutation.ResetOrder()
	su.mutation.SetOrder(i)
	return su
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (su *SubjectUpdate) SetNillableOrder(i *int8) *SubjectUpdate {
	if i != nil {
		su.SetOrder(*i)
	}
	return su
}

// AddOrder adds i to the "order" field.
func (su *SubjectUpdate) AddOrder(i int8) *SubjectUpdate {
	su.mutation.AddOrder(i)
	return su
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (su *SubjectUpdate) SetCategoryID(id uuid.UUID) *SubjectUpdate {
	su.mutation.SetCategoryID(id)
	return su
}

// SetCategory sets the "category" edge to the Category entity.
func (su *SubjectUpdate) SetCategory(c *Category) *SubjectUpdate {
	return su.SetCategoryID(c.ID)
}

// AddStudyLogIDs adds the "study_logs" edge to the StudyLog entity by IDs.
func (su *SubjectUpdate) AddStudyLogIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.AddStudyLogIDs(ids...)
	return su
}

// AddStudyLogs adds the "study_logs" edges to the StudyLog entity.
func (su *SubjectUpdate) AddStudyLogs(s ...*StudyLog) *SubjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddStudyLogIDs(ids...)
}

// AddTimerIDs adds the "timers" edge to the Timer entity by IDs.
func (su *SubjectUpdate) AddTimerIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.AddTimerIDs(ids...)
	return su
}

// AddTimers adds the "timers" edges to the Timer entity.
func (su *SubjectUpdate) AddTimers(t ...*Timer) *SubjectUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTimerIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (su *SubjectUpdate) Mutation() *SubjectMutation {
	return su.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (su *SubjectUpdate) ClearCategory() *SubjectUpdate {
	su.mutation.ClearCategory()
	return su
}

// ClearStudyLogs clears all "study_logs" edges to the StudyLog entity.
func (su *SubjectUpdate) ClearStudyLogs() *SubjectUpdate {
	su.mutation.ClearStudyLogs()
	return su
}

// RemoveStudyLogIDs removes the "study_logs" edge to StudyLog entities by IDs.
func (su *SubjectUpdate) RemoveStudyLogIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.RemoveStudyLogIDs(ids...)
	return su
}

// RemoveStudyLogs removes "study_logs" edges to StudyLog entities.
func (su *SubjectUpdate) RemoveStudyLogs(s ...*StudyLog) *SubjectUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveStudyLogIDs(ids...)
}

// ClearTimers clears all "timers" edges to the Timer entity.
func (su *SubjectUpdate) ClearTimers() *SubjectUpdate {
	su.mutation.ClearTimers()
	return su
}

// RemoveTimerIDs removes the "timers" edge to Timer entities by IDs.
func (su *SubjectUpdate) RemoveTimerIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.RemoveTimerIDs(ids...)
	return su
}

// RemoveTimers removes "timers" edges to Timer entities.
func (su *SubjectUpdate) RemoveTimers(t ...*Timer) *SubjectUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTimerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubjectUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubjectUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubjectUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubjectUpdate) check() error {
	if _, ok := su.mutation.CategoryID(); su.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subject.category"`)
	}
	return nil
}

func (su *SubjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Color(); ok {
		_spec.SetField(subject.FieldColor, field.TypeString, value)
	}
	if value, ok := su.mutation.Order(); ok {
		_spec.SetField(subject.FieldOrder, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedOrder(); ok {
		_spec.AddField(subject.FieldOrder, field.TypeInt8, value)
	}
	if su.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.CategoryTable,
			Columns: []string{subject.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.CategoryTable,
			Columns: []string{subject.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StudyLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedStudyLogsIDs(); len(nodes) > 0 && !su.mutation.StudyLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StudyLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.TimersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTimersIDs(); len(nodes) > 0 && !su.mutation.TimersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TimersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubjectUpdateOne is the builder for updating a single Subject entity.
type SubjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubjectMutation
}

// SetName sets the "name" field.
func (suo *SubjectUpdateOne) SetName(s string) *SubjectUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableName(s *string) *SubjectUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetColor sets the "color" field.
func (suo *SubjectUpdateOne) SetColor(s string) *SubjectUpdateOne {
	suo.mutation.SetColor(s)
	return suo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableColor(s *string) *SubjectUpdateOne {
	if s != nil {
		suo.SetColor(*s)
	}
	return suo
}

// SetOrder sets the "order" field.
func (suo *SubjectUpdateOne) SetOrder(i int8) *SubjectUpdateOne {
	suo.mutation.ResetOrder()
	suo.mutation.SetOrder(i)
	return suo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableOrder(i *int8) *SubjectUpdateOne {
	if i != nil {
		suo.SetOrder(*i)
	}
	return suo
}

// AddOrder adds i to the "order" field.
func (suo *SubjectUpdateOne) AddOrder(i int8) *SubjectUpdateOne {
	suo.mutation.AddOrder(i)
	return suo
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (suo *SubjectUpdateOne) SetCategoryID(id uuid.UUID) *SubjectUpdateOne {
	suo.mutation.SetCategoryID(id)
	return suo
}

// SetCategory sets the "category" edge to the Category entity.
func (suo *SubjectUpdateOne) SetCategory(c *Category) *SubjectUpdateOne {
	return suo.SetCategoryID(c.ID)
}

// AddStudyLogIDs adds the "study_logs" edge to the StudyLog entity by IDs.
func (suo *SubjectUpdateOne) AddStudyLogIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.AddStudyLogIDs(ids...)
	return suo
}

// AddStudyLogs adds the "study_logs" edges to the StudyLog entity.
func (suo *SubjectUpdateOne) AddStudyLogs(s ...*StudyLog) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddStudyLogIDs(ids...)
}

// AddTimerIDs adds the "timers" edge to the Timer entity by IDs.
func (suo *SubjectUpdateOne) AddTimerIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.AddTimerIDs(ids...)
	return suo
}

// AddTimers adds the "timers" edges to the Timer entity.
func (suo *SubjectUpdateOne) AddTimers(t ...*Timer) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTimerIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (suo *SubjectUpdateOne) Mutation() *SubjectMutation {
	return suo.mutation
}

// ClearCategory clears the "category" edge to the Category entity.
func (suo *SubjectUpdateOne) ClearCategory() *SubjectUpdateOne {
	suo.mutation.ClearCategory()
	return suo
}

// ClearStudyLogs clears all "study_logs" edges to the StudyLog entity.
func (suo *SubjectUpdateOne) ClearStudyLogs() *SubjectUpdateOne {
	suo.mutation.ClearStudyLogs()
	return suo
}

// RemoveStudyLogIDs removes the "study_logs" edge to StudyLog entities by IDs.
func (suo *SubjectUpdateOne) RemoveStudyLogIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.RemoveStudyLogIDs(ids...)
	return suo
}

// RemoveStudyLogs removes "study_logs" edges to StudyLog entities.
func (suo *SubjectUpdateOne) RemoveStudyLogs(s ...*StudyLog) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveStudyLogIDs(ids...)
}

// ClearTimers clears all "timers" edges to the Timer entity.
func (suo *SubjectUpdateOne) ClearTimers() *SubjectUpdateOne {
	suo.mutation.ClearTimers()
	return suo
}

// RemoveTimerIDs removes the "timers" edge to Timer entities by IDs.
func (suo *SubjectUpdateOne) RemoveTimerIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.RemoveTimerIDs(ids...)
	return suo
}

// RemoveTimers removes "timers" edges to Timer entities.
func (suo *SubjectUpdateOne) RemoveTimers(t ...*Timer) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTimerIDs(ids...)
}

// Where appends a list predicates to the SubjectUpdate builder.
func (suo *SubjectUpdateOne) Where(ps ...predicate.Subject) *SubjectUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubjectUpdateOne) Select(field string, fields ...string) *SubjectUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subject entity.
func (suo *SubjectUpdateOne) Save(ctx context.Context) (*Subject, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveX(ctx context.Context) *Subject {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubjectUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubjectUpdateOne) check() error {
	if _, ok := suo.mutation.CategoryID(); suo.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Subject.category"`)
	}
	return nil
}

func (suo *SubjectUpdateOne) sqlSave(ctx context.Context) (_node *Subject, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(subject.Table, subject.Columns, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for _, f := range fields {
			if !subject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Color(); ok {
		_spec.SetField(subject.FieldColor, field.TypeString, value)
	}
	if value, ok := suo.mutation.Order(); ok {
		_spec.SetField(subject.FieldOrder, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedOrder(); ok {
		_spec.AddField(subject.FieldOrder, field.TypeInt8, value)
	}
	if suo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.CategoryTable,
			Columns: []string{subject.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.CategoryTable,
			Columns: []string{subject.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StudyLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedStudyLogsIDs(); len(nodes) > 0 && !suo.mutation.StudyLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StudyLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.StudyLogsTable,
			Columns: []string{subject.StudyLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(studylog.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.TimersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTimersIDs(); len(nodes) > 0 && !suo.mutation.TimersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TimersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.TimersTable,
			Columns: []string{subject.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subject{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

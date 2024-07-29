// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
)

// Timer is the model entity for the Timer schema.
type Timer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimerQuery when eager-loading is set.
	Edges          TimerEdges `json:"edges"`
	subject_timers *uuid.UUID
	selectValues   sql.SelectValues
}

// TimerEdges holds the relations/edges for other nodes in the graph.
type TimerEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Subject holds the value of the subject edge.
	Subject *Subject `json:"subject,omitempty"`
	// SharedGroup holds the value of the shared_group edge.
	SharedGroup []*Group `json:"shared_group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimerEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimerEdges) SubjectOrErr() (*Subject, error) {
	if e.Subject != nil {
		return e.Subject, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: subject.Label}
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// SharedGroupOrErr returns the SharedGroup value or an error if the edge
// was not loaded in eager-loading.
func (e TimerEdges) SharedGroupOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.SharedGroup, nil
	}
	return nil, &NotLoadedError{edge: "shared_group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Timer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timer.FieldContent:
			values[i] = new(sql.NullString)
		case timer.FieldStartAt:
			values[i] = new(sql.NullTime)
		case timer.FieldID, timer.FieldUserID:
			values[i] = new(uuid.UUID)
		case timer.ForeignKeys[0]: // subject_timers
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Timer fields.
func (t *Timer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case timer.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				t.StartAt = value.Time
			}
		case timer.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				t.Content = value.String
			}
		case timer.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				t.UserID = *value
			}
		case timer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field subject_timers", values[i])
			} else if value.Valid {
				t.subject_timers = new(uuid.UUID)
				*t.subject_timers = *value.S.(*uuid.UUID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Timer.
// This includes values selected through modifiers, order, etc.
func (t *Timer) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Timer entity.
func (t *Timer) QueryUser() *UserQuery {
	return NewTimerClient(t.config).QueryUser(t)
}

// QuerySubject queries the "subject" edge of the Timer entity.
func (t *Timer) QuerySubject() *SubjectQuery {
	return NewTimerClient(t.config).QuerySubject(t)
}

// QuerySharedGroup queries the "shared_group" edge of the Timer entity.
func (t *Timer) QuerySharedGroup() *GroupQuery {
	return NewTimerClient(t.config).QuerySharedGroup(t)
}

// Update returns a builder for updating this Timer.
// Note that you need to call Timer.Unwrap() before calling this method if this Timer
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Timer) Update() *TimerUpdateOne {
	return NewTimerClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Timer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Timer) Unwrap() *Timer {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Timer is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Timer) String() string {
	var builder strings.Builder
	builder.WriteString("Timer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("start_at=")
	builder.WriteString(t.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(t.Content)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Timers is a parsable slice of Timer.
type Timers []*Timer

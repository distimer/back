// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/group"
	"pentag.kr/distimer/ent/invitecode"
)

// InviteCode is the model entity for the InviteCode schema.
type InviteCode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Used holds the value of the "used" field.
	Used int32 `json:"used,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InviteCodeQuery when eager-loading is set.
	Edges              InviteCodeEdges `json:"edges"`
	group_invite_codes *uuid.UUID
	selectValues       sql.SelectValues
}

// InviteCodeEdges holds the relations/edges for other nodes in the graph.
type InviteCodeEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InviteCodeEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InviteCode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invitecode.FieldID, invitecode.FieldUsed:
			values[i] = new(sql.NullInt64)
		case invitecode.FieldCode:
			values[i] = new(sql.NullString)
		case invitecode.ForeignKeys[0]: // group_invite_codes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InviteCode fields.
func (ic *InviteCode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case invitecode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ic.ID = int(value.Int64)
		case invitecode.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ic.Code = value.String
			}
		case invitecode.FieldUsed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field used", values[i])
			} else if value.Valid {
				ic.Used = int32(value.Int64)
			}
		case invitecode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_invite_codes", values[i])
			} else if value.Valid {
				ic.group_invite_codes = new(uuid.UUID)
				*ic.group_invite_codes = *value.S.(*uuid.UUID)
			}
		default:
			ic.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InviteCode.
// This includes values selected through modifiers, order, etc.
func (ic *InviteCode) Value(name string) (ent.Value, error) {
	return ic.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the InviteCode entity.
func (ic *InviteCode) QueryGroup() *GroupQuery {
	return NewInviteCodeClient(ic.config).QueryGroup(ic)
}

// Update returns a builder for updating this InviteCode.
// Note that you need to call InviteCode.Unwrap() before calling this method if this InviteCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *InviteCode) Update() *InviteCodeUpdateOne {
	return NewInviteCodeClient(ic.config).UpdateOne(ic)
}

// Unwrap unwraps the InviteCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *InviteCode) Unwrap() *InviteCode {
	_tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: InviteCode is not a transactional entity")
	}
	ic.config.driver = _tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *InviteCode) String() string {
	var builder strings.Builder
	builder.WriteString("InviteCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ic.ID))
	builder.WriteString("code=")
	builder.WriteString(ic.Code)
	builder.WriteString(", ")
	builder.WriteString("used=")
	builder.WriteString(fmt.Sprintf("%v", ic.Used))
	builder.WriteByte(')')
	return builder.String()
}

// InviteCodes is a parsable slice of InviteCode.
type InviteCodes []*InviteCode

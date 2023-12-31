// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/linelog"
	"entdemo/ent/lineuser"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// LineLog is the model entity for the LineLog schema.
type LineLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LineLogQuery when eager-loading is set.
	Edges              LineLogEdges `json:"edges"`
	line_user_linelogs *int
	selectValues       sql.SelectValues
}

// LineLogEdges holds the relations/edges for other nodes in the graph.
type LineLogEdges struct {
	// Owner holds the value of the owner edge.
	Owner *LineUser `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LineLogEdges) OwnerOrErr() (*LineUser, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: lineuser.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LineLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case linelog.FieldID:
			values[i] = new(sql.NullInt64)
		case linelog.FieldAction, linelog.FieldMessage:
			values[i] = new(sql.NullString)
		case linelog.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case linelog.ForeignKeys[0]: // line_user_linelogs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LineLog fields.
func (ll *LineLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case linelog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ll.ID = int(value.Int64)
		case linelog.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				ll.Action = value.String
			}
		case linelog.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				ll.Message = value.String
			}
		case linelog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ll.CreatedAt = value.Time
			}
		case linelog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field line_user_linelogs", value)
			} else if value.Valid {
				ll.line_user_linelogs = new(int)
				*ll.line_user_linelogs = int(value.Int64)
			}
		default:
			ll.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LineLog.
// This includes values selected through modifiers, order, etc.
func (ll *LineLog) Value(name string) (ent.Value, error) {
	return ll.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the LineLog entity.
func (ll *LineLog) QueryOwner() *LineUserQuery {
	return NewLineLogClient(ll.config).QueryOwner(ll)
}

// Update returns a builder for updating this LineLog.
// Note that you need to call LineLog.Unwrap() before calling this method if this LineLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ll *LineLog) Update() *LineLogUpdateOne {
	return NewLineLogClient(ll.config).UpdateOne(ll)
}

// Unwrap unwraps the LineLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ll *LineLog) Unwrap() *LineLog {
	_tx, ok := ll.config.driver.(*txDriver)
	if !ok {
		panic("ent: LineLog is not a transactional entity")
	}
	ll.config.driver = _tx.drv
	return ll
}

// String implements the fmt.Stringer.
func (ll *LineLog) String() string {
	var builder strings.Builder
	builder.WriteString("LineLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ll.ID))
	builder.WriteString("action=")
	builder.WriteString(ll.Action)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(ll.Message)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ll.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LineLogs is a parsable slice of LineLog.
type LineLogs []*LineLog

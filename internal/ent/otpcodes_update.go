// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pets-backend/internal/ent/otpcodes"
	"pets-backend/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OtpCodesUpdate is the builder for updating OtpCodes entities.
type OtpCodesUpdate struct {
	config
	hooks    []Hook
	mutation *OtpCodesMutation
}

// Where appends a list predicates to the OtpCodesUpdate builder.
func (ocu *OtpCodesUpdate) Where(ps ...predicate.OtpCodes) *OtpCodesUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetCode sets the "code" field.
func (ocu *OtpCodesUpdate) SetCode(s string) *OtpCodesUpdate {
	ocu.mutation.SetCode(s)
	return ocu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ocu *OtpCodesUpdate) SetNillableCode(s *string) *OtpCodesUpdate {
	if s != nil {
		ocu.SetCode(*s)
	}
	return ocu
}

// SetEmail sets the "email" field.
func (ocu *OtpCodesUpdate) SetEmail(s string) *OtpCodesUpdate {
	ocu.mutation.SetEmail(s)
	return ocu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ocu *OtpCodesUpdate) SetNillableEmail(s *string) *OtpCodesUpdate {
	if s != nil {
		ocu.SetEmail(*s)
	}
	return ocu
}

// SetNextSendTime sets the "next_send_time" field.
func (ocu *OtpCodesUpdate) SetNextSendTime(t time.Time) *OtpCodesUpdate {
	ocu.mutation.SetNextSendTime(t)
	return ocu
}

// SetNillableNextSendTime sets the "next_send_time" field if the given value is not nil.
func (ocu *OtpCodesUpdate) SetNillableNextSendTime(t *time.Time) *OtpCodesUpdate {
	if t != nil {
		ocu.SetNextSendTime(*t)
	}
	return ocu
}

// SetExparationTime sets the "exparation_time" field.
func (ocu *OtpCodesUpdate) SetExparationTime(t time.Time) *OtpCodesUpdate {
	ocu.mutation.SetExparationTime(t)
	return ocu
}

// SetNillableExparationTime sets the "exparation_time" field if the given value is not nil.
func (ocu *OtpCodesUpdate) SetNillableExparationTime(t *time.Time) *OtpCodesUpdate {
	if t != nil {
		ocu.SetExparationTime(*t)
	}
	return ocu
}

// Mutation returns the OtpCodesMutation object of the builder.
func (ocu *OtpCodesUpdate) Mutation() *OtpCodesMutation {
	return ocu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OtpCodesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ocu.sqlSave, ocu.mutation, ocu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocu *OtpCodesUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OtpCodesUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OtpCodesUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocu *OtpCodesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(otpcodes.Table, otpcodes.Columns, sqlgraph.NewFieldSpec(otpcodes.FieldID, field.TypeUUID))
	if ps := ocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocu.mutation.Code(); ok {
		_spec.SetField(otpcodes.FieldCode, field.TypeString, value)
	}
	if value, ok := ocu.mutation.Email(); ok {
		_spec.SetField(otpcodes.FieldEmail, field.TypeString, value)
	}
	if value, ok := ocu.mutation.NextSendTime(); ok {
		_spec.SetField(otpcodes.FieldNextSendTime, field.TypeTime, value)
	}
	if value, ok := ocu.mutation.ExparationTime(); ok {
		_spec.SetField(otpcodes.FieldExparationTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{otpcodes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ocu.mutation.done = true
	return n, nil
}

// OtpCodesUpdateOne is the builder for updating a single OtpCodes entity.
type OtpCodesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OtpCodesMutation
}

// SetCode sets the "code" field.
func (ocuo *OtpCodesUpdateOne) SetCode(s string) *OtpCodesUpdateOne {
	ocuo.mutation.SetCode(s)
	return ocuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ocuo *OtpCodesUpdateOne) SetNillableCode(s *string) *OtpCodesUpdateOne {
	if s != nil {
		ocuo.SetCode(*s)
	}
	return ocuo
}

// SetEmail sets the "email" field.
func (ocuo *OtpCodesUpdateOne) SetEmail(s string) *OtpCodesUpdateOne {
	ocuo.mutation.SetEmail(s)
	return ocuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ocuo *OtpCodesUpdateOne) SetNillableEmail(s *string) *OtpCodesUpdateOne {
	if s != nil {
		ocuo.SetEmail(*s)
	}
	return ocuo
}

// SetNextSendTime sets the "next_send_time" field.
func (ocuo *OtpCodesUpdateOne) SetNextSendTime(t time.Time) *OtpCodesUpdateOne {
	ocuo.mutation.SetNextSendTime(t)
	return ocuo
}

// SetNillableNextSendTime sets the "next_send_time" field if the given value is not nil.
func (ocuo *OtpCodesUpdateOne) SetNillableNextSendTime(t *time.Time) *OtpCodesUpdateOne {
	if t != nil {
		ocuo.SetNextSendTime(*t)
	}
	return ocuo
}

// SetExparationTime sets the "exparation_time" field.
func (ocuo *OtpCodesUpdateOne) SetExparationTime(t time.Time) *OtpCodesUpdateOne {
	ocuo.mutation.SetExparationTime(t)
	return ocuo
}

// SetNillableExparationTime sets the "exparation_time" field if the given value is not nil.
func (ocuo *OtpCodesUpdateOne) SetNillableExparationTime(t *time.Time) *OtpCodesUpdateOne {
	if t != nil {
		ocuo.SetExparationTime(*t)
	}
	return ocuo
}

// Mutation returns the OtpCodesMutation object of the builder.
func (ocuo *OtpCodesUpdateOne) Mutation() *OtpCodesMutation {
	return ocuo.mutation
}

// Where appends a list predicates to the OtpCodesUpdate builder.
func (ocuo *OtpCodesUpdateOne) Where(ps ...predicate.OtpCodes) *OtpCodesUpdateOne {
	ocuo.mutation.Where(ps...)
	return ocuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OtpCodesUpdateOne) Select(field string, fields ...string) *OtpCodesUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OtpCodes entity.
func (ocuo *OtpCodesUpdateOne) Save(ctx context.Context) (*OtpCodes, error) {
	return withHooks(ctx, ocuo.sqlSave, ocuo.mutation, ocuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocuo *OtpCodesUpdateOne) SaveX(ctx context.Context) *OtpCodes {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OtpCodesUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OtpCodesUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ocuo *OtpCodesUpdateOne) sqlSave(ctx context.Context) (_node *OtpCodes, err error) {
	_spec := sqlgraph.NewUpdateSpec(otpcodes.Table, otpcodes.Columns, sqlgraph.NewFieldSpec(otpcodes.FieldID, field.TypeUUID))
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OtpCodes.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, otpcodes.FieldID)
		for _, f := range fields {
			if !otpcodes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != otpcodes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocuo.mutation.Code(); ok {
		_spec.SetField(otpcodes.FieldCode, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.Email(); ok {
		_spec.SetField(otpcodes.FieldEmail, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.NextSendTime(); ok {
		_spec.SetField(otpcodes.FieldNextSendTime, field.TypeTime, value)
	}
	if value, ok := ocuo.mutation.ExparationTime(); ok {
		_spec.SetField(otpcodes.FieldExparationTime, field.TypeTime, value)
	}
	_node = &OtpCodes{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{otpcodes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ocuo.mutation.done = true
	return _node, nil
}

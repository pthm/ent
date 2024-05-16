// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/attachedfile"
	"entgo.io/ent/entc/integration/edgeschema/ent/file"
	"entgo.io/ent/entc/integration/edgeschema/ent/process"
	"entgo.io/ent/schema/field"
)

// AttachedFileCreate is the builder for creating a AttachedFile entity.
type AttachedFileCreate struct {
	config
	mutation *AttachedFileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAttachTime sets the "attach_time" field.
func (afc *AttachedFileCreate) SetAttachTime(t time.Time) *AttachedFileCreate {
	afc.mutation.SetAttachTime(t)
	return afc
}

// SetNillableAttachTime sets the "attach_time" field if the given value is not nil.
func (afc *AttachedFileCreate) SetNillableAttachTime(t *time.Time) *AttachedFileCreate {
	if t != nil {
		afc.SetAttachTime(*t)
	}
	return afc
}

// SetFID sets the "f_id" field.
func (afc *AttachedFileCreate) SetFID(i int) *AttachedFileCreate {
	afc.mutation.SetFID(i)
	return afc
}

// SetProcID sets the "proc_id" field.
func (afc *AttachedFileCreate) SetProcID(i int) *AttachedFileCreate {
	afc.mutation.SetProcID(i)
	return afc
}

// SetFiID sets the "fi" edge to the File entity by ID.
func (afc *AttachedFileCreate) SetFiID(id int) *AttachedFileCreate {
	afc.mutation.SetFiID(id)
	return afc
}

// SetFi sets the "fi" edge to the File entity.
func (afc *AttachedFileCreate) SetFi(f *File) *AttachedFileCreate {
	return afc.SetFiID(f.ID)
}

// SetProc sets the "proc" edge to the Process entity.
func (afc *AttachedFileCreate) SetProc(p *Process) *AttachedFileCreate {
	return afc.SetProcID(p.ID)
}

// Mutation returns the AttachedFileMutation object of the builder.
func (afc *AttachedFileCreate) Mutation() *AttachedFileMutation {
	return afc.mutation
}

// Save creates the AttachedFile in the database.
func (afc *AttachedFileCreate) Save(ctx context.Context) (*AttachedFile, error) {
	afc.defaults()
	return withHooks(ctx, afc.sqlSave, afc.mutation, afc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (afc *AttachedFileCreate) SaveX(ctx context.Context) *AttachedFile {
	v, err := afc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afc *AttachedFileCreate) Exec(ctx context.Context) error {
	_, err := afc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afc *AttachedFileCreate) ExecX(ctx context.Context) {
	if err := afc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afc *AttachedFileCreate) defaults() {
	if _, ok := afc.mutation.AttachTime(); !ok {
		v := attachedfile.DefaultAttachTime()
		afc.mutation.SetAttachTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afc *AttachedFileCreate) check() error {
	if _, ok := afc.mutation.AttachTime(); !ok {
		return &ValidationError{Name: "attach_time", err: errors.New(`ent: missing required field "AttachedFile.attach_time"`)}
	}
	if _, ok := afc.mutation.FID(); !ok {
		return &ValidationError{Name: "f_id", err: errors.New(`ent: missing required field "AttachedFile.f_id"`)}
	}
	if _, ok := afc.mutation.ProcID(); !ok {
		return &ValidationError{Name: "proc_id", err: errors.New(`ent: missing required field "AttachedFile.proc_id"`)}
	}
	if _, ok := afc.mutation.FiID(); !ok {
		return &ValidationError{Name: "fi", err: errors.New(`ent: missing required edge "AttachedFile.fi"`)}
	}
	if _, ok := afc.mutation.ProcID(); !ok {
		return &ValidationError{Name: "proc", err: errors.New(`ent: missing required edge "AttachedFile.proc"`)}
	}
	return nil
}

func (afc *AttachedFileCreate) sqlSave(ctx context.Context) (*AttachedFile, error) {
	if err := afc.check(); err != nil {
		return nil, err
	}
	_node, _spec := afc.createSpec()
	if err := sqlgraph.CreateNode(ctx, afc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	afc.mutation.id = &_node.ID
	afc.mutation.done = true
	return _node, nil
}

func (afc *AttachedFileCreate) createSpec() (*AttachedFile, *sqlgraph.CreateSpec) {
	var (
		_node = &AttachedFile{config: afc.config}
		_spec = sqlgraph.NewCreateSpec(attachedfile.Table, sqlgraph.NewFieldSpec(attachedfile.FieldID, field.TypeInt))
	)
	_spec.OnConflict = afc.conflict
	if value, ok := afc.mutation.AttachTime(); ok {
		_spec.SetField(attachedfile.FieldAttachTime, field.TypeTime, value)
		_node.AttachTime = value
	}
	if nodes := afc.mutation.FiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachedfile.FiTable,
			Columns: []string{attachedfile.FiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := afc.mutation.ProcIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attachedfile.ProcTable,
			Columns: []string{attachedfile.ProcColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProcID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttachedFile.Create().
//		SetAttachTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachedFileUpsert) {
//			SetAttachTime(v+v).
//		}).
//		Exec(ctx)
func (afc *AttachedFileCreate) OnConflict(opts ...sql.ConflictOption) *AttachedFileUpsertOne {
	afc.conflict = opts
	return &AttachedFileUpsertOne{
		create: afc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (afc *AttachedFileCreate) OnConflictColumns(columns ...string) *AttachedFileUpsertOne {
	afc.conflict = append(afc.conflict, sql.ConflictColumns(columns...))
	return &AttachedFileUpsertOne{
		create: afc,
	}
}

type (
	// AttachedFileUpsertOne is the builder for "upsert"-ing
	//  one AttachedFile node.
	AttachedFileUpsertOne struct {
		create *AttachedFileCreate
	}

	// AttachedFileUpsert is the "OnConflict" setter.
	AttachedFileUpsert struct {
		*sql.UpdateSet
	}
)

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsert) SetAttachTime(v time.Time) *AttachedFileUpsert {
	u.Set(attachedfile.FieldAttachTime, v)
	return u
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateAttachTime() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldAttachTime)
	return u
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsert) SetFID(v int) *AttachedFileUpsert {
	u.Set(attachedfile.FieldFID, v)
	return u
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateFID() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldFID)
	return u
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsert) SetProcID(v int) *AttachedFileUpsert {
	u.Set(attachedfile.FieldProcID, v)
	return u
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsert) UpdateProcID() *AttachedFileUpsert {
	u.SetExcluded(attachedfile.FieldProcID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttachedFileUpsertOne) UpdateNewValues() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttachedFileUpsertOne) Ignore() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachedFileUpsertOne) DoNothing() *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachedFileCreate.OnConflict
// documentation for more info.
func (u *AttachedFileUpsertOne) Update(set func(*AttachedFileUpsert)) *AttachedFileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachedFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsertOne) SetAttachTime(v time.Time) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetAttachTime(v)
	})
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateAttachTime() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateAttachTime()
	})
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsertOne) SetFID(v int) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetFID(v)
	})
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateFID() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateFID()
	})
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsertOne) SetProcID(v int) *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetProcID(v)
	})
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsertOne) UpdateProcID() *AttachedFileUpsertOne {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateProcID()
	})
}

// Exec executes the query.
func (u *AttachedFileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachedFileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachedFileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttachedFileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttachedFileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttachedFileCreateBulk is the builder for creating many AttachedFile entities in bulk.
type AttachedFileCreateBulk struct {
	config
	err      error
	builders []*AttachedFileCreate
	conflict []sql.ConflictOption
}

// Save creates the AttachedFile entities in the database.
func (afcb *AttachedFileCreateBulk) Save(ctx context.Context) ([]*AttachedFile, error) {
	if afcb.err != nil {
		return nil, afcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(afcb.builders))
	nodes := make([]*AttachedFile, len(afcb.builders))
	mutators := make([]Mutator, len(afcb.builders))
	for i := range afcb.builders {
		func(i int, root context.Context) {
			builder := afcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachedFileMutation)
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
					_, err = mutators[i+1].Mutate(root, afcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = afcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, afcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, afcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (afcb *AttachedFileCreateBulk) SaveX(ctx context.Context) []*AttachedFile {
	v, err := afcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afcb *AttachedFileCreateBulk) Exec(ctx context.Context) error {
	_, err := afcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afcb *AttachedFileCreateBulk) ExecX(ctx context.Context) {
	if err := afcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AttachedFile.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttachedFileUpsert) {
//			SetAttachTime(v+v).
//		}).
//		Exec(ctx)
func (afcb *AttachedFileCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttachedFileUpsertBulk {
	afcb.conflict = opts
	return &AttachedFileUpsertBulk{
		create: afcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (afcb *AttachedFileCreateBulk) OnConflictColumns(columns ...string) *AttachedFileUpsertBulk {
	afcb.conflict = append(afcb.conflict, sql.ConflictColumns(columns...))
	return &AttachedFileUpsertBulk{
		create: afcb,
	}
}

// AttachedFileUpsertBulk is the builder for "upsert"-ing
// a bulk of AttachedFile nodes.
type AttachedFileUpsertBulk struct {
	create *AttachedFileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttachedFileUpsertBulk) UpdateNewValues() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AttachedFile.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttachedFileUpsertBulk) Ignore() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttachedFileUpsertBulk) DoNothing() *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttachedFileCreateBulk.OnConflict
// documentation for more info.
func (u *AttachedFileUpsertBulk) Update(set func(*AttachedFileUpsert)) *AttachedFileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttachedFileUpsert{UpdateSet: update})
	}))
	return u
}

// SetAttachTime sets the "attach_time" field.
func (u *AttachedFileUpsertBulk) SetAttachTime(v time.Time) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetAttachTime(v)
	})
}

// UpdateAttachTime sets the "attach_time" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateAttachTime() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateAttachTime()
	})
}

// SetFID sets the "f_id" field.
func (u *AttachedFileUpsertBulk) SetFID(v int) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetFID(v)
	})
}

// UpdateFID sets the "f_id" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateFID() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateFID()
	})
}

// SetProcID sets the "proc_id" field.
func (u *AttachedFileUpsertBulk) SetProcID(v int) *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.SetProcID(v)
	})
}

// UpdateProcID sets the "proc_id" field to the value that was provided on create.
func (u *AttachedFileUpsertBulk) UpdateProcID() *AttachedFileUpsertBulk {
	return u.Update(func(s *AttachedFileUpsert) {
		s.UpdateProcID()
	})
}

// Exec executes the query.
func (u *AttachedFileUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if b == nil {
			return fmt.Errorf("ent: missing builder at index %d, unexpected nil builder passed to CreateBulk", i)
		}
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttachedFileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttachedFileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttachedFileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

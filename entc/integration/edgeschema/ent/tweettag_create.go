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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetTagCreate is the builder for creating a TweetTag entity.
type TweetTagCreate struct {
	config
	mutation *TweetTagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAddedAt sets the "added_at" field.
func (ttc *TweetTagCreate) SetAddedAt(t time.Time) *TweetTagCreate {
	ttc.mutation.SetAddedAt(t)
	return ttc
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (ttc *TweetTagCreate) SetNillableAddedAt(t *time.Time) *TweetTagCreate {
	if t != nil {
		ttc.SetAddedAt(*t)
	}
	return ttc
}

// SetTagID sets the "tag_id" field.
func (ttc *TweetTagCreate) SetTagID(i int) *TweetTagCreate {
	ttc.mutation.SetTagID(i)
	return ttc
}

// SetTweetID sets the "tweet_id" field.
func (ttc *TweetTagCreate) SetTweetID(i int) *TweetTagCreate {
	ttc.mutation.SetTweetID(i)
	return ttc
}

// SetID sets the "id" field.
func (ttc *TweetTagCreate) SetID(u uuid.UUID) *TweetTagCreate {
	ttc.mutation.SetID(u)
	return ttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttc *TweetTagCreate) SetNillableID(u *uuid.UUID) *TweetTagCreate {
	if u != nil {
		ttc.SetID(*u)
	}
	return ttc
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttc *TweetTagCreate) SetTag(t *Tag) *TweetTagCreate {
	return ttc.SetTagID(t.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (ttc *TweetTagCreate) SetTweet(t *Tweet) *TweetTagCreate {
	return ttc.SetTweetID(t.ID)
}

// Mutation returns the TweetTagMutation object of the builder.
func (ttc *TweetTagCreate) Mutation() *TweetTagMutation {
	return ttc.mutation
}

// Save creates the TweetTag in the database.
func (ttc *TweetTagCreate) Save(ctx context.Context) (*TweetTag, error) {
	ttc.defaults()
	return withHooks(ctx, ttc.sqlSave, ttc.mutation, ttc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TweetTagCreate) SaveX(ctx context.Context) *TweetTag {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TweetTagCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TweetTagCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TweetTagCreate) defaults() {
	if _, ok := ttc.mutation.AddedAt(); !ok {
		v := tweettag.DefaultAddedAt()
		ttc.mutation.SetAddedAt(v)
	}
	if _, ok := ttc.mutation.ID(); !ok {
		v := tweettag.DefaultID()
		ttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TweetTagCreate) check() error {
	if _, ok := ttc.mutation.AddedAt(); !ok {
		return &ValidationError{Name: "added_at", err: errors.New(`ent: missing required field "TweetTag.added_at"`)}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "TweetTag.tag_id"`)}
	}
	if _, ok := ttc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "TweetTag.tweet_id"`)}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required edge "TweetTag.tag"`)}
	}
	if _, ok := ttc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "TweetTag.tweet"`)}
	}
	return nil
}

func (ttc *TweetTagCreate) sqlSave(ctx context.Context) (*TweetTag, error) {
	if err := ttc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
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
	ttc.mutation.id = &_node.ID
	ttc.mutation.done = true
	return _node, nil
}

func (ttc *TweetTagCreate) createSpec() (*TweetTag, *sqlgraph.CreateSpec) {
	var (
		_node = &TweetTag{config: ttc.config}
		_spec = sqlgraph.NewCreateSpec(tweettag.Table, sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ttc.conflict
	if id, ok := ttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ttc.mutation.AddedAt(); ok {
		_spec.SetField(tweettag.FieldAddedAt, field.TypeTime, value)
		_node.AddedAt = value
	}
	if nodes := ttc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TweetTag.Create().
//		SetAddedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TweetTagUpsert) {
//			SetAddedAt(v+v).
//		}).
//		Exec(ctx)
func (ttc *TweetTagCreate) OnConflict(opts ...sql.ConflictOption) *TweetTagUpsertOne {
	ttc.conflict = opts
	return &TweetTagUpsertOne{
		create: ttc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ttc *TweetTagCreate) OnConflictColumns(columns ...string) *TweetTagUpsertOne {
	ttc.conflict = append(ttc.conflict, sql.ConflictColumns(columns...))
	return &TweetTagUpsertOne{
		create: ttc,
	}
}

type (
	// TweetTagUpsertOne is the builder for "upsert"-ing
	//  one TweetTag node.
	TweetTagUpsertOne struct {
		create *TweetTagCreate
	}

	// TweetTagUpsert is the "OnConflict" setter.
	TweetTagUpsert struct {
		*sql.UpdateSet
	}
)

// SetAddedAt sets the "added_at" field.
func (u *TweetTagUpsert) SetAddedAt(v time.Time) *TweetTagUpsert {
	u.Set(tweettag.FieldAddedAt, v)
	return u
}

// UpdateAddedAt sets the "added_at" field to the value that was provided on create.
func (u *TweetTagUpsert) UpdateAddedAt() *TweetTagUpsert {
	u.SetExcluded(tweettag.FieldAddedAt)
	return u
}

// SetTagID sets the "tag_id" field.
func (u *TweetTagUpsert) SetTagID(v int) *TweetTagUpsert {
	u.Set(tweettag.FieldTagID, v)
	return u
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TweetTagUpsert) UpdateTagID() *TweetTagUpsert {
	u.SetExcluded(tweettag.FieldTagID)
	return u
}

// SetTweetID sets the "tweet_id" field.
func (u *TweetTagUpsert) SetTweetID(v int) *TweetTagUpsert {
	u.Set(tweettag.FieldTweetID, v)
	return u
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *TweetTagUpsert) UpdateTweetID() *TweetTagUpsert {
	u.SetExcluded(tweettag.FieldTweetID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tweettag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TweetTagUpsertOne) UpdateNewValues() *TweetTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tweettag.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TweetTagUpsertOne) Ignore() *TweetTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TweetTagUpsertOne) DoNothing() *TweetTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TweetTagCreate.OnConflict
// documentation for more info.
func (u *TweetTagUpsertOne) Update(set func(*TweetTagUpsert)) *TweetTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TweetTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddedAt sets the "added_at" field.
func (u *TweetTagUpsertOne) SetAddedAt(v time.Time) *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetAddedAt(v)
	})
}

// UpdateAddedAt sets the "added_at" field to the value that was provided on create.
func (u *TweetTagUpsertOne) UpdateAddedAt() *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateAddedAt()
	})
}

// SetTagID sets the "tag_id" field.
func (u *TweetTagUpsertOne) SetTagID(v int) *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TweetTagUpsertOne) UpdateTagID() *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateTagID()
	})
}

// SetTweetID sets the "tweet_id" field.
func (u *TweetTagUpsertOne) SetTweetID(v int) *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetTweetID(v)
	})
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *TweetTagUpsertOne) UpdateTweetID() *TweetTagUpsertOne {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateTweetID()
	})
}

// Exec executes the query.
func (u *TweetTagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TweetTagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetTagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TweetTagUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TweetTagUpsertOne.ID is not supported by MySQL driver. Use TweetTagUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TweetTagUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TweetTagCreateBulk is the builder for creating many TweetTag entities in bulk.
type TweetTagCreateBulk struct {
	config
	err      error
	builders []*TweetTagCreate
	conflict []sql.ConflictOption
}

// Save creates the TweetTag entities in the database.
func (ttcb *TweetTagCreateBulk) Save(ctx context.Context) ([]*TweetTag, error) {
	if ttcb.err != nil {
		return nil, ttcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TweetTag, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TweetTagMutation)
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
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ttcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TweetTagCreateBulk) SaveX(ctx context.Context) []*TweetTag {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TweetTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TweetTagCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TweetTag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TweetTagUpsert) {
//			SetAddedAt(v+v).
//		}).
//		Exec(ctx)
func (ttcb *TweetTagCreateBulk) OnConflict(opts ...sql.ConflictOption) *TweetTagUpsertBulk {
	ttcb.conflict = opts
	return &TweetTagUpsertBulk{
		create: ttcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ttcb *TweetTagCreateBulk) OnConflictColumns(columns ...string) *TweetTagUpsertBulk {
	ttcb.conflict = append(ttcb.conflict, sql.ConflictColumns(columns...))
	return &TweetTagUpsertBulk{
		create: ttcb,
	}
}

// TweetTagUpsertBulk is the builder for "upsert"-ing
// a bulk of TweetTag nodes.
type TweetTagUpsertBulk struct {
	create *TweetTagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tweettag.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TweetTagUpsertBulk) UpdateNewValues() *TweetTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tweettag.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TweetTag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TweetTagUpsertBulk) Ignore() *TweetTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TweetTagUpsertBulk) DoNothing() *TweetTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TweetTagCreateBulk.OnConflict
// documentation for more info.
func (u *TweetTagUpsertBulk) Update(set func(*TweetTagUpsert)) *TweetTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TweetTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddedAt sets the "added_at" field.
func (u *TweetTagUpsertBulk) SetAddedAt(v time.Time) *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetAddedAt(v)
	})
}

// UpdateAddedAt sets the "added_at" field to the value that was provided on create.
func (u *TweetTagUpsertBulk) UpdateAddedAt() *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateAddedAt()
	})
}

// SetTagID sets the "tag_id" field.
func (u *TweetTagUpsertBulk) SetTagID(v int) *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *TweetTagUpsertBulk) UpdateTagID() *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateTagID()
	})
}

// SetTweetID sets the "tweet_id" field.
func (u *TweetTagUpsertBulk) SetTweetID(v int) *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.SetTweetID(v)
	})
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *TweetTagUpsertBulk) UpdateTweetID() *TweetTagUpsertBulk {
	return u.Update(func(s *TweetTagUpsert) {
		s.UpdateTweetID()
	})
}

// Exec executes the query.
func (u *TweetTagUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if b == nil {
			return fmt.Errorf("ent: missing builder at index %d, unexpected nil builder passed to CreateBulk", i)
		}
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TweetTagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TweetTagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetTagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

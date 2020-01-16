// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/spec"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	create_time *time.Time
	update_time *time.Time
	number      *string
	name        *string
	owner       map[string]struct{}
	spec        map[string]struct{}
}

// SetCreateTime sets the create_time field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.create_time = &t
	return cc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the update_time field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.update_time = &t
	return cc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetNumber sets the number field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.number = &s
	return cc
}

// SetName sets the name field.
func (cc *CardCreate) SetName(s string) *CardCreate {
	cc.name = &s
	return cc
}

// SetNillableName sets the name field if the given value is not nil.
func (cc *CardCreate) SetNillableName(s *string) *CardCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CardCreate) SetOwnerID(id string) *CardCreate {
	if cc.owner == nil {
		cc.owner = make(map[string]struct{})
	}
	cc.owner[id] = struct{}{}
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(id *string) *CardCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cc *CardCreate) AddSpecIDs(ids ...string) *CardCreate {
	if cc.spec == nil {
		cc.spec = make(map[string]struct{})
	}
	for i := range ids {
		cc.spec[ids[i]] = struct{}{}
	}
	return cc
}

// AddSpec adds the spec edges to Spec.
func (cc *CardCreate) AddSpec(s ...*Spec) *CardCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSpecIDs(ids...)
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	if cc.create_time == nil {
		v := card.DefaultCreateTime()
		cc.create_time = &v
	}
	if cc.update_time == nil {
		v := card.DefaultUpdateTime()
		cc.update_time = &v
	}
	if cc.number == nil {
		return nil, errors.New("ent: missing required field \"number\"")
	}
	if err := card.NumberValidator(*cc.number); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"number\": %v", err)
	}
	if cc.name != nil {
		if err := card.NameValidator(*cc.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	var (
		c     = &Card{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: card.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: card.FieldID,
			},
		}
	)
	if value := cc.create_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: card.FieldCreateTime,
		})
		c.CreateTime = *value
	}
	if value := cc.update_time; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: card.FieldUpdateTime,
		})
		c.UpdateTime = *value
	}
	if value := cc.number; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: card.FieldNumber,
		})
		c.Number = *value
	}
	if value := cc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: card.FieldName,
		})
		c.Name = *value
	}
	if nodes := cc.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.spec; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: spec.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = strconv.FormatInt(id, 10)
	return c, nil
}

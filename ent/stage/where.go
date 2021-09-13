// Code generated by entc, DO NOT EDIT.

package stage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// TuitionAmount applies equality check predicate on the "tuition_amount" field. It's identical to TuitionAmountEQ.
func TuitionAmount(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTuitionAmount), v))
	})
}

// Directory applies equality check predicate on the "directory" field. It's identical to DirectoryEQ.
func Directory(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDirectory), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TuitionAmountEQ applies the EQ predicate on the "tuition_amount" field.
func TuitionAmountEQ(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTuitionAmount), v))
	})
}

// TuitionAmountNEQ applies the NEQ predicate on the "tuition_amount" field.
func TuitionAmountNEQ(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTuitionAmount), v))
	})
}

// TuitionAmountIn applies the In predicate on the "tuition_amount" field.
func TuitionAmountIn(vs ...int) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTuitionAmount), v...))
	})
}

// TuitionAmountNotIn applies the NotIn predicate on the "tuition_amount" field.
func TuitionAmountNotIn(vs ...int) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTuitionAmount), v...))
	})
}

// TuitionAmountGT applies the GT predicate on the "tuition_amount" field.
func TuitionAmountGT(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTuitionAmount), v))
	})
}

// TuitionAmountGTE applies the GTE predicate on the "tuition_amount" field.
func TuitionAmountGTE(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTuitionAmount), v))
	})
}

// TuitionAmountLT applies the LT predicate on the "tuition_amount" field.
func TuitionAmountLT(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTuitionAmount), v))
	})
}

// TuitionAmountLTE applies the LTE predicate on the "tuition_amount" field.
func TuitionAmountLTE(v int) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTuitionAmount), v))
	})
}

// DirectoryEQ applies the EQ predicate on the "directory" field.
func DirectoryEQ(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDirectory), v))
	})
}

// DirectoryNEQ applies the NEQ predicate on the "directory" field.
func DirectoryNEQ(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDirectory), v))
	})
}

// DirectoryIn applies the In predicate on the "directory" field.
func DirectoryIn(vs ...string) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDirectory), v...))
	})
}

// DirectoryNotIn applies the NotIn predicate on the "directory" field.
func DirectoryNotIn(vs ...string) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDirectory), v...))
	})
}

// DirectoryGT applies the GT predicate on the "directory" field.
func DirectoryGT(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDirectory), v))
	})
}

// DirectoryGTE applies the GTE predicate on the "directory" field.
func DirectoryGTE(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDirectory), v))
	})
}

// DirectoryLT applies the LT predicate on the "directory" field.
func DirectoryLT(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDirectory), v))
	})
}

// DirectoryLTE applies the LTE predicate on the "directory" field.
func DirectoryLTE(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDirectory), v))
	})
}

// DirectoryContains applies the Contains predicate on the "directory" field.
func DirectoryContains(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDirectory), v))
	})
}

// DirectoryHasPrefix applies the HasPrefix predicate on the "directory" field.
func DirectoryHasPrefix(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDirectory), v))
	})
}

// DirectoryHasSuffix applies the HasSuffix predicate on the "directory" field.
func DirectoryHasSuffix(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDirectory), v))
	})
}

// DirectoryEqualFold applies the EqualFold predicate on the "directory" field.
func DirectoryEqualFold(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDirectory), v))
	})
}

// DirectoryContainsFold applies the ContainsFold predicate on the "directory" field.
func DirectoryContainsFold(v string) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDirectory), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Stage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HasSchool applies the HasEdge predicate on the "school" edge.
func HasSchool() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchoolTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchoolTable, SchoolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchoolWith applies the HasEdge predicate on the "school" edge with a given conditions (other predicates).
func HasSchoolWith(preds ...predicate.School) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchoolInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchoolTable, SchoolColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClasses applies the HasEdge predicate on the "classes" edge.
func HasClasses() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassesWith applies the HasEdge predicate on the "classes" edge with a given conditions (other predicates).
func HasClassesWith(preds ...predicate.Class) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayments applies the HasEdge predicate on the "payments" edge.
func HasPayments() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentsTable, PaymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentsWith applies the HasEdge predicate on the "payments" edge with a given conditions (other predicates).
func HasPaymentsWith(preds ...predicate.TuitionPayment) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentsTable, PaymentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudents applies the HasEdge predicate on the "students" edge.
func HasStudents() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentsWith applies the HasEdge predicate on the "students" edge with a given conditions (other predicates).
func HasStudentsWith(preds ...predicate.User) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCourseGrades applies the HasEdge predicate on the "course_grades" edge.
func HasCourseGrades() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseGradesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseGradesTable, CourseGradesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseGradesWith applies the HasEdge predicate on the "course_grades" edge with a given conditions (other predicates).
func HasCourseGradesWith(preds ...predicate.CourseGrade) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseGradesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CourseGradesTable, CourseGradesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifications applies the HasEdge predicate on the "notifications" edge.
func HasNotifications() predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
func HasNotificationsWith(preds ...predicate.Notification) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stage) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stage) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stage) predicate.Stage {
	return predicate.Stage(func(s *sql.Selector) {
		p(s.Not())
	})
}

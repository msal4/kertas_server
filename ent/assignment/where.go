// Code generated by entc, DO NOT EDIT.

package assignment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/msal4/hassah_school_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// IsExam applies equality check predicate on the "is_exam" field. It's identical to IsExamEQ.
func IsExam(v bool) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsExam), v))
	})
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDueDate), v))
	})
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// IsExamEQ applies the EQ predicate on the "is_exam" field.
func IsExamEQ(v bool) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsExam), v))
	})
}

// IsExamNEQ applies the NEQ predicate on the "is_exam" field.
func IsExamNEQ(v bool) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsExam), v))
	})
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDueDate), v))
	})
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDueDate), v))
	})
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDueDate), v...))
	})
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDueDate), v...))
	})
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDueDate), v))
	})
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDueDate), v))
	})
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDueDate), v))
	})
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDueDate), v))
	})
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDuration), v))
	})
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDuration), v...))
	})
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDuration), v...))
	})
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDuration), v))
	})
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDuration), v))
	})
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDuration), v))
	})
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDuration), v))
	})
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDuration)))
	})
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDuration)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...time.Time) predicate.Assignment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Assignment(func(s *sql.Selector) {
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
func DeletedAtGT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubmissions applies the HasEdge predicate on the "submissions" edge.
func HasSubmissions() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubmissionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmissionsTable, SubmissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionsWith applies the HasEdge predicate on the "submissions" edge with a given conditions (other predicates).
func HasSubmissionsWith(preds ...predicate.AssignmentSubmission) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubmissionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmissionsTable, SubmissionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGrades applies the HasEdge predicate on the "grades" edge.
func HasGrades() predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GradesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GradesTable, GradesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGradesWith applies the HasEdge predicate on the "grades" edge with a given conditions (other predicates).
func HasGradesWith(preds ...predicate.Grade) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GradesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GradesTable, GradesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Assignment) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Assignment) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
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
func Not(p predicate.Assignment) predicate.Assignment {
	return predicate.Assignment(func(s *sql.Selector) {
		p(s.Not())
	})
}

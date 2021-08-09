// Code generated by entc, DO NOT EDIT.

package tuitionpayment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// PaidAmount applies equality check predicate on the "paid_amount" field. It's identical to PaidAmountEQ.
func PaidAmount(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaidAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// PaidAmountEQ applies the EQ predicate on the "paid_amount" field.
func PaidAmountEQ(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaidAmount), v))
	})
}

// PaidAmountNEQ applies the NEQ predicate on the "paid_amount" field.
func PaidAmountNEQ(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaidAmount), v))
	})
}

// PaidAmountIn applies the In predicate on the "paid_amount" field.
func PaidAmountIn(vs ...int) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaidAmount), v...))
	})
}

// PaidAmountNotIn applies the NotIn predicate on the "paid_amount" field.
func PaidAmountNotIn(vs ...int) predicate.TuitionPayment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TuitionPayment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaidAmount), v...))
	})
}

// PaidAmountGT applies the GT predicate on the "paid_amount" field.
func PaidAmountGT(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaidAmount), v))
	})
}

// PaidAmountGTE applies the GTE predicate on the "paid_amount" field.
func PaidAmountGTE(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaidAmount), v))
	})
}

// PaidAmountLT applies the LT predicate on the "paid_amount" field.
func PaidAmountLT(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaidAmount), v))
	})
}

// PaidAmountLTE applies the LTE predicate on the "paid_amount" field.
func PaidAmountLTE(v int) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaidAmount), v))
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.User) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStage applies the HasEdge predicate on the "stage" edge.
func HasStage() predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StageTable, StageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStageWith applies the HasEdge predicate on the "stage" edge with a given conditions (other predicates).
func HasStageWith(preds ...predicate.Stage) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StageTable, StageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TuitionPayment) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TuitionPayment) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
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
func Not(p predicate.TuitionPayment) predicate.TuitionPayment {
	return predicate.TuitionPayment(func(s *sql.Selector) {
		p(s.Not())
	})
}

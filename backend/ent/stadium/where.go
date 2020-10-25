// Code generated by entc, DO NOT EDIT.

package stadium

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/tew147258/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Namestadium applies equality check predicate on the "namestadium" field. It's identical to NamestadiumEQ.
func Namestadium(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamestadium), v))
	})
}

// NamestadiumEQ applies the EQ predicate on the "namestadium" field.
func NamestadiumEQ(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNamestadium), v))
	})
}

// NamestadiumNEQ applies the NEQ predicate on the "namestadium" field.
func NamestadiumNEQ(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNamestadium), v))
	})
}

// NamestadiumIn applies the In predicate on the "namestadium" field.
func NamestadiumIn(vs ...string) predicate.Stadium {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stadium(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNamestadium), v...))
	})
}

// NamestadiumNotIn applies the NotIn predicate on the "namestadium" field.
func NamestadiumNotIn(vs ...string) predicate.Stadium {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stadium(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNamestadium), v...))
	})
}

// NamestadiumGT applies the GT predicate on the "namestadium" field.
func NamestadiumGT(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNamestadium), v))
	})
}

// NamestadiumGTE applies the GTE predicate on the "namestadium" field.
func NamestadiumGTE(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNamestadium), v))
	})
}

// NamestadiumLT applies the LT predicate on the "namestadium" field.
func NamestadiumLT(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNamestadium), v))
	})
}

// NamestadiumLTE applies the LTE predicate on the "namestadium" field.
func NamestadiumLTE(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNamestadium), v))
	})
}

// NamestadiumContains applies the Contains predicate on the "namestadium" field.
func NamestadiumContains(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNamestadium), v))
	})
}

// NamestadiumHasPrefix applies the HasPrefix predicate on the "namestadium" field.
func NamestadiumHasPrefix(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNamestadium), v))
	})
}

// NamestadiumHasSuffix applies the HasSuffix predicate on the "namestadium" field.
func NamestadiumHasSuffix(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNamestadium), v))
	})
}

// NamestadiumEqualFold applies the EqualFold predicate on the "namestadium" field.
func NamestadiumEqualFold(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNamestadium), v))
	})
}

// NamestadiumContainsFold applies the ContainsFold predicate on the "namestadium" field.
func NamestadiumContainsFold(v string) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNamestadium), v))
	})
}

// HasStadiumConfirmation applies the HasEdge predicate on the "StadiumConfirmation" edge.
func HasStadiumConfirmation() predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StadiumConfirmationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StadiumConfirmationTable, StadiumConfirmationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStadiumConfirmationWith applies the HasEdge predicate on the "StadiumConfirmation" edge with a given conditions (other predicates).
func HasStadiumConfirmationWith(preds ...predicate.Confirmation) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StadiumConfirmationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StadiumConfirmationTable, StadiumConfirmationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
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
func Not(p predicate.Stadium) predicate.Stadium {
	return predicate.Stadium(func(s *sql.Selector) {
		p(s.Not())
	})
}
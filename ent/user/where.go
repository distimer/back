// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"pentag.kr/distimer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// OauthID applies equality check predicate on the "oauth_id" field. It's identical to OauthIDEQ.
func OauthID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOauthID, v))
}

// OauthProvider applies equality check predicate on the "oauth_provider" field. It's identical to OauthProviderEQ.
func OauthProvider(v int8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOauthProvider, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// OauthIDEQ applies the EQ predicate on the "oauth_id" field.
func OauthIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOauthID, v))
}

// OauthIDNEQ applies the NEQ predicate on the "oauth_id" field.
func OauthIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOauthID, v))
}

// OauthIDIn applies the In predicate on the "oauth_id" field.
func OauthIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldOauthID, vs...))
}

// OauthIDNotIn applies the NotIn predicate on the "oauth_id" field.
func OauthIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOauthID, vs...))
}

// OauthIDGT applies the GT predicate on the "oauth_id" field.
func OauthIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldOauthID, v))
}

// OauthIDGTE applies the GTE predicate on the "oauth_id" field.
func OauthIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOauthID, v))
}

// OauthIDLT applies the LT predicate on the "oauth_id" field.
func OauthIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldOauthID, v))
}

// OauthIDLTE applies the LTE predicate on the "oauth_id" field.
func OauthIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOauthID, v))
}

// OauthIDContains applies the Contains predicate on the "oauth_id" field.
func OauthIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldOauthID, v))
}

// OauthIDHasPrefix applies the HasPrefix predicate on the "oauth_id" field.
func OauthIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldOauthID, v))
}

// OauthIDHasSuffix applies the HasSuffix predicate on the "oauth_id" field.
func OauthIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldOauthID, v))
}

// OauthIDEqualFold applies the EqualFold predicate on the "oauth_id" field.
func OauthIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldOauthID, v))
}

// OauthIDContainsFold applies the ContainsFold predicate on the "oauth_id" field.
func OauthIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldOauthID, v))
}

// OauthProviderEQ applies the EQ predicate on the "oauth_provider" field.
func OauthProviderEQ(v int8) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOauthProvider, v))
}

// OauthProviderNEQ applies the NEQ predicate on the "oauth_provider" field.
func OauthProviderNEQ(v int8) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOauthProvider, v))
}

// OauthProviderIn applies the In predicate on the "oauth_provider" field.
func OauthProviderIn(vs ...int8) predicate.User {
	return predicate.User(sql.FieldIn(FieldOauthProvider, vs...))
}

// OauthProviderNotIn applies the NotIn predicate on the "oauth_provider" field.
func OauthProviderNotIn(vs ...int8) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOauthProvider, vs...))
}

// OauthProviderGT applies the GT predicate on the "oauth_provider" field.
func OauthProviderGT(v int8) predicate.User {
	return predicate.User(sql.FieldGT(FieldOauthProvider, v))
}

// OauthProviderGTE applies the GTE predicate on the "oauth_provider" field.
func OauthProviderGTE(v int8) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOauthProvider, v))
}

// OauthProviderLT applies the LT predicate on the "oauth_provider" field.
func OauthProviderLT(v int8) predicate.User {
	return predicate.User(sql.FieldLT(FieldOauthProvider, v))
}

// OauthProviderLTE applies the LTE predicate on the "oauth_provider" field.
func OauthProviderLTE(v int8) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOauthProvider, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// HasJoinedGroups applies the HasEdge predicate on the "joined_groups" edge.
func HasJoinedGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, JoinedGroupsTable, JoinedGroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJoinedGroupsWith applies the HasEdge predicate on the "joined_groups" edge with a given conditions (other predicates).
func HasJoinedGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newJoinedGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnedGroups applies the HasEdge predicate on the "owned_groups" edge.
func HasOwnedGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OwnedGroupsTable, OwnedGroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnedGroupsWith applies the HasEdge predicate on the "owned_groups" edge with a given conditions (other predicates).
func HasOwnedGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOwnedGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudyLogs applies the HasEdge predicate on the "study_logs" edge.
func HasStudyLogs() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StudyLogsTable, StudyLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudyLogsWith applies the HasEdge predicate on the "study_logs" edge with a given conditions (other predicates).
func HasStudyLogsWith(preds ...predicate.StudyLog) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newStudyLogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTimers applies the HasEdge predicate on the "timers" edge.
func HasTimers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TimersTable, TimersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTimersWith applies the HasEdge predicate on the "timers" edge with a given conditions (other predicates).
func HasTimersWith(preds ...predicate.Timer) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTimersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRefreshTokens applies the HasEdge predicate on the "refresh_tokens" edge.
func HasRefreshTokens() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RefreshTokensTable, RefreshTokensColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRefreshTokensWith applies the HasEdge predicate on the "refresh_tokens" edge with a given conditions (other predicates).
func HasRefreshTokensWith(preds ...predicate.RefreshToken) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRefreshTokensStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnedCategories applies the HasEdge predicate on the "owned_categories" edge.
func HasOwnedCategories() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OwnedCategoriesTable, OwnedCategoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnedCategoriesWith applies the HasEdge predicate on the "owned_categories" edge with a given conditions (other predicates).
func HasOwnedCategoriesWith(preds ...predicate.Category) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOwnedCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAffiliations applies the HasEdge predicate on the "affiliations" edge.
func HasAffiliations() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AffiliationsTable, AffiliationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAffiliationsWith applies the HasEdge predicate on the "affiliations" edge with a given conditions (other predicates).
func HasAffiliationsWith(preds ...predicate.Affiliation) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAffiliationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}

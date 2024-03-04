// Code generated by ent, DO NOT EDIT.

package apitoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldCreatedAt, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldExpiresAt, v))
}

// RevokedAt applies equality check predicate on the "revoked_at" field. It's identical to RevokedAtEQ.
func RevokedAt(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldRevokedAt, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldOrganizationID, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.APIToken {
	return predicate.APIToken(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLTE(FieldCreatedAt, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLTE(FieldExpiresAt, v))
}

// ExpiresAtIsNil applies the IsNil predicate on the "expires_at" field.
func ExpiresAtIsNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldIsNull(FieldExpiresAt))
}

// ExpiresAtNotNil applies the NotNil predicate on the "expires_at" field.
func ExpiresAtNotNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldNotNull(FieldExpiresAt))
}

// RevokedAtEQ applies the EQ predicate on the "revoked_at" field.
func RevokedAtEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldRevokedAt, v))
}

// RevokedAtNEQ applies the NEQ predicate on the "revoked_at" field.
func RevokedAtNEQ(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldRevokedAt, v))
}

// RevokedAtIn applies the In predicate on the "revoked_at" field.
func RevokedAtIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldRevokedAt, vs...))
}

// RevokedAtNotIn applies the NotIn predicate on the "revoked_at" field.
func RevokedAtNotIn(vs ...time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldRevokedAt, vs...))
}

// RevokedAtGT applies the GT predicate on the "revoked_at" field.
func RevokedAtGT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGT(FieldRevokedAt, v))
}

// RevokedAtGTE applies the GTE predicate on the "revoked_at" field.
func RevokedAtGTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldGTE(FieldRevokedAt, v))
}

// RevokedAtLT applies the LT predicate on the "revoked_at" field.
func RevokedAtLT(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLT(FieldRevokedAt, v))
}

// RevokedAtLTE applies the LTE predicate on the "revoked_at" field.
func RevokedAtLTE(v time.Time) predicate.APIToken {
	return predicate.APIToken(sql.FieldLTE(FieldRevokedAt, v))
}

// RevokedAtIsNil applies the IsNil predicate on the "revoked_at" field.
func RevokedAtIsNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldIsNull(FieldRevokedAt))
}

// RevokedAtNotNil applies the NotNil predicate on the "revoked_at" field.
func RevokedAtNotNil() predicate.APIToken {
	return predicate.APIToken(sql.FieldNotNull(FieldRevokedAt))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.APIToken {
	return predicate.APIToken(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.APIToken {
	return predicate.APIToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.APIToken {
	return predicate.APIToken(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.APIToken) predicate.APIToken {
	return predicate.APIToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.APIToken) predicate.APIToken {
	return predicate.APIToken(func(s *sql.Selector) {
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
func Not(p predicate.APIToken) predicate.APIToken {
	return predicate.APIToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
// lint_ext_cert_policy_disallowed_any_policy_qualifier.go
/*******************************************************************
RFC 5280: 4.2.1.4
To promote interoperability, this profile RECOMMENDS that policy
information terms consist of only an OID.  Where an OID alone is
insufficient, this profile strongly recommends that the use of
qualifiers be limited to those identified in this section.  When
qualifiers are used with the special policy anyPolicy, they MUST be
limited to the qualifiers identified in this section.  Only those
qualifiers returned as a result of path validation are considered.
********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type unrecommendedQualifier struct{}

func (l *unrecommendedQualifier) Initialize() error {
	return nil
}

func (l *unrecommendedQualifier) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CertPolicyOID)
}

func (l *unrecommendedQualifier) Execute(c *x509.Certificate) *LintResult {
	for _, firstLvl := range c.QualifierId {
		for _, qualifierId := range firstLvl {
			if !qualifierId.Equal(util.CpsOID) && !qualifierId.Equal(util.UserNoticeOID) {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_cert_policy_disallowed_any_policy_qualifier",
		Description:   "When qualifiers are used with the special policy anyPolicy, they must be limited to qualifiers identified in this section: (4.2.1.4)",
		Citation:      "RFC 5280: 4.2.1.4",
		Source:        RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &unrecommendedQualifier{},
	})
}

// lint_ext_name_constraints_not_in_ca.go
/***********************************************************************
RFC 5280: 4.2.1.10
The name constraints extension, which MUST be used only in a CA
   certificate, indicates a name space within which all subject names in
   subsequent certificates in a certification path MUST be located.
   Restrictions apply to the subject distinguished name and apply to
   subject alternative names.  Restrictions apply only when the
   specified name form is present.  If no name of the type is in the
   certificate, the certificate is acceptable.
***********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintNotCa struct{}

func (l *nameConstraintNotCa) Initialize() error {
	return nil
}

func (l *nameConstraintNotCa) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintNotCa) Execute(c *x509.Certificate) *LintResult {
	if !util.IsCACert(c) {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_name_constraints_not_in_ca",
		Description:   "The name constraints extension MUST only be used in CA certificates",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &nameConstraintNotCa{},
	})
}

// lint_old_sub_ca_rsa_mod_less_than_1024_bits.go
// CHANGE THIS COMMENT TO MATCH SOURCE TEXT

package lints

import (
	"crypto/rsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaModSize struct{}

func (l *subCaModSize) Initialize() error {
	return nil
}

func (l *subCaModSize) CheckApplies(c *x509.Certificate) bool {
	issueDate := c.NotBefore
	endDate := c.NotAfter
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && util.IsSubCA(c) && issueDate.Before(util.NoRSA1024RootDate) && endDate.Before(util.NoRSA1024Date)
}

func (l *subCaModSize) Execute(c *x509.Certificate) *LintResult {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.N.BitLen() < 1024 {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:        "e_old_sub_ca_rsa_mod_less_than_1024_bits",
		Description: "In a validity period beginning on or before 31 Dec 2010 and ending on or before 31 Dec 2013, subordinate CA certificates using RSA public key algorithm MUST use a 1024 bit modulus",
		Citation:    "BRs: 6.1.5",
		Source:      CABFBaselineRequirements,
		// since effective date should be checked against end date in this specific case, putting time check into checkApplies instead, ZeroDate here to automatically pass NE test
		EffectiveDate: util.ZeroDate,
		Lint:          &subCaModSize{},
	})
}

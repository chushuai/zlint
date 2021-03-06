package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"regexp"
	"strings"
)

type DNSNameProperCharacters struct {
	CompiledExpression *regexp.Regexp
}

func (l *DNSNameProperCharacters) Initialize() error {
	const dnsLabelRegex = "^[A-Za-z0-9*_-]+$"
	l.CompiledExpression = regexp.MustCompile(dnsLabelRegex)
	return nil
}

func (l *DNSNameProperCharacters) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func (l *DNSNameProperCharacters) labelContainsBadCharacters(domain string) bool {
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if !l.CompiledExpression.MatchString(label) {
			return true
		}
	}
	return false
}

func (l *DNSNameProperCharacters) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" {
		commonNameWithoutWildcard := util.RemovePrependedWildcard(c.Subject.CommonName)
		commonNameWithoutQuestionMarks := util.RemovePrependedQuestionMarks(commonNameWithoutWildcard)
		badCharacterFound := l.labelContainsBadCharacters(commonNameWithoutQuestionMarks)
		if badCharacterFound {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		domainWithoutWildcard := util.RemovePrependedWildcard(dns)
		domainWithoutQuestionMarks := util.RemovePrependedQuestionMarks(domainWithoutWildcard)
		badCharacterFound := l.labelContainsBadCharacters(domainWithoutQuestionMarks)
		if badCharacterFound {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_bad_character_in_label",
		Description:   "Characters in labels of DNSNames MUST be alphanumeric, - , _ or *",
		Citation:      "BRs: 7.1.4.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameProperCharacters{},
	})
}

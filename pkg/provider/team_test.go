package provider

import (
	"testing"
)

func TestValidateAuthEntry(t *testing.T) {
	valid := []string{
		"user:local:admin",
		"user:github:tlwr",
		"user:saml:platform-owner@example.com",
		"group:saml:platform-owners",
		"group:saml:concourse-admins",
		"group:github:alphagov:paas-team",
		"group:ldap:cn=devs,ou=groups,dc=example,dc=com",
	}
	for _, v := range valid {
		_, errs := validateAuthEntry(v, "owners")
		if len(errs) != 0 {
			t.Errorf("expected %q to be valid, got errors: %v", v, errs)
		}
	}

	invalid := []string{
		"saml:platform-owners",     // missing user:/group: discriminator — was silently dropped before this fix
		"platform-owners",          // bare group name, no prefix
		"",                         // empty string
		"local:admin",              // missing discriminator
		"github:org:team",          // missing discriminator
	}
	for _, v := range invalid {
		_, errs := validateAuthEntry(v, "owners")
		if len(errs) == 0 {
			t.Errorf("expected %q to be invalid, but got no errors", v)
		}
	}
}

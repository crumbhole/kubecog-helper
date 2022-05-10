package schema

import (
	"testing"
)

func TestValid(t *testing.T) {
	vals := []CogValues{{Platform: Platform{Provider: "rke"}},
		{ArgoCD: ArgoCD{HA: true},
			Platform: Platform{Provider: "aks"}},
		{ArgoCD: ArgoCD{HA: false},
			Platform: Platform{Provider: "aks"}},
	}
	v := KubecogValidator{}
	for i, val := range vals {
		t.Logf("Valid test %d\n", i)
		errorString := v.ValidateToSingleString(val)
		if errorString != "" {
			t.Error(errorString)
		}
	}
}

type expectedErrors []string

func ensureAllErrors(t *testing.T, received []string, expected expectedErrors) {
	for _, rec := range received {
		found := false
		for _, exp := range expected {
			if rec == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpectedly got error <%v>\n", rec)
		}
	}
	for _, exp := range expected {
		found := false
		for _, rec := range received {
			if rec == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Did not receive error <%v>\n", exp)
		}
	}
}

func TestEmpty(t *testing.T) {
	expectedErrors := expectedErrors{`CogValues.Platform.Provider: Provider is a required field`}
	v := KubecogValidator{}
	received := v.ValidateToStrings(CogValues{})
	ensureAllErrors(t, received, expectedErrors)
}

type failTest struct {
	values CogValues
	expect expectedErrors
}

func TestFailing(t *testing.T) {
	tests := []failTest{{CogValues{Platform: Platform{Provider: "poo"}},
		expectedErrors{`CogValues.Platform.Provider: Provider must be one of [rke k3s aks eks]`},
	}}
	v := KubecogValidator{}
	for i, val := range tests {
		t.Logf("Failing test %d\n", i)
		received := v.ValidateToStrings(val.values)
		ensureAllErrors(t, received, val.expect)
	}
}

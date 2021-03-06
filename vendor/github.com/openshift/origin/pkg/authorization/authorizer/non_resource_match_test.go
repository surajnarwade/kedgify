package authorizer

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
	kauthorizer "k8s.io/apiserver/pkg/authorization/authorizer"

	authorizationapi "github.com/openshift/origin/pkg/authorization/apis/authorization"
)

type nonResourceMatchTest struct {
	url            string
	matcher        string
	expectedResult bool
}

func TestNonResourceMatchStar(t *testing.T) {
	test := &nonResourceMatchTest{
		url:            "first/second",
		matcher:        "first/*",
		expectedResult: true,
	}
	test.run(t)
}

func TestNonResourceMatchExact(t *testing.T) {
	test := &nonResourceMatchTest{
		url:            "first/second",
		matcher:        "first/second",
		expectedResult: true,
	}
	test.run(t)
}

func TestNonResourceMatchMatcherEndsShort(t *testing.T) {
	test := &nonResourceMatchTest{
		url:            "first/second",
		matcher:        "first",
		expectedResult: false,
	}
	test.run(t)
}

func TestNonResourceMatchURLEndsShort(t *testing.T) {
	test := &nonResourceMatchTest{
		url:            "first",
		matcher:        "first/second",
		expectedResult: false,
	}
	test.run(t)
}

func TestNonResourceMatchNoSimilarity(t *testing.T) {
	test := &nonResourceMatchTest{
		url:            "first/second",
		matcher:        "foo",
		expectedResult: false,
	}
	test.run(t)
}

func (test *nonResourceMatchTest) run(t *testing.T) {
	attributes := kauthorizer.AttributesRecord{
		ResourceRequest: false,
		Path:            test.url,
	}

	rule := authorizationapi.PolicyRule{NonResourceURLs: sets.NewString(test.matcher)}

	result := nonResourceMatches(attributes, rule)

	if result != test.expectedResult {
		t.Errorf("Expected %v, got %v", test.expectedResult, result)
	}

}

package gomega_matchers

import "github.com/onsi/gomega/types"

func MatchXML(xml interface{}) types.GomegaMatcher {
	return &MatchXMLMatcher{
		XMLToMatch: xml,
	}
}

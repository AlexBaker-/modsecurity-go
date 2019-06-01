package main

import (
	"github.com/sirupsen/logrus"

	"github.com/senghoo/modsecurity-go/libmodsecurity"
)

const (
	rulesFile = "basic_rules.conf"
)

func main() {
	lms := libmodsecurity.NewLibModSecurity()
	if err := lms.AddRuleFromFile(rulesFile); err != nil {
		logrus.Errorf("Failed to load %s: %v", rulesFile, err)
	}
	tr := lms.NewTransaction()
	tr.ProcessConnection("127.0.0.1", "127.0.0.1", 12345, 80)
	tr.ProcessURL(
		"http://www.modsecurity.org/test?key1=value1&key2=value2&key3=value3",
		"GET",
		1,
		1)
	tr.ProcessRequestHeader()
	tr.ProcessRequestBody()
	tr.ProcessResponseHeader()
	tr.ProcessResponseBody()

}

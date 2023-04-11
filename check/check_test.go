package check

import (
	"fmt"
	"testing"
)

func TestDoAllCheck(t *testing.T) {
	checker := SupportedChecker{}
	check := checker.DoAllCheck("dl.google.com")
	fmt.Println(check)

}

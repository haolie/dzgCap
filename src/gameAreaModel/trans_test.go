package gameAreaModel

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrans(t *testing.T) {
	// TransFile("mini")

	str:="sldfjasdf{0}sldfjasdf"
	str=strings.Replace(str,"{0}","%d",1)
	fmt.Printf(str,3232)
}

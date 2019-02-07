package util

import (
	"fmt"

	"github.com/edge-go/core"
)

func MakeHttpUrl(sd core.ServiceDef, sp core.ServicePath) string {
	return fmt.Sprintf("%s://%s:%d/%s", sd.Proto, sd.Host, sd.Port, sp.Path)
}

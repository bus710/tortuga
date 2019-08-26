package tortuga

import (
	"testing"

	tortuga "github.com/bus710/tortuga"
)

func TestInit(t *testing.T) {
	tConn := tortuga.Connection{}
	tConn.Init()
	Dummy()

}

package stopserver

import (
	"context"
	. "github.com/comodo/comodoca-status-api/startserver"
)

func StopStatusServer(ctx context.Context) {
	StatusServer.Shutdown(ctx)
}

package stopserver

import (
	"context"
	. "github.com/comodo/comodoca-status-api/startserver"
	"time"
)

func StopStatusServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	StatusServer.Shutdown(ctx)
	defer cancel()
}

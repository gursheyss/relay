package relay

import (
	"fmt"

	"github.com/gursheyss/relay/aws"
	"github.com/gursheyss/relay/server"
)

func Init() (startFunc func() error, err error) {
	if err := aws.InitAWS(); err != nil {
		return nil, fmt.Errorf("failed to initialize AWS: %w", err)
	}

	startFunc, err = server.SetupServer()
	if err != nil {
		return nil, fmt.Errorf("failed to start HTTP server: %w", err)
	}

	return startFunc, nil
}

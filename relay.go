package relay

import (
	"fmt"

	"github.com/gursheyss/relay/aws"
	"github.com/gursheyss/relay/options"
	"github.com/gursheyss/relay/server"
)

func Init(options ...options.Option) (startFunc func() error, err error) {
	if err := aws.InitAWS(); err != nil {
		return nil, fmt.Errorf("failed to initialize AWS: %w", err)
	}

	startFunc, err = server.SetupServer(options...)
	if err != nil {
		return nil, fmt.Errorf("failed to start HTTP server: %w", err)
	}

	return startFunc, nil
}

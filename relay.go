package relay

import (
	"fmt"

	"github.com/gursheyss/relay/aws"
	"github.com/gursheyss/relay/server"
)

func Init() error {
	if err := aws.InitAWS(); err != nil {
		return fmt.Errorf("failed to initialize AWS: %w", err)
	}

	if err := server.StartHTTPServer(); err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}

	return nil
}

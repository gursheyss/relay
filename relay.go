package relay

import (
	"github.com/gursheyss/relay/aws"
	"github.com/gursheyss/relay/server"
)

func Init() error {
	if err := aws.InitAWS(); err != nil {
		return err
	}

	if err := server.StartHTTPServer(); err != nil {
		return err
	}

	return nil
}

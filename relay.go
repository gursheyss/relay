package relay

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	_, err = credentials.NewEnvCredentials().Get()
	if err != nil {
		return err
	}

	svc := sts.New(sess)

	_, err = svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}

	return nil
}

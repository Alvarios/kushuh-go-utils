package extLogs

import (
	"fmt"
	cbErrors "github.com/Alvarios/kushuh-go-utils/ext-apis/couchbase/cb-errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Params struct {
	Context *gin.Context
	Config *ServerConfig
	Error error
	StringError string
	Reason string
	File string
}

func UnexpectedCouchbaseAbort(p Params) {
	serr := p.StringError
	if serr == "" {
		serr = p.Error.Error()
	}
	prettyError, err := cbErrors.Pretty(fmt.Errorf(serr))
	if err != nil {
		prettyError = err.Error()
	}

	UnexpectedAbort(Params{
		StringError: prettyError,
		Context: p.Context,
		Config: p.Config,
		Reason: p.Reason,
		File: p.File,
	})
}

func UnexpectedAbort(p Params) {
	serr := p.StringError
	if serr == "" {
		serr = p.Error.Error()
	}

	p.Config.Error(
		p.File,
		fmt.Sprintf(
			"*%s* -> <http://%s%s>\n\n```%s```\n\n",
			p.Context.Request.Method,
			p.Context.Request.Host,
			p.Context.Request.URL.Path,
			serr,
		),
	)

	// Abort request.
	p.Context.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"message" : serr},
	)
}

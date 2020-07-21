package cbRequest

import (
	"fmt"
	extLogs "github.com/Alvarios/kushuh-go-utils/ext-logs"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFile(c *gocb.Collection, k string, o interface{}, p extLogs.Params, permissive bool) error {
	post, err := c.Get(k, nil)

	if err != nil && !permissive {
		if err == gocb.ErrDocumentNotFound || err == gocb.ErrPathNotFound {
			p.Context.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"message" : fmt.Sprintf("unable to find document with id %s", k)},
			)

			return nil
		}

		p.Reason = "unable to fetch document with id " + k
		p.Error = err
		extLogs.UnexpectedCouchbaseAbort(p)

		return nil
	} else if err != nil {
		return err
	}

	err = post.Content(o)

	// Handle error in case request has an unexpected structure.
	if err != nil && !permissive {
		// Cannot expect to have wrong structure. This is a critical error.
		p.Reason = "unexpected document structure for id " + k
		p.Error = err
		extLogs.UnexpectedCouchbaseAbort(p)
	}

	return err
}

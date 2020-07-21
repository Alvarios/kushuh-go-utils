package cbRequest

import (
	"fmt"
	extLogs "github.com/Alvarios/kushuh-go-utils/ext-logs"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CbFile struct {
	Collection *gocb.Collection
	ID string
	OutputPtr interface{}
	ErrorHandler func(c *CbFile, err error) error
	FatalHandler func(c *CbFile, err error) error
}

func DefaultErrorHandler (log extLogs.Params) func (c *CbFile, err error) error {
	return func (c *CbFile, err error) error {
		log.Context.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message" : err.Error()},
		)

		return err
	}
}

func DefaultFatalHandler (log extLogs.Params) func (c *CbFile, err error) error {
	return func (c *CbFile, err error) error {
		log.Error = err
		extLogs.UnexpectedCouchbaseAbort(log)
		return err
	}
}

func (f *CbFile) Fetch() error {
	post, err := f.Collection.Get(f.ID, nil)

	if err != nil && (err == gocb.ErrDocumentNotFound || err == gocb.ErrPathNotFound) {
		err = fmt.Errorf("unable to find document with id %s : %s", f.ID, err.Error())

		if f.ErrorHandler != nil {
			nerr := f.ErrorHandler(f, err)
			return nerr
		}

		return err
	}

	if err != nil {
		err = fmt.Errorf("unable to fetch document with id %s : %s", f.ID, err.Error())

		if f.FatalHandler != nil {
			nerr := f.FatalHandler(f, err)
			return nerr
		}

		return err
	}

	err = post.Content(f.OutputPtr)

	if err != nil && f.FatalHandler != nil {
		err = fmt.Errorf("\"unexpected document structure for id %s : %s", f.ID, err.Error())

		nerr := f.FatalHandler(f, err)
		return nerr
	}

	return err
}


package server

import (
	"github.com/digital-dream-labs/api/go/jdocspb"
	"github.com/digital-dream-labs/hugh/log"
)

const JDOCS_PATH = "/home/pi/jdocs/"

type options struct {
	log log.Logger
}

// Option is the list of options
type Option func(*options)

type Server struct {
	jdocspb.UnimplementedJdocsServer
}

// New accepts a list of args and returns the service
func New(opts ...Option) (*Server, error) {
	cfg := options{
		log: log.Base(),
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	s := Server{}

	return &s, nil
}

func getDocName(fname string, uid string, thing string) string {
	return getDocFolder(uid, thing) + "/" + fname
}

func getDocFolder(uid string, thing string) string {
	return JDOCS_PATH + uid + "/" + thing
}

func getUserFolder(uid string) string {
	return JDOCS_PATH + uid
}

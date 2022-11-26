package jdocs

import (
	"context"
	"github.com/digital-dream-labs/api/go/jdocspb"
	"github.com/digital-dream-labs/hugh/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
)

const JDOCS_PATH = "/home/pi/jdocs/"

type options struct {
	log log.Logger
}

// Option is the list of options
type Option func(*options)

type JDocsServer struct {
	jdocspb.UnimplementedJdocsServer
}

// New accepts a list of args and returns the service
func New(opts ...Option) (*JDocsServer, error) {
	cfg := options{
		log: log.Base(),
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	s := JDocsServer{}

	println("JDocs server created")

	return &s, nil
}

func (*JDocsServer) DeleteDoc(ctx context.Context, req *jdocspb.DeleteDocReq) (*jdocspb.DeleteDocResp, error) {
	fname := req.GetDocName()

	println("[DeleteDoc] NAME: " + fname + ", USERID: " + req.UserId + ", THING: " + req.Thing)

	_ = os.Remove(getDocName(fname, req.UserId, req.Thing))

	response := jdocspb.DeleteDocResp{}

	return &response, nil
}

func (*JDocsServer) PurgeAccountDocs(ctx context.Context, req *jdocspb.PurgeAccountDocsReq) (*jdocspb.PurgeAccountDocsResp, error) {
	println("[PurgeAccountDocs] REQ: " + req.String())
	return nil, status.Errorf(codes.Unimplemented, "method PurgeAccountDocs not implemented")
}

func (*JDocsServer) ReadDocs(ctx context.Context, req *jdocspb.ReadDocsReq) (*jdocspb.ReadDocsResp, error) {
	println("[ReadDocs] REQ: " + req.String())
	return nil, status.Errorf(codes.Unimplemented, "method ReadDocs not implemented")
}

func (*JDocsServer) WriteDoc(ctx context.Context, req *jdocspb.WriteDocReq) (*jdocspb.WriteDocResp, error) {
	fname := req.GetDocName()
	uid := req.UserId
	thing := req.Thing
	version := req.GetDoc().GetDocVersion()
	fileContent := req.GetDoc().JsonDoc

	println("[WriteDoc] NAME: " + fname + ", UID: " + uid + ", THING:" + thing + ", CONTENT: " + fileContent)

	folder := getDocFolder(uid, thing)
	fc := []byte(fileContent)
	os.MkdirAll(folder, os.ModePerm)
	_ = os.WriteFile(getDocName(fname, uid, thing), fc, 0644)

	response := jdocspb.WriteDocResp{
		Status:           jdocspb.WriteDocResp_ACCEPTED,
		LatestDocVersion: version,
	}

	return &response, nil
}

func (*JDocsServer) ViewAccountDocs(ctx context.Context, req *jdocspb.ViewAccountDocsReq) (*jdocspb.ViewDocsResp, error) {
	uid := req.UserId

	println("[ViewAccountDocs] USERID: " + uid + ", REQ: " + req.String())

	docs := make([]*jdocspb.ViewDoc, 0)

	folder := getUserFolder(uid)

	var things []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			things = append(things, info.Name())
		}
		return nil
	})
	if err == nil {
		for i := 0; i < len(things); i++ {
			err = filepath.Walk(folder+"/"+things[i], func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					content, readErr := os.ReadFile(path)
					if readErr == nil {
						doc := jdocspb.ViewDoc{
							UserId:  uid,
							DocName: info.Name(),
							Thing:   things[i],
							JsonDoc: string(content),
						}
						docs = append(docs, &doc)
					}
				}
				return nil
			})
		}
	}

	response := jdocspb.ViewDocsResp{
		Docs: docs,
	}

	return &response, nil
}

func (*JDocsServer) ViewAccountDocsWithPII(ctx context.Context, req *jdocspb.ViewAccountDocsReq) (*jdocspb.ViewDocsResp, error) {
	println("[ViewAccountDocs] REQ: " + req.String())

	return nil, status.Errorf(codes.Unimplemented, "method ViewAccountDocsWithPII not implemented")
}

/**********************************************************************************************************************/
/*                                                   PRIVATE FUNCTIONS                                                */
/**********************************************************************************************************************/

func getDocName(fname string, uid string, thing string) string {
	return getDocFolder(uid, thing) + "/" + fname
}

func getDocFolder(uid string, thing string) string {
	return JDOCS_PATH + uid + "/" + thing
}

func getUserFolder(uid string) string {
	return JDOCS_PATH + uid
}

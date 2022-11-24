package server

import (
	"context"
	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
	"os"
)

func (*Server) DeleteDoc(ctx context.Context, req *jdocspb.DeleteDocReq) (*jdocspb.DeleteDocResp, error) {
	fname := req.GetDocName()

	print("[DeleteDoc] NAME: " + fname + ", USERID: " + req.UserId + ", THING: " + req.Thing)

	_ = os.Remove(getDocName(fname, req.UserId, req.Thing))

	response := jdocspb.DeleteDocResp{}

	return &response, nil
}

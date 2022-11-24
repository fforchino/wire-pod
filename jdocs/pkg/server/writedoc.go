package server

import (
	"context"
	"os"

	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
)

func (*Server) WriteDoc(ctx context.Context, req *jdocspb.WriteDocReq) (*jdocspb.WriteDocResp, error) {
	fname := req.GetDocName()
	uid := req.UserId
	thing := req.Thing
	version := req.GetDoc().GetDocVersion()
	fileContent := req.GetDoc().JsonDoc

	print("[WriteDoc] NAME: " + fname + ", UID: " + uid + ", THING:" + thing + ", CONTENT: " + fileContent)

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

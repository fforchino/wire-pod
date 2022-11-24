package server

import (
	"context"
	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
	"os"
	"path/filepath"
)

func (*Server) ViewAccountDocs(ctx context.Context, req *jdocspb.ViewAccountDocsReq) (*jdocspb.ViewDocsResp, error) {
	uid := req.UserId

	print("[ViewAccountDocs] USERID: " + uid)

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

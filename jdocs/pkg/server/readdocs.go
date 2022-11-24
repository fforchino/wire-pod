package server

import (
	"context"
	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) ReadDocs(context.Context, *jdocspb.ReadDocsReq) (*jdocspb.ReadDocsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDocs not implemented")
}

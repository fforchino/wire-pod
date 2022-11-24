package server

import (
	"context"
	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) ViewAccountDocsWithPII(context.Context, *jdocspb.ViewAccountDocsReq) (*jdocspb.ViewDocsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAccountDocsWithPII not implemented")
}

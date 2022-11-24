package server

import (
	"context"
	jdocspb "github.com/digital-dream-labs/api/go/jdocspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) PurgeAccountDocs(context.Context, *jdocspb.PurgeAccountDocsReq) (*jdocspb.PurgeAccountDocsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeAccountDocs not implemented")
}

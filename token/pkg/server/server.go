package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/digital-dream-labs/api/go/tokenpb"
	"github.com/digital-dream-labs/hugh/log"
	"time"
)

type options struct {
	log log.Logger
}

// Option is the list of options
type Option func(*options)

type Server struct {
	tokenpb.UnimplementedTokenServer
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

func getToken(appId string, clientName string, generateStsToken bool, revokeClientTokens bool, expirationMinutes uint32, sessionCertificate []byte) tokenpb.TokenBundle {
	sessionToken := jwt.Token{}
	if expirationMinutes == 0 {
		expirationMinutes = 120
	}
	expiresAt := time.Now().Add(time.Duration(expirationMinutes) * time.Second)

	stsToken := tokenpb.StsToken{
		AccessKeyId:     "",
		SecretAccessKey: "",
		SessionToken:    sessionToken.Signature,
		Expiration:      string(expiresAt.Unix()),
	}

	data := tokenpb.TokenBundle{
		Token:       "pincopallo",
		ClientToken: "pallopinco",
		StsToken:    &stsToken,
	}

	return data
}

func (*Server) AssociatePrimaryUser(ctx context.Context, req *tokenpb.AssociatePrimaryUserRequest) (*tokenpb.AssociatePrimaryUserResponse, error) {
	print("[AssociatePrimaryUser] " + req.String())

	data := getToken(req.AppId, req.ClientName, req.GenerateStsToken, req.RevokeClientTokens, req.ExpirationMinutes, req.SessionCertificate)
	response := tokenpb.AssociatePrimaryUserResponse{
		Data: &data,
	}

	return &response, nil
}

func (*Server) ReassociatePrimaryUser(ctx context.Context, req *tokenpb.ReassociatePrimaryUserRequest) (*tokenpb.ReassociatePrimaryUserResponse, error) {
	print("[ReassociatePrimaryUser] " + req.String())

	var cert []byte
	data := getToken(req.AppId, req.ClientName, req.GenerateStsToken, true, req.ExpirationMinutes, cert)
	response := tokenpb.ReassociatePrimaryUserResponse{
		Data: &data,
	}

	return &response, nil
}

func (*Server) AssociateSecondaryClient(ctx context.Context, req *tokenpb.AssociateSecondaryClientRequest) (*tokenpb.AssociateSecondaryClientResponse, error) {
	print("[AssociateSecondaryClientResponse] " + req.String())

	var cert []byte
	data := getToken(req.AppId, req.ClientName, false, false, 0, cert)
	response := tokenpb.AssociateSecondaryClientResponse{
		Data: &data,
	}

	return &response, nil
}

func (*Server) DisassociatePrimaryUser(ctx context.Context, req *tokenpb.DisassociatePrimaryUserRequest) (*tokenpb.DisassociatePrimaryUserResponse, error) {
	print("[DisassociatePrimaryUser] " + req.String())

	response := tokenpb.DisassociatePrimaryUserResponse{}

	return &response, nil
}

func (*Server) RefreshToken(ctx context.Context, req *tokenpb.RefreshTokenRequest) (*tokenpb.RefreshTokenResponse, error) {
	print("[RefreshToken] " + req.String())

	var cert []byte
	data := getToken("", "", req.RefreshStsTokens, req.RefreshJwtTokens, req.ExpirationMinutes, cert)
	response := tokenpb.RefreshTokenResponse{
		Data: &data,
	}

	return &response, nil
}

func (*Server) ListRevokedTokens(ctx context.Context, req *tokenpb.ListRevokedTokensRequest) (*tokenpb.ListRevokedTokensResponse, error) {
	print("[ListRevokedTokens] " + req.String())
	var tokens []string

	data := tokenpb.TokensPage{
		Tokens:  tokens,
		LastKey: "",
		Done:    true,
	}

	response := tokenpb.ListRevokedTokensResponse{
		Data: &data,
	}

	return &response, nil
}

func (*Server) RevokeFactoryCertificate(ctx context.Context, req *tokenpb.RevokeFactoryCertificateRequest) (*tokenpb.RevokeFactoryCertificateResponse, error) {
	print("[RevokeFactoryCertificate] " + req.String())

	response := tokenpb.RevokeFactoryCertificateResponse{}

	return &response, nil
}

func (*Server) RevokeTokens(ctx context.Context, req *tokenpb.RevokeTokensRequest) (*tokenpb.RevokeTokensResponse, error) {
	print("[RevokeFactoryCertificate] " + req.String())

	response := tokenpb.RevokeTokensResponse{
		TokensRevoked: 0,
	}

	return &response, nil
}

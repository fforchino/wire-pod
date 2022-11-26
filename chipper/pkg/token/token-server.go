package token

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

type TokenServer struct {
	tokenpb.UnimplementedTokenServer
}

// New accepts a list of args and returns the service
func New(opts ...Option) (*TokenServer, error) {
	cfg := options{
		log: log.Base(),
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	s := TokenServer{}

	println("Token server created")

	return &s, nil
}

func (*TokenServer) AssociatePrimaryUser(ctx context.Context, req *tokenpb.AssociatePrimaryUserRequest) (*tokenpb.AssociatePrimaryUserResponse, error) {
	println("[AssociatePrimaryUser] " + req.String())

	data := getToken(req.AppId, req.ClientName, req.GenerateStsToken, req.RevokeClientTokens, req.ExpirationMinutes, req.SessionCertificate)
	response := tokenpb.AssociatePrimaryUserResponse{
		Data: &data,
	}
	println("[AssociatePrimaryUser] Sending response: " + response.String())

	return &response, nil
}

func (*TokenServer) ReassociatePrimaryUser(ctx context.Context, req *tokenpb.ReassociatePrimaryUserRequest) (*tokenpb.ReassociatePrimaryUserResponse, error) {
	println("[ReassociatePrimaryUser] " + req.String())

	var cert []byte
	data := getToken(req.AppId, req.ClientName, req.GenerateStsToken, true, req.ExpirationMinutes, cert)
	response := tokenpb.ReassociatePrimaryUserResponse{
		Data: &data,
	}

	return &response, nil
}

func (*TokenServer) AssociateSecondaryClient(ctx context.Context, req *tokenpb.AssociateSecondaryClientRequest) (*tokenpb.AssociateSecondaryClientResponse, error) {
	println("[AssociateSecondaryClientResponse] " + req.String())

	var cert []byte
	data := getToken(req.AppId, req.ClientName, false, false, 0, cert)
	response := tokenpb.AssociateSecondaryClientResponse{
		Data: &data,
	}

	return &response, nil
}

func (*TokenServer) DisassociatePrimaryUser(ctx context.Context, req *tokenpb.DisassociatePrimaryUserRequest) (*tokenpb.DisassociatePrimaryUserResponse, error) {
	println("[DisassociatePrimaryUser] " + req.String())

	response := tokenpb.DisassociatePrimaryUserResponse{}

	return &response, nil
}

func (*TokenServer) RefreshToken(ctx context.Context, req *tokenpb.RefreshTokenRequest) (*tokenpb.RefreshTokenResponse, error) {
	println("[RefreshToken] " + req.String())

	var cert []byte
	data := getToken("", "", req.RefreshStsTokens, req.RefreshJwtTokens, req.ExpirationMinutes, cert)
	response := tokenpb.RefreshTokenResponse{
		Data: &data,
	}

	return &response, nil
}

func (*TokenServer) ListRevokedTokens(ctx context.Context, req *tokenpb.ListRevokedTokensRequest) (*tokenpb.ListRevokedTokensResponse, error) {
	println("[ListRevokedTokens] " + req.String())
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

func (*TokenServer) RevokeFactoryCertificate(ctx context.Context, req *tokenpb.RevokeFactoryCertificateRequest) (*tokenpb.RevokeFactoryCertificateResponse, error) {
	println("[RevokeFactoryCertificate] " + req.String())

	response := tokenpb.RevokeFactoryCertificateResponse{}

	return &response, nil
}

func (*TokenServer) RevokeTokens(ctx context.Context, req *tokenpb.RevokeTokensRequest) (*tokenpb.RevokeTokensResponse, error) {
	println("[RevokeFactoryCertificate] " + req.String())

	response := tokenpb.RevokeTokensResponse{
		TokensRevoked: 0,
	}

	return &response, nil
}

/**********************************************************************************************************************/
/*                                                   PRIVATE FUNCTIONS                                                */
/**********************************************************************************************************************/

func getToken(appId string, clientName string, generateStsToken bool, revokeClientTokens bool, expirationMinutes uint32, sessionCertificate []byte) tokenpb.TokenBundle {
	if expirationMinutes == 0 {
		expirationMinutes = 120
	}
	expiresAt := time.Now().Add(time.Duration(expirationMinutes) * time.Second)

	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "test",
		},
	}
	sessionToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, _ := sessionToken.SignedString(mySigningKey)

	stsToken := tokenpb.StsToken{
		AccessKeyId:     "11",
		SecretAccessKey: "AllYourBase",
		SessionToken:    signedString,
		Expiration:      string(expiresAt.Unix()),
	}

	data := tokenpb.TokenBundle{
		Token:       "pincopallo",
		ClientToken: "pallopinco",
		StsToken:    &stsToken,
	}

	return data
}

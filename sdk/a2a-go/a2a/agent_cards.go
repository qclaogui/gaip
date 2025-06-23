package a2a

import (
	"context"
	"fmt"

	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
)

type AgentCards struct {
	apiClient *apiClient
}

// GetAgentCard retrieves an agent card by its model.
func (a AgentCards) GetAgentCard(ctx context.Context) (*AgentCard, error) {
	req := &pb.GetAgentCardRequest{}
	debugPrint(req)
	res, err := a.apiClient.a2aClient.GetAgentCard(ctx, req)
	if err != nil {
		return nil, err
	}
	return fromProto[AgentCard](res)
}

type OAuthFlows interface {
	toOAuthFlows() *pb.OAuthFlows
}

func oAuthFlowsToProto(o OAuthFlows) *pb.OAuthFlows {
	if o == nil {
		return nil
	}
	return o.toOAuthFlows()
}

func oAuthFlowsFromProto(o *pb.OAuthFlows) OAuthFlows {
	switch d := o.Flow.(type) {
	case *pb.OAuthFlows_AuthorizationCode:
		return (AuthorizationCodeOAuthFlow{}).fromProto(d.AuthorizationCode)
	case *pb.OAuthFlows_ClientCredentials:
		return (ClientCredentialsOAuthFlow{}).fromProto(d.ClientCredentials)
	case *pb.OAuthFlows_Implicit:
		return (ImplicitOAuthFlow{}).fromProto(d.Implicit)
	case *pb.OAuthFlows_Password:
		return (PasswordOAuthFlow{}).fromProto(d.Password)
	default:
		panic(fmt.Errorf("unknown OAuthFlows.Flow type %T", o.Flow))
	}
}

func (o *AuthorizationCodeOAuthFlow) toOAuthFlows() *pb.OAuthFlows {
	return &pb.OAuthFlows{
		Flow: &pb.OAuthFlows_AuthorizationCode{
			AuthorizationCode: o.toProto(),
		},
	}
}

func (o *ClientCredentialsOAuthFlow) toOAuthFlows() *pb.OAuthFlows {
	return &pb.OAuthFlows{
		Flow: &pb.OAuthFlows_ClientCredentials{
			ClientCredentials: o.toProto(),
		},
	}
}

func (o *ImplicitOAuthFlow) toOAuthFlows() *pb.OAuthFlows {
	return &pb.OAuthFlows{
		Flow: &pb.OAuthFlows_Implicit{
			Implicit: o.toProto(),
		},
	}
}

func (o *PasswordOAuthFlow) toOAuthFlows() *pb.OAuthFlows {
	return &pb.OAuthFlows{
		Flow: &pb.OAuthFlows_Password{
			Password: o.toProto(),
		},
	}
}

type SecurityScheme interface {
	toSecurityScheme() *pb.SecurityScheme
}

func securitySchemeToProto(s SecurityScheme) *pb.SecurityScheme {
	if s == nil {
		return nil
	}
	return s.toSecurityScheme()
}

func securitySchemeFromProto(s *pb.SecurityScheme) SecurityScheme {
	switch d := s.Scheme.(type) {
	case *pb.SecurityScheme_ApiKeySecurityScheme:
		return (APIKeySecurityScheme{}).fromProto(d.ApiKeySecurityScheme)
	case *pb.SecurityScheme_HttpAuthSecurityScheme:
		return (HTTPAuthSecurityScheme{}).fromProto(d.HttpAuthSecurityScheme)
	case *pb.SecurityScheme_Oauth2SecurityScheme:
		return (OAuth2SecurityScheme{}).fromProto(d.Oauth2SecurityScheme)
	case *pb.SecurityScheme_OpenIdConnectSecurityScheme:
		return (OpenIdConnectSecurityScheme{}).fromProto(d.OpenIdConnectSecurityScheme)
	default:
		panic(fmt.Errorf("unknown SecurityScheme.Scheme type %T", s.Scheme))
	}
}

func (s *APIKeySecurityScheme) toSecurityScheme() *pb.SecurityScheme {
	return &pb.SecurityScheme{
		Scheme: &pb.SecurityScheme_ApiKeySecurityScheme{
			ApiKeySecurityScheme: s.toProto(),
		},
	}
}

func (s *HTTPAuthSecurityScheme) toSecurityScheme() *pb.SecurityScheme {
	return &pb.SecurityScheme{
		Scheme: &pb.SecurityScheme_HttpAuthSecurityScheme{
			HttpAuthSecurityScheme: s.toProto(),
		},
	}
}

func (s *OAuth2SecurityScheme) toSecurityScheme() *pb.SecurityScheme {
	return &pb.SecurityScheme{
		Scheme: &pb.SecurityScheme_Oauth2SecurityScheme{
			Oauth2SecurityScheme: s.toProto(),
		},
	}
}

func (s *OpenIdConnectSecurityScheme) toSecurityScheme() *pb.SecurityScheme {
	return &pb.SecurityScheme{
		Scheme: &pb.SecurityScheme_OpenIdConnectSecurityScheme{
			OpenIdConnectSecurityScheme: s.toProto(),
		},
	}
}

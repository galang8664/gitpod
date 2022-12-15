// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package oidc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestGetStartParams(t *testing.T) {
	const (
		issuerG  = "https://accounts.google.com"
		clientID = "client-id-123"
	)
	service := NewOIDCService()
	config := &OIDCClientConfig{
		ID:         "google-1",
		Issuer:     issuerG,
		OIDCConfig: &oidc.Config{},
		OAuth2Config: &oauth2.Config{
			ClientID: clientID,
			Endpoint: oauth2.Endpoint{
				AuthURL: issuerG + "/o/oauth2/v2/auth",
			},
		},
	}

	params, err := service.GetStartParams(config)

	require.NoError(t, err)
	require.NotNil(t, params.Nonce)
	require.NotNil(t, params.State)

	// AuthCodeURL example:
	// https://accounts.google.com/o/oauth2/v2/auth
	// ?client_id=client-id-123
	// &nonce=UFTMxxUtc5jVZbp2a2R9XEoRwpfzs-04FcmVQ-HdCsw
	// &response_type=code
	// &state=Q4XzRcdo4jtOYeRbF17T9LHHwX-4HacT1_5pZH8mXLI
	require.NotNil(t, params.AuthCodeURL)
	require.Contains(t, params.AuthCodeURL, issuerG)
	require.Contains(t, params.AuthCodeURL, clientID)
	require.Contains(t, params.AuthCodeURL, params.Nonce)
	require.Contains(t, params.AuthCodeURL, params.State)
}

func TestGetClientConfigFromRequest(t *testing.T) {
	const (
		issuerG = "https://accounts.google.com"
	)
	testCases := []struct {
		Location      string
		ExpectedError bool
		ExpectedId    string
	}{
		{
			Location:      "/start?word=abc",
			ExpectedError: true,
			ExpectedId:    "",
		},
		{
			Location:      "/start?issuer=https%3A%2F%2Faccounts.google.com",
			ExpectedError: false,
			ExpectedId:    "google-1",
		},
		{
			Location:      "/start?issuer=UNKNOWN",
			ExpectedError: true,
			ExpectedId:    "",
		},
	}

	service := NewOIDCService()
	err := service.AddClientConfig(&OIDCClientConfig{
		ID:           "google-1",
		Issuer:       issuerG,
		OIDCConfig:   &oidc.Config{},
		OAuth2Config: &oauth2.Config{},
	})
	require.NoError(t, err, "failed to initialize test")

	for _, tc := range testCases {
		t.Run(tc.Location, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tc.Location, nil)
			config, err := service.GetClientConfigFromRequest(request)
			if tc.ExpectedError == true {
				require.Error(t, err)
			}
			if tc.ExpectedError != true {
				require.NoError(t, err)
				require.NotNil(t, config)
				require.Equal(t, tc.ExpectedId, config.ID)
			}
		})
	}
}

func TestAuthenticate(t *testing.T) {
	t.Skip() //
	const (
		issuerG = "https://accounts.google.com"
	)
	service := NewOIDCService()
	err := service.AddClientConfig(&OIDCClientConfig{
		ID:     "google-1",
		Issuer: issuerG,
		OIDCConfig: &oidc.Config{
			SkipClientIDCheck: true,
		},
		OAuth2Config: &oauth2.Config{},
	})
	require.NoError(t, err, "failed to initialize test")

	token := oauth2.Token{}
	extra := map[string]interface{}{
		"id_token": "foo123",
	}

	ctx := context.Background()
	nonceCookieValue := "foobar123"
	result, err := service.Authenticate(ctx, &OAuth2Result{
		OAuth2Token: token.WithExtra(extra),
	}, issuerG, nonceCookieValue)

	require.NoError(t, err, "failed to authenticate")
	require.NotNil(t, result)

}

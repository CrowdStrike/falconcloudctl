// Copyright (c) 2022 CrowdStrike, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package config

import (
	"context"
	"fmt"

	"github.com/crowdstrike/falcon-cli/internal/build"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/spf13/viper"
)

// Struct to hold persistent configuration for falcon
type Config struct {
	// The Falcon Customer ID
	CID string `mapstructure:"cid,omitempty"`
	// The Falcon API client ID.
	ClientID string `mapstructure:"client_id"`
	// The Falcon API client secret.
	ClientSecret string `mapstructure:"client_secret"`
	// The Falcon API base URL.
	MemberCID string `mapstructure:"member_cid,omitempty"`
	// The Falcon API cloud region.
	Cloud string `mapstructure:"cloud,omitempty"`
	// The OAuth token returned from the Falcon API.
	OauthToken string `mapstructure:"oauth_token,omitempty"`
	// The Container Registry OAuth token returned from the Falcon API.
	RegistryToken string `mapstructure:"registry_token,omitempty"`
}

var ConfigFile string

func NewConfig() (Config, error) {
	c := &Config{}

	if c.Cloud == "" {
		c.Cloud = "autodiscover"
	}

	c.CID = viper.GetViper().GetString("cid")
	c.ClientID = viper.GetViper().GetString("client_id")
	c.ClientSecret = viper.GetViper().GetString("client_secret")
	c.MemberCID = viper.GetViper().GetString("member_cid")
	c.Cloud = viper.GetViper().GetString("cloud")

	return *c, nil
}

func (c Config) ApiConfig(appVersion string) *falcon.ApiConfig {
	return &falcon.ApiConfig{
		ClientId:          c.ClientID,
		ClientSecret:      c.ClientSecret,
		MemberCID:         c.MemberCID,
		Cloud:             falcon.Cloud(c.Cloud),
		Context:           context.Background(),
		UserAgentOverride: fmt.Sprintf("falcon-cli/%s", build.Version),
	}
}

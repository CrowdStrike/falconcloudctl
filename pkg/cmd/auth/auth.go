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

package auth

import (
	"github.com/crowdstrike/falcon-cli/pkg/utils"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"

	authConfigCmd "github.com/crowdstrike/falcon-cli/pkg/cmd/auth/config"
)

var (
	longDescAuth = templates.LongDesc(`Initialize the Falcon CLI tool`)
	examplesAuth = templates.Examples(`
	    # Initialize the CrowdStrike Falcon CLI tool
	    falcon init
		
		# Initialize the CrowdStrike Falcon CLI tool with defining OAuth2 client ID and secret
		falcon init --client-id <client_id> --client-secret <client_secret>
    `)
)

func NewAuthCmd(f *utils.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auth <command>",
		Short:   "Authenticate falcon CLI with CrowdStrike Falcon API",
		Long:    longDescAuth,
		Example: examplesAuth,
	}

	// Add subcommands
	cmd.AddCommand(authConfigCmd.NewCmdConfig(f))

	return cmd
}
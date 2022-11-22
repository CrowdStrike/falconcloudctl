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

package version

import (
	"fmt"
	"strings"

	"github.com/crowdstrike/falcon-cli/internal/build"
	"github.com/crowdstrike/falcon-cli/pkg/utils"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	shortDesc = `Print the falcon version`
	longDesc  = templates.LongDesc(`Print the falcon version`)
	examples  = templates.Examples(`
        # Print the Falcon CLI and GO version information
        falcon version
    `)
)

func NewCmdVersion(f *utils.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   shortDesc,
		Long:    longDesc,
		Example: examples,
		RunE:    runVer(f),
		Annotations: map[string]string{
			"authSkipCheck": "true",
		},
	}
	return cmd
}

func VersionFormat(version string) string {
	version = strings.TrimPrefix(version, "v")

	return fmt.Sprintf("falcon cli version: %s commit: %s, build date: %s\n", version, build.Commit, build.Date)
}

func runVer(f *utils.Factory) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Fprint(f.IOStreams.Out, VersionFormat(build.Version))
		return nil
	}
}

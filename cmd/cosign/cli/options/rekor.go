//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"github.com/spf13/cobra"
)

// RekorOptions is the wrapper for Rekor related options.
type RekorOptions struct {
	URL string
}

var _ Interface = (*RekorOptions)(nil)

// AddFlags implements Interface
func (o *RekorOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.URL, "rekor-url", "https://rekor.sigstore.dev",
		"[EXPERIMENTAL] address of rekor STL server")
}

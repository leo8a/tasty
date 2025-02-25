// Package cmd
/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/karmab/tasty/pkg/operator"
	"github.com/spf13/cobra"
)

func NewInstaller() *cobra.Command {
	var o *operator.Operator
	var wait, out bool
	var ns, ch string
	cmd := &cobra.Command{
		Use:          "install [operator]",
		Short:        "install operators",
		Long:         `install operators`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			o = operator.NewOperator()
			return o.InstallOperator(wait, out, ns, ch, args)
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&wait, "wait", "w", false, "Wait for crd to show up")
	flags.BoolVarP(&out, "stdout", "s", false, "Print to stdout")
	flags.StringVarP(&ns, "namespace", "n", "", "Target namespace")
	flags.StringVarP(&ch, "channel", "c", "", "Target channel")

	return cmd
}

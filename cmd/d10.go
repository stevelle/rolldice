/*
Copyright © 2020 Steve Lewis <stevelle@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/stevelle/rolldice/lib"

	"github.com/spf13/cobra"
)

// d10Cmd represents the d10 command
var d10Cmd = &cobra.Command{
	Use:   "d10",
	Short: "Roll a single 10-sided die",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("d10 => %d\n", lib.Roll(10))
	},
}

func init() {
	rootCmd.AddCommand(d10Cmd)
}

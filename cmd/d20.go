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

	"github.com/spf13/cobra"
	"github.com/stevelle/rolldice/lib"
)

// d20Cmd represents the d20 command
var d20Cmd = &cobra.Command{
	Use:   "d20",
	Short: "Roll a single 20-sided die",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("d20 => %d\n", lib.Roll(20))
	},
}

func init() {
	rootCmd.AddCommand(d20Cmd)
}

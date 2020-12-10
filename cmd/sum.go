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
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stevelle/rolldice/lib"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Aliases: []string {"total", "complex", "add"},
	Short: "Roll an arbitrary combination of dice",
	Long: `Roll any combination of dice. 

Specify a set of dice by the number of dice, followed by a 'd', followed by
the number of sides on that set of dice.

Specify multiple sets of dice by separating each set by either a space ' ' or
with a plus '+'.

Specify a fixed bonus set by specifying the bonus number just as you would
for another set of dice. 

For example:

2d8 will calculate the sum of d8 + d8 in the range of 2-16.

3d6 + 2 will calculate the sum of d6 + d6 + d6 + 2 in the range of 5-20.
Notice the first set and the bonus set may be separated by spaces and a plus.

3d12 2d6 5 will calculate the sum of d12 + d12 + d12 + d6 + d6 + 5 in the
range of 10-53.
Notice the second dice set and the bonus set may be separated by a space.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sum := int64(0)

		sets := splitArgs(args)
		for _, set := range sets {
			rolls := lib.RollDice(set)

			if fumbledWildDie(rolls) {
				lost := maxIn(rolls)
				fmt.Printf("Fumble: -%d\n", lost)
				// exclude the largest value and the fumble die
				sum -= lost + 1
			}

			// handle critical success on wild die
			for last := lastOf(rolls); last == set.Sides; {
				fmt.Printf("Blown up\n")
				// add a bonus die
				last = lib.Roll(set.Sides)
				rolls = append(rolls, last)
			}

			for _, v := range rolls {
				sum += v
			}
			// TODO format the rolls to display with --verbose
			fmt.Println("Rolls: ", rolls)
			sum += set.Bonus
		}
		fmt.Printf("%d\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
}

/*
	Splits arguments into DiceSets

	Splits occur over whitespace or plus

	A DiceSet can be expressed as either
      a number of dice, a "D" (case-insensitive), then the number of sides of each of the dice in that set
      a bonus
 */
func splitArgs(args []string) []*lib.DiceSet {
	sets := make([]*lib.DiceSet, 0)
	for _, part := range args {
		for _, segment := range strings.Split(part, "+") {
			// ignore blank segments
			if isBlank(segment) {
				continue
			}
			if strings.ContainsAny(segment, "Dd") {
				// the segment represents a number of dice
				sets = append(sets, parseDiceSet(segment, "d"))
			} else {
				// the segment represents a bonus
				if len(sets) > 0 {
					// add the bonus to the prior dice set
					sets[len(sets)-1].Bonus += parsePosInt64(segment)
				} else {
					// this is the first dice set, create a standalone-bonus
					sets = append(sets, lib.NewBonus(parsePosInt64(segment)))
				}
			}
		}
	}
	return sets
}

func isBlank(set string) bool {
	return len(strings.TrimSpace(strings.ReplaceAll(set, "+", ""))) == 0
}

/*
	Parse an expression into a DiceSet, given a delimiter

	The delimiter separates the number of dice from the number of sides for those dice.
	The expression and delimiter are treated case-insensitively.

	Fatal errors resulting in an error message and an exit code of 1 include:
		finding illegal characters (those not used to define a number) on either side of the separator
		finding negative numbers
		finding more than one occurrence of the separator.
 */
func parseDiceSet(expr string, sep string) *lib.DiceSet {
	parts := strings.Split(strings.ToLower(expr), strings.ToLower(sep))
	// TODO FUTURE? handle "d6" using default dice=1 and "3d" using default sides=6
 	if len(parts) != 2 {
		fatal("Could not determine desired dice from \"s\"", expr)
	}
	set := lib.NewDiceSet(parsePosInt64(parts[0]), parsePosInt64(parts[1]))
	return set
}

// Parse a positive integer from a string
func parsePosInt64(expr string) int64 {
	// TODO better error message when len(expr) == 0
	parsed, err := strconv.ParseInt(expr, 10, 64)
	if err != nil {
		fatal("Could not read number from \"%s\"", expr)
	}
	if parsed <= 0 {
		fatal("Cannot operate against non-positive values like \"%d\"", parsed)
	}
	return parsed
}

// Print a formatted error to stderr and exit
func fatal(msg string, param interface{}) {
	os.Stderr.WriteString(fmt.Sprintf("ERROR: " + msg + "\n", param))
	os.Exit(1)
}

// Return the last value
func lastOf(values []int64) int64 {
	return values[len(values)-1]
}

// is the wild die a '1'
func fumbledWildDie(rolls []int64) bool {
	return lastOf(rolls) == 1
}

// identify the largest value in a list
func maxIn(values []int64) int64 {
	maximum := int64(0)
	for _, val := range values {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

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
package lib

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type DiceSet struct {
	Dice  int64
	Sides int64
	Bonus int64
	IsWild bool
}

func NewDiceSet(dice int64, sides int64, isWild bool) *DiceSet {
	return &DiceSet{Dice: dice, Sides: sides, Bonus: 0, IsWild: isWild}
}

func NewBonus(bonus int64) *DiceSet {
	return &DiceSet{Dice: 0, Sides: 0, Bonus: bonus, IsWild: false}
}

// Roll a single die with a given number of sides
func Roll(sides int64) int64 {
	value, err := rand.Int(rand.Reader, big.NewInt(sides))
	if err != nil {
		panic(err)
	}
	return value.Int64() + 1
}

// Roll a DiceSet and return the results for each die, or bonus
func (d *DiceSet) Roll() []int64 {
	results := make([]int64, d.Dice)
	for i := range results {
		results[i] = Roll(d.Sides)
	}
	return results
}

// Describe a DiceSet with the results it produced in a single line
func Describe(set *DiceSet, r []int64) string {
	if set.Bonus != 0 {
		return fmt.Sprintf("%dd%d+%d\t=> [%d]", set.Dice, set.Sides, set.Bonus, r[0])
	}

	return fmt.Sprintf("%dd%d\t%s=> []", set.Dice, set.Sides, strings.Join(formatInts(r), ", "))
}

// Convert a slice of ints to a slice of strings
func formatInts(r []int64) []string {
	s := make([]string, len(r))
	for i, n := range r {
		s[i] = strconv.FormatInt(n, 10)
	}
	return s
}
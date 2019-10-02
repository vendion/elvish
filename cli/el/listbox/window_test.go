package listbox

import (
	"testing"

	"github.com/elves/elvish/tt"
)

var Args = tt.Args

func TestGetVerticalWindow(t *testing.T) {
	tt.Test(t, tt.Fn("getVerticalWindow", getVertialWindow), tt.Table{
		// selected = 0: always show a widow starting from 0, regardless of
		// the value of oldFirst
		Args(State{TestItems{NItems: 10}, 0, 0}, 6).Rets(0, 0),
		Args(State{TestItems{NItems: 10}, 0, 1}, 6).Rets(0, 0),
		// selected = n-1: always show a window ending at n-1, regardless of the
		// value of oldFirst
		Args(State{TestItems{NItems: 10}, 9, 0}, 6).Rets(4, 0),
		Args(State{TestItems{NItems: 10}, 9, 8}, 6).Rets(4, 0),
		// selected = 3, oldFirst = 2 (likely because previous selected = 4).
		// Adjust first -> 1 to satisfy the upward respect distance of 2.
		Args(State{TestItems{NItems: 10}, 3, 2}, 6).Rets(1, 0),
		// selected = 6, oldFirst = 2 (likely because previous selected = 7).
		// Adjust first -> 3 to satisfy the downward respect distance of 2.
		Args(State{TestItems{NItems: 10}, 6, 2}, 6).Rets(3, 0),

		// There is not enough budget to achieve respect distance on both sides.
		// Split the budget in half.
		Args(State{TestItems{NItems: 10}, 3, 1}, 3).Rets(2, 0),
		Args(State{TestItems{NItems: 10}, 3, 0}, 3).Rets(2, 0),

		// There is just enough distance to fit the selected item. Only show the
		// selected item.
		Args(State{TestItems{NItems: 10}, 2, 0}, 1).Rets(2, 0),
	})
}

func TestGetHorizontalWindow(t *testing.T) {
	tt.Test(t, tt.Fn("getHorizontalWindow", getHorizontalWindow), tt.Table{
		// All items fit in a single column.
		Args(State{TestItems{NItems: 10}, 4, 0}, 6, 10).Rets(0, true),
		// All items fit in multiple columns.
		Args(State{TestItems{Prefix: "x", NItems: 10}, 4, 0}, 6, 5).Rets(0, true),
		// All items cannot fit, selected = 0; show a window from 0.
		Args(State{TestItems{Prefix: "x", NItems: 11}, 0, 0}, 6, 5).Rets(0, false),
		// All items cannot fit. Columsn are 0-3, 4-7, 8-10 (height reduced from
		// 5 to 4 for scrollbar). Selecting last item; show last two columns.
		Args(State{TestItems{Prefix: "x", NItems: 11}, 10, 0}, 7, 5).Rets(4, false),
		// Items are wider than terminal. Treat them as if they are trimmed.
		// Columns are 0-4, 5-9.
		Args(State{TestItems{Prefix: "long prefix", NItems: 10}, 9, 0},
			6, 6).Rets(5, false),

		// The following cases only differ in State.First and shows that the
		// algorithm respects it. In alls cases, the columns are 0-4, 5-9,
		// 10-14, 15-19, item 10 is selected, and the terminal can fit 2 columns.

		// First = 0. Try to reach as far as possible to that, ending up showing
		// columns 5-9 and 10-14.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 0}, 8, 6).Rets(5, false),
		// First = 2. Ditto.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 2}, 8, 6).Rets(5, false),
		// First = 5. Show columns 5-9 and 10-14.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 5}, 8, 6).Rets(5, false),
		// First = 7. Ditto.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 7}, 8, 6).Rets(5, false),
		// First = 10. No need to any columns to the left.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 10}, 8, 6).Rets(10, false),
		// First = 12. Ditto.
		Args(State{TestItems{Prefix: "x", NItems: 20}, 10, 12}, 8, 6).Rets(10, false),
	})
}
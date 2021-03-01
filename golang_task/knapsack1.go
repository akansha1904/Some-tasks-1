package main

import "fmt"

/*
1. A vase that weights 3 pounds and is worth 50 dollars.
2. A silver nugget that weights 6 pounds and is worth 30 dollars.
3. A painting that weights 4 pounds and is worth 40 dollars.
4. A mirror that weights 5 pounds and is worth 10 dollars.
*/

type Item struct {
	name   string
	weight int
	value  int
}

type Solution struct {
	items       *[]Item // To avoid unneccessary copies, I just reference the array
	totalValue  int
	totalWeight int
}

func knapsack(items []Item, knapsackSize int) Solution {
	// solution[x] is for the solutions for knapsack size x.
	// solution[x][y] is for the solution for knapsack size x, before adding item[y]
	// So the final solution will be in solution[knapsackSize][number_of_items + 1]
	solutions := make([][]Solution, knapsackSize+1)

	empty := make([]Item, 0)

	for j := 0; j <= knapsackSize; j++ {
		solutions[j] = make([]Solution, len(items)+1)
		for i := 0; i <= len(items); i++ {
			solutions[j][i].items = &empty
		}
	}

	// We build up our possible solutions by adding the items one by one
	for i := 0; i < len(items); i++ {
		// ... for all possible knapsack sizes up to the size we care about
		for j := 1; j <= knapsackSize+1; j++ {
			if items[i].weight < j {
				solution := &solutions[j-1][i]

				altSolution := &solutions[j-1-items[i].weight][i]

				if solution.totalValue < altSolution.totalValue+items[i].value {
					// extra work here, because we used an array ref for Solution.items
					newItems := make([]Item, len(*altSolution.items)+1)
					copy(newItems, *altSolution.items)
					newItems[len(*altSolution.items)] = items[i]

					newSolution := Solution{
						&newItems,
						altSolution.totalValue + items[i].value,
						altSolution.totalWeight + items[i].weight,
					}
					solution = &newSolution
				}
				solutions[j-1][i+1] = *solution
			} else {
				solutions[j-1][i+1] = solutions[j-1][i]
			}
		}
	}

	// Since this is an iterative process, we would have solved the knapsack for all sizes up to the
	// given size. Show them.
	for j := 1; j <= knapsackSize; j++ {
		fmt.Println("Solution for size", j,
			"is", *(solutions[j][len(items)].items),
			"weight", solutions[j][len(items)].totalWeight,
			"value", solutions[j][len(items)].totalValue,
		)
	}

	return solutions[knapsackSize][len(items)]
}

func main() {
	items := []Item{
		{"salt", 3, 50},
		{"sugar", 6, 30},
		{"vegetables", 4, 40},
		{"dal", 5, 10},
	}

	knapsacksize := 20

	fmt.Println("For items", items)
	s := knapsack(items, knapsacksize)
	fmt.Println("Final solution for size", knapsacksize,
		"is", *s.items,
		"weight", s.totalWeight,
		"value", s.totalValue,
	)

	items = append(items, Item{"diamond", 1, 100})
	items = append(items, Item{"feather", 1, 1})

	fmt.Println("For items", items)
	s = knapsack(items, knapsacksize)
	fmt.Println("Final solution for size", knapsacksize,
		"is", *s.items,
		"weight", s.totalWeight,
		"value", s.totalValue,
	)

	items = append(items, Item{"dust", 0, 1})
	fmt.Println("For items", items)
	s = knapsack(items, knapsacksize)
	fmt.Println("Final solution for size", knapsacksize,
		"is", *s.items,
		"weight", s.totalWeight,
		"value", s.totalValue,
	)

}

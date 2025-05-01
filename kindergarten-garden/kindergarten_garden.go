package kindergarten

import (
	"strings"
	"errors"
	"slices"
)

// Define the Garden type here.
type Garden map[string][]string

var plants = map[rune]string{
	'G': "grass",
	'V': "violets",
	'C': "clover",
	'R': "radishes",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func PreChecks(diagram string, children []string) bool {
	if !strings.HasPrefix(diagram,"\n") {
		return false
	}
	if len(strings.Split(strings.Replace(diagram,"\n","",1), "\n")) != 2 {
		return false
	}
	if len(strings.Split(strings.Replace(diagram,"\n","",1), "\n")[0]) !=  len(strings.Split(strings.Replace(diagram,"\n","",1), "\n")[1]) || len(strings.Split(strings.Replace(diagram,"\n","",1), "\n")[0]) % 2 != 0 {
		return false
	}
	uniqueChildren := []string{}
	for _,x := range children {
		if !slices.Contains(uniqueChildren, x) {
			uniqueChildren = append(uniqueChildren, x)
		}
	}
	if len(uniqueChildren) != len(children) {
		return false
	}
	return true
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	preCheck := PreChecks(diagram, children)
	if !preCheck {
		return nil, errors.New("Wrong input format")
	}
	diagram = strings.Replace(diagram,"\n","",1)
	diagrams := strings.Split(diagram, "\n")
	windowDiagram, internalDiagram := diagrams[0], diagrams[1]
	kidsPlants := map[string][]string{}
	kinder := slices.Clone(children)
	if !slices.IsSorted(kinder) {
		slices.Sort(kinder)
	}
	for i,_ := range windowDiagram {
		if i%2 == 0 {
			kidPlants := windowDiagram[i:i+2] + internalDiagram[i:i+2]
			nPlants := []string{}
			for _,x := range kidPlants {
				plant, ok := plants[x]
				if !ok {
					return nil, errors.New("Wrong input format")
				} 
				nPlants = append(nPlants, plant)
			}
			kidsPlants[kinder[i/2]] = nPlants
		}
	}
	var g Garden
	g = kidsPlants
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	val, ok := (*g)[child]
	return val, ok
}
package sublist

import (
	"strings"
	"strconv"
)

func Sublist(l1, l2 []int) Relation {
	if len(l1) != 0 && len(l2) == 0 {
		return RelationSuperlist
	}
	if len(l1) == 0 && len(l2) != 0 {
		return RelationSublist
	}
	l1S := []string{}
	l2S := []string{}
	for _,x := range l1 {
		l1S = append(l1S, strconv.Itoa(x))
	}
	for _,x := range l2 {
		l2S = append(l2S, strconv.Itoa(x))
	}
	l1str := strings.Join(l1S, " ")+" "
	l2str := strings.Join(l2S, " ")+" "
	if l1str == l2str {
		return RelationEqual
	}
	if strings.Contains(l1str, l2str) {
		return RelationSuperlist
	} 
	if strings.Contains(l2str, l1str) {
		return RelationSublist
	} 
	return RelationUnequal
}

package strand

import "strings"

var dnaToRna = map[string]string{"A": "U", "T": "A", "C": "G", "G": "C"}

func ToRNA(dna string) string {
	rna := ""
	for _,x := range dna {
		rna += dnaToRna[strings.ToUpper(string(x))]
	}
	return rna
}

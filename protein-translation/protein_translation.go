package protein

import (
	"errors"
)

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("This is a stop codon!")
var ErrInvalidBase = errors.New("This codon has an invalid base!")

func FromRNA(rna string) ([]string, error) {
	codons := []string{}
	for i,_ := range rna {
		if i%3 == 0 && i+2 < len(rna) {
			codons = append(codons, rna[i:i+3])
		}
	}
	translatedProt := []string{}
	for _,x := range codons {
		amino, err := FromCodon(x)
		if err != nil {
			if codonToProtein[x] == "STOP" {
				break
			} else {
				return nil, ErrInvalidBase
			}
		} else {
			translatedProt = append(translatedProt, amino)
		}
	}
	return translatedProt, nil
}

func FromCodon(codon string) (string, error) {
	amino, ok := codonToProtein[codon]
	if ok {
		if amino == "STOP" {
			return "", ErrStop
		} 
		return amino, nil
	} else {
		return "", ErrInvalidBase
	}
}

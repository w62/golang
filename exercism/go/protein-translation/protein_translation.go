package protein

import (
	"fmt"
	"errors"
)

const rlength = 3

var ErrStop error 
var ErrInvalidBase error 

func FromRNA(rna string) ([]string, error) {
	var codon []string
	var protein []string

	fmt.Println("rna", rna)

	if len(rna) % 3 != 0 {
		return nil, errors.New("Invalid rna")
	}

	for i := 0; i < len(rna); i += rlength {
		fmt.Println("Here?", i)
		codon[i] = string(rna[i]+rlength-1)
		val, err := FromCodon(codon[i])
		if err != nil {
			return nil, errors.New("Invalid codon")
		}
//		protein[i] = val
		protein = append(protein, val)
		
	}
	fmt.Printf("codon=%v\n", codon)
	fmt.Printf("protein=%v\n", protein)
	return protein, nil
	panic("Please implement the FromRNA function")
}


func FromCodon(codon string) (string, error) {
	fmt.Println(codon)
	protein := ""
	switch {
	case codon == "AUG":
		protein = "Methionine"
	case codon == "UUU" || codon == "UUC":
		protein = "Phenylalanine"
	case codon == "UUA" || codon == "UUG":
		protein = "Leucine"
	case codon == "UCU" || codon == "UCC" ||
	codon == "UCA" || codon == "UCG":
		protein = "Serine"
	case codon == "UAU" || codon == "UAC":
		protein = "Tyrosine"
	case codon == "UGU" || codon == "UGC":
		protein = "Cysteine"
	case codon == "UGG":
		protein = "Tryptophan"
	case codon == "UAA" || codon == "UAG" ||
	codon == "UGA" :
		protein = "" //STOP
	default:
		return "" , nil
		
	}
	return protein, nil
	panic("Please implement the FromCodon function")
}

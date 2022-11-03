package protein

import (
	"errors"
)

const rlength = 3

var ErrStop error = errors.New("stop")
var ErrInvalidBase error = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var codon string
	var protein []string

	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	for i := 0; i < len(rna); i += rlength {
		codon = rna[i : i+rlength]
		val, err := FromCodon(codon)
		if err == ErrStop  {
			break
		} else if val != "" {
			protein = append(protein, val)
	
		}
	}

	return protein, nil
	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
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
		codon == "UGA":
		protein = "" //STOP
		return "", ErrStop
	default:
		return "", ErrInvalidBase

	}
	return protein, nil
	panic("Please implement the FromCodon function")
}

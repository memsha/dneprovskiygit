package strand

func ToRNA(dna string) string {
	mapa := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	rna := ""
	for _, b := range dna {
		rna += string(mapa[b])
	}
	return rna
}

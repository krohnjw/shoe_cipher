package main

import "fmt"

var CIPHER = []string{
	"A", "D", "G", "J", "L", "Z", "C", "B", "M", "W", "R", "Y", "I", "P", "S", "F", "H", "K", "X", "V", "N", "Q", "E", "T", "U", "O",
}

var ALPHA = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func main() {

	// July 4th letter
	encoded := "GR VRRW BCMHAQ. S MCV'O XCHR OPR MCKK. OPSU QKCMR SU YAKK NY MNXQAORT PCMHRTU CVW OPR QPNVRU CTR CKK BAIIRW. UQL PNOKSVR MNWR SU PSIPONQ'U OSMHRO VAXBRT XSVAU CWXSUUSNV QTSMR."
	match := "C"
	decoded := decodeString(encoded, match, false)
	fmt.Printf("Decoded string:\n%v\n", decoded)

}

func decodeString(encoded, match string, debug bool) string {
	// Create a hash of ciphered letters back to alpha for replacement
	decoded := ""
	fmt.Printf("Encoded String:\n%v\n", encoded)
	// Find Cipher Index
	cipherIndex := findCipherIndex(match)
	// Use this index to build a map of Cipered letters to Alpha
	alphaIndex := 0
	matched := 0
	total := 26
	decoder := make(map[string]string)
	for matched < total {
		cipherChar := CIPHER[cipherIndex]
		alphaChar := ALPHA[alphaIndex]
		decoder[cipherChar] = alphaChar
		cipherIndex = cipherIndex + 1
		if cipherIndex > len(CIPHER)-1 {
			cipherIndex = 0
		}
		alphaIndex = alphaIndex + 1
		if alphaIndex > len(ALPHA)-1 {
			alphaIndex = 0
		}
		matched = matched + 1
	}

	for _, r := range encoded {
		a, ok := decoder[string(r)]
		if ok {
			decoded = decoded + a
		} else {
			if debug {
				fmt.Printf("Missing cipher char for %v\n", string(r))
			}
			decoded = decoded + string(r)
		}
	}

	return decoded

}
func findCipherIndex(match string) int {
	for i, v := range CIPHER {
		if v == match {
			return i
		}
	}
	return 0
}

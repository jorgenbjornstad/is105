package huffman

import (
  "errors"
)


func DecodeHuffmanString(code string) (string, error) {
  var i int
  var decoded string
  var err error
  // Opretter et map som viser hvilken kode hvert fakultet har.
  key := getKey()
  for {
    i = 2
    // Ser om lengden på koden er 2 eller større.
    // Om den ikke er det er ikke koden lang nok til å returnere noe.
    if len(code) < i {
      err = errors.New("Invalid length of code")
      break
      // Dersom de 2 neste elementene i code ikke gir en rune, settes i til
      // 3 slik av vi kan se om de 3 neste elementene gir en rune.
    } else if key[code[:i]] == 0 {
      i = 3
    }
    // Ser om lengen på koden er 3 eller større.
    // Om den ikke er det er ikke koden lang nok til å returnere noe.
    if len(code) < i {
      err = errors.New("Invalid length of code")
      break
    }
    // Dersom det er 3 eller flere elementer i code ser vi om de 3 neste
    // elementene tilsvarer en rune. Dersom de ikke gjør det er det noe
    // feil med stringen som er gitt til funksjonen.
    if key[code[:i]] == 0 {
      err = errors.New("Malformed huffman code, no matching key")
      break
    }
    // Legger til de 2 eller 3 første (avhengig av if-setningene over) til en
    // ny liste (decoded).
    decoded += string(key[code[:i]])
    // Oppdaterer stringen som ble tatt inn som parameter slik at vi kan
    // begynne på nytt i neste iterering i for-loopen.
    code = code[i:]
    // Når listen blir tom avsluttes for-loopen.
    if len(code) == 0 {
      break
    }
  }
  if err != nil {
    return "", err
  } else {
    return decoded, nil
  }
}

func getKey() map[string]rune {
  return map[string]rune{
    "000": 'A',
    "001": 'B',
    "111": 'C',
    "10": 'D',
    "110": 'E',
    "01": 'F',
  }
}

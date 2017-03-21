package scan

import (
    "os"
    "log"
    "fmt"
    "bufio"
)

// Scans the a file and prints out every word ony by one.

func Skan(filename string) {

    // Open file and create scanner on top of it
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    // Default scanner is bufio.ScanLines. Lets use ScanWords.
    // Could also use a custom function of SplitFunc type
    scanner.Split(bufio.ScanWords)


    // Scan for next token.
    success := scanner.Scan()
    if success == false {
        // False on error or EOF. Check error
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        } else {
            log.Fatal(err)
        }
    }

    // Get data from scan with Bytes() or Text()
    fmt.Println("Words scanned:")
    fmt.Println(scanner.Text())

    for i := 0; i < 1; {
      x := scanner.Scan()
      if x == true {

        fmt.Println(scanner.Text())
      } else {
        break
      }

    }


    // Call scanner.Scan() again to find next token
}

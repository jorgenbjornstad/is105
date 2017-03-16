package functions

import ("fmt"
"os"
	"log"
	"io/ioutil"
	"bytes"
)

func Lineshift_iterate() {
	arg := os.Args[1]

	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if var1 := bytes.Contains([]byte(data)), []byte("\x0d\x0a"); var1 == true {
		fmt.Println("This file uses CRLF.")
	}else if var2 := bytes.Contains([]byte(data)), []byte("\x0a"); var2 == true {
		fmt.Println("This file uses LF.")
	}else if var3 := bytes.Contains([]byte(data)), []byte("\x0d"); var3 == true {
			fmt.Println("This file uses CR.")
			}



	fmt.Printf("Data as hex: %x\n", var1)


	fmt.Println("Number of bytes read:", len(data))
}
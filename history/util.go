package history

import (
	"fmt"
	"os"
)

func readLastNLines(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 32)
	n, err := file.ReadAt(buf, fi.Size()-int64(len(buf)))
	if err != nil {
		fmt.Println(err)
	}
	buf = buf[:n]
	fmt.Printf("%s", buf)

}

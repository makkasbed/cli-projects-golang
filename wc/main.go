package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"flag"
)
func main(){
	lines := flag.Bool("l", false,"Count Lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	//set up scanner on reader
	scanner := bufio.NewScanner(r)
	//split based on words
	if !countLines{
		scanner.Split(bufio.ScanWords)
	}

	//split based on bytes
	if countBytes{
		scanner.Split(bufio.ScanBytes)
	}
	
	//initialize word count to 0
	wc := 0
	//scan and increase word count
	for scanner.Scan(){
		wc++
	}
	return wc
}
// Created by Noel Willems for COS350 Asg 1 - Golang Text Cloud.
// Sources:
// Ethan Chen helped me understand map syntax in Go.
// https://golangcode.com/how-to-remove-all-non-alphanumerical-characters-from-a-string/
// https://gobyexample.com/
// https://golang.org/pkg/strings/
// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
// https://golangbyexample.com/read-large-file-word-by-word-go/

package main
 import (
	"fmt"; "os"; "log"; "bufio"; "strings"; "regexp"; "unicode"; "sort"; "strconv";
 )

func main() {
		
	if(len(os.Args) != 4) {
		fmt.Println("Please enter 3 args (inputFileName, excludeFileName, outputFileName)")
		os.Exit(1)
	}

	excludeFile, err := os.Open(os.Args[2])
	textFile, err := os.Open(os.Args[1])
	outputName := os.Args[3]

	if err != nil {
        log.Fatal(err)
	}
	exclMap := make(map[string]int)
	textMap := make(map[string]int)

	// Loading excluded text into "exclMap"
	scanner0 := bufio.NewScanner(excludeFile)
	for scanner0.Scan() {
		word := scanner0.Text()
		exclMap[word] = 1;
	}

	// Loading text into "textMap"
	scanner1 := bufio.NewScanner(textFile)
	scanner1.Split(bufio.ScanWords)
	for scanner1.Scan() {
		word := scanner1.Text()
		word = strings.ToLower(word)
		word = (strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))
		reg, err := regexp.Compile("[^A-Za-z' ]+")
   		 if err != nil {
        	log.Fatal(err)
   		 }
		word = reg.ReplaceAllString(word, "")
		// Sending to method to check if the word is excluded 
		 if(!(excluded(word, exclMap))) {
			if val, ok := textMap[word]; ok {
				textMap[word] = val + 1;
			} else {
				textMap[word] = 1;
			}
		 }
	}

	// Creating a struct that contains key and value
	type kv struct {
		Key string
		Value int
	}
	// Initializing sorted[] to values of textmap
	var sorted []kv
	for k, v := range textMap {
		sorted = append(sorted, kv{k, v})
	}
	// Sorting sorted based on key
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	// Creating array "top50" of size 50 of type kv
	var top50[50] kv
	for i:=0; i < 50; i++ {
		top50[i].Key = sorted[i].Key
		top50[i].Value = sorted[i].Value
	}
	// Alphabetizing ... bubble sort, and not most effective, but it'll do
	var temp kv
	for i:= 0; i < 49; i++ {
		for j:= 0; j < 49; j++ {
			if top50[j].Key > top50[j+1].Key {
				temp = top50[j]
				top50[j] = top50[j+1]
				top50[j+1] = temp
			}
		}
	}
	// Creating sizeFactor: had to convert to float64
	sizeFactor := 1000.0 / (float64(sorted[0].Value) - float64(sorted[49].Value))
	// Creating output file to write HTML to 
	outputFile, err := os.Create(outputName)
	if err != nil {
        fmt.Println(err)
        return
	}
	// HTML
	outputFile.WriteString("<!DOCTYPE html>\n")
	outputFile.WriteString("<html>\n")
	outputFile.WriteString("<body>\n")
	// Produces an ombre effect!
	for i:=0; i<50; i++ {
		sf := (float64(top50[i].Value) * sizeFactor)
		stfa := strconv.FormatFloat(sf, 'f', 6, 64)
		outputFile.WriteString("<span style=\"font-size:")
		outputFile.WriteString(stfa)
		outputFile.WriteString("%; color:rgb(") //, i * 30, "\">")
		outputFile.WriteString(strconv.Itoa(i+146))
		outputFile.WriteString(", ")
		outputFile.WriteString(strconv.Itoa(i+10))
		outputFile.WriteString(", ")
		outputFile.WriteString(strconv.Itoa(i+80))
		outputFile.WriteString("\">")
		outputFile.WriteString(top50[i].Key)
		outputFile.WriteString("               </span>\n")
	}
	outputFile.WriteString("</body>")
	outputFile.WriteString("</html>")
	outputFile.Close()
}
// Func excluded
func excluded(word string, words map[string]int) bool {
	ex := false
	if val, ok:= words[word]; ok {
		words[word] = val - 1
		ex = true
	}
	return ex;
}





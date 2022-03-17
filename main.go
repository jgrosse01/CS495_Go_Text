package main

// necessary imports (go alphabetizes them which is really nifty!)
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// loads the quote generator :)
func main() {
	// initialize random seed
	rand.Seed(time.Now().Unix())
	// load quotes from file
	quotes := readFile()

	// define console reader
	consoleReader := bufio.NewReader(os.Stdin)

	// print welcome
	fmt.Println("Welcome to the Random Quote Generator (RQG)! I hope you have a beautiful time!")
	// fencepost here so we can change the string slightly after
	fmt.Println("Would you like a quote? (Input 'Yes' or 'No'): ")
	for true {
		// read text from the console and replace the newline characters with empty space
		text, _ := consoleReader.ReadString('\n')
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println()

		// switch to manage inputs
		switch strings.ToLower(text) {
		// they say yes
		case "yes", "y":
			fmt.Println(getQuote(quotes) + "\n")
			fmt.Println("Would you like another quote? (Input 'Yes' or 'No'): ")
		// they say no
		case "no", "n":
			fmt.Println("We're sorry to see you leave! Bye bye!")
			os.Exit(0)
		// they say literally anything else INCLUDING saying nothing
		default:
			fmt.Println("Please input 'Yes' or 'No'")
		}
	}
}

// getQuote takes a slice of type string and returns a randomly selected string from that list
func getQuote(quoteList []string) string {
	// store size of the list
	size := len(quoteList)
	// if the list is empty, print an error regarding the slice content and exit the program
	if size == 0 {
		fmt.Println("[ERROR] Quote list was not properly loaded! This is an error of read_file. Exiting program.")
		os.Exit(2)
	}
	// get a random number using the seed defined in main
	var numSelection = rand.Intn(size - 1)

	// return the randomly selected quote from the slice of strings passed in
	return quoteList[numSelection]
}

// function to read the file
func readFile() []string {
	// make a return string slice
	var retList []string
	// make a file and an error as contingency when opening file
	file, err := os.Open("data/quotes")
	// if there is a file error (i.e. if the file cannot be opened)
	if err != nil {
		fmt.Println("[ERROR] Quotes File Not Found or File Permission Set Incorrectly! " +
			"Exiting Program. (This is an error in the file-structure of the project).")
		os.Exit(1)
	}

	// sub function that will deal with errors from closing the file
	// defers file close until this method returns (really cool feature)
	defer func(file *os.File) {
		// get the error value from the closure of the file if there is one
		err := file.Close()
		// if there is then print that
		if err != nil {
			fmt.Println("[ERROR] Failed to close file! This is unexpected. " +
				"Memory will be freed upon exiting program")
		}
	}(file)

	// make a file reader instance for the file that is currently open
	fileScanner := bufio.NewScanner(file)
	// scan the file
	for fileScanner.Scan() {
		// for each line, append the string to the slice for later returning
		retList = append(retList, fileScanner.Text())
	}

	// if the file scanner had an error while running then print that
	if err := fileScanner.Err(); err != nil {
		fmt.Println("[ERROR] File Scanner Errored in the Process of Reading File.")
	}

	// return the slice containing the strings read
	return retList
}

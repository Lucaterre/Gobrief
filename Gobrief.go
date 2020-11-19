package main

import (
	"flag"
	"fmt"
	"github.com/euskadi31/go-tokenizer"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"

	"github.com/bbalet/stopwords"
)

type keyval struct {
    key string
    val int
}



func wordCount(txt_file string) []string{
	wordList := strings.Fields(txt_file)
	counts := make(map[string]int)

	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1 
		}
	}
	var keyvals []keyval
	for k, v := range counts {
		keyvals = append(keyvals, keyval{k, v})
	}
	sort.Slice(keyvals, func(i, j int) bool {
		return keyvals[i].val > keyvals[j].val
	})
	s := report(keyvals)

	return s
}




func report(pairs []keyval) []string {
	topTenwords := make([]string, 10)
	fmt.Println("\nI-/ Top Ten words most uses : \n")
	fmt.Println("Rank                Word                  Frequency")
	fmt.Println("====  =================================   =========")
	for rank := 1; rank <= 10; rank++ {
		word := pairs[rank-1].key
		topTenwords[rank-1] += word
		freq := pairs[rank-1].val
		fmt.Printf("%2d   |     %-20s         |     %5d\n", rank, word, freq)
	}

	return topTenwords
}

func main() {

	textCmd := flag.NewFlagSet("text", flag.ExitOnError)
	webCmd := flag.NewFlagSet("web", flag.ExitOnError)

	if len(os.Args) < 2 {
        fmt.Println("expected 'text' or 'web' subcommands")
        os.Exit(1)
    }
	
	switch os.Args[1]{

	case "text":
		textCmd.Parse(os.Args[2:])
		text_file := os.Args[2]

		fmt.Println("\nYour text file to analysis is : " + text_file + "\n")

		// 1) Read entire file content
		content, err := ioutil.ReadFile(text_file)

		if err != nil {
			log.Fatal(err)
		}

		// Pre-process text to analysis
		// - Convert []byte to string and pass string in lower case
	    text := strings.ToLower(string(content))

		// - suppress eliding french words
		r, _ := regexp.Compile("\\w{1,2}'(\\w)")
		textNoElide := r.ReplaceAllString(text, "$1")

		// - suppress all non-alphabetic characters
		//reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
		//processedString := reg.ReplaceAllString(textNoElide, " ")

	    // - tokenize text
	    t := tokenizer.New()
		tokens := t.Tokenize(textNoElide)
		tokensStr := strings.Join(tokens, ",")

		// remove stops words
	    cleanContent := stopwords.CleanString(tokensStr, "fr", true)

	    // remove characters alone
	    reg, _ := regexp.Compile(" \\w ")
	    textNoCharAlone := reg.ReplaceAllString(cleanContent, " ")


		// 4) display frequency of words in ten rank list and stock top ten words in list
	    s := wordCount(textNoCharAlone)
	    listPosTag := strings.Join(s, ",")

	    cmd := exec.Command("python", "./adj_postagger_util.py", listPosTag, text)
	    out, err := cmd.Output()
	    
	    fmt.Println(string(out))
	    if err != nil {
    		log.Fatal(err)
		}


	case "web":

		webCmd.Parse(os.Args[2:])
		fmt.Println("Ok pour le web")
		os.Exit(3)


	}
	

	//r, _ := regexp.Compile("http[s]?")

	//status := r.MatchString(arg)
	//fmt.Println(status)

	
	

	


	



	// fmt.Println("Your file is : " + arg)

	

	// 2) Convert []byte to string to pass in wordCount() func
	//text := string(content)

	// 3) Clean text action : Remove french stop words from text 
	//cleanContent := stopwords.CleanString(text, "fr", true)

	// 4) Return frequency of words in ten rank list  
	//wordCount(cleanContent)

}


package main

import (
  "fmt"
  "flag"
  "strings"
  "net/http"
  "io/ioutil"
  "regexp"
  "os"
  "strconv"
)

func getAllWords(text string) []string {
  //creates regexp object based on string
  //then returns all matches of regex as array of strings
  words:= regexp.MustCompile("\\w+") // \\w+
  return words.FindAllString(text, -1)
}

func countFrequencyOfWords(words []string) map[string]int {
  //initialize hash table with strings as key, int as value
  wordCounts := make(map[string]int)
  for _, word :=range words {
      wordCounts[word]++
  }
  return wordCounts;
}

func dumpHashTableToFile(url string, fileNum int, wordCountsMap map[string]int) {
  //truncates if found, creates if not, opens write only
  outputFile, err := os.OpenFile("url" + strconv.Itoa(fileNum) + ".txt",
    os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
    0666)
  if err != nil {
    panic(err)
  }
  fmt.Fprintf(outputFile, "url: %v\n", url)
  for word, wordCount :=range wordCountsMap {
    fmt.Fprintf(outputFile, "  %v: %v\n",word, wordCount)
  }
}

func fetchURL(url string, fileNum int) {
  fmt.Println("Now Working on URL: ", url)
  res, err := http.Get(url)
  if err != nil {
    fmt.Println("Error: URL \"" + url + "\"" + " unreachable.")
    return
  }
  body := res.Body
  defer body.Close() //closes down at end of function

  bodyText, err := ioutil.ReadAll(body)
  dumpHashTableToFile(url, fileNum, countFrequencyOfWords(getAllWords(string(bodyText[:]))))

}

func main() {
    fmt.Printf("hello, world\n")

    urlPtr := flag.String("urls", "http://www.google.com",
                          "comma seperated list of urls")
    flag.Parse()
    arrayOfURLs := strings.Split(*urlPtr, ",")

    fmt.Println("All URLs:", arrayOfURLs)

    fileNum := 1
    for _,url := range arrayOfURLs {
      fetchURL(url, fileNum)
      fileNum++
    }
}

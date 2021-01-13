package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the hackerrankInString function below.
func hackerrankInString(s string) string {
    // Can build a trie as well
    hackerrank := "hackerrank"
    sPtr := 0
    hPtr := 0
    for hPtr < len(hackerrank) {
        // Find first instance
        for sPtr < len(s) {
            if s[sPtr] == hackerrank[hPtr] {
                hPtr++
                sPtr++
                break
            }
            sPtr++
            // Can't find
            if (sPtr == len(s)) {
                // fmt.Printf("Can't find %c", hackerrank[hPtr])
                return "NO"
            }
        }
        // Can't find
        if (sPtr == len(s) && hPtr != len(hackerrank)) {
            // fmt.Printf("Can't find %c", hackerrank[hPtr])
            return "NO"
        }
    }
    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        s := readLine(reader)

        result := hackerrankInString(s)

        fmt.Fprintf(writer, "%s\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}


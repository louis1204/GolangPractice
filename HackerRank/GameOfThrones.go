package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the gameOfThrones function below.
func gameOfThrones(s string) string {
    // Only one odd
    count := make(map[string]int)
    for i := 0; i < len(s); i++ {
        if _, ok := count[string(s[i])]; ok {
            count[string(s[i])]++
        } else {
            count[string(s[i])] = 1
        }
    }
    odds := 0
    for _, v := range count {
        if v % 2 != 0 {
            odds++
        }
        if odds > 1 {
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

    s := readLine(reader)

    result := gameOfThrones(s)

    fmt.Fprintf(writer, "%s\n", result)

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


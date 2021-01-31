package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
    map1 := make(map[string]bool)
    for i := 0; i < len(s1); i++ {
        map1[string(s1[i])] = true
    }
    for i := 0; i < len(s2); i++ {
        if _, isOk := map1[string(s2[i])]; isOk {
            return "YES"
        }
    }
    return "NO"
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
        s1 := readLine(reader)

        s2 := readLine(reader)

        result := twoStrings(s1, s2)

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

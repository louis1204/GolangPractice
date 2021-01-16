package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the gemstones function below.
func gemstones(arr []string) int32 {
    count := make(map[byte]int)
    target := len(arr)
    for i := 0; i < len(arr); i++ {
        seen := make(map[byte]bool)
        for j := 0; j < len(arr[i]); j++ {
            if _, alreadySeen := seen[arr[i][j]]; !alreadySeen {
                if _, ok := count[arr[i][j]]; ok {
                    count[arr[i][j]]++
                } else {
                    count[arr[i][j]] = 1
                }
                seen[arr[i][j]] = true
            }
        }
    }
    var res int32 = 0
    for _, value := range count {
        if (value == target) {
            res++
        }
    }
    return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var arr []string

    for i := 0; i < int(n); i++ {
        arrItem := readLine(reader)
        arr = append(arr, arrItem)
    }

    result := gemstones(arr)

    fmt.Fprintf(writer, "%d\n", result)

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


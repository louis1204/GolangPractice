package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
    "reflect"
)

// Complete the funnyString function below.
func funnyString(s string) string {
    ptr1 := 0
    ptr2 := len(s) - 1
    first := []int{}
    second := []int{}
    for ptr1 < len(s) - 1 {
        first = append(first, int(math.Abs(float64(s[ptr1 + 1]) - float64(s[ptr1]))))
        ptr1++
    }
    for ptr2 > 0 {
        second = append(second, int(math.Abs(float64(s[ptr2]) - float64(s[ptr2 - 1]))))
        ptr2--
    }
    if reflect.DeepEqual(first, second) {
        return "Funny"
    } else {
        return "Not Funny"
    }
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

        result := funnyString(s)

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


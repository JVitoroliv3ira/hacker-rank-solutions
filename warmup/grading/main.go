package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func getClosestMultiple(value int32) int32 {
    residue := value % 5
    if residue == 0 {
        return value
    }
    return value + (5 - residue)
}

func getDiffToClosestMultiple(value int32) int32 {
    closestMultiple := getClosestMultiple(value)
    return closestMultiple - value
}

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
    // Write your code here
    
    for i := 0; i < len(grades); i++ {
        grade := grades[i]
        diff := getDiffToClosestMultiple(grade)
        if grade >= 38 && diff < 3 {
            grades[i] = grade + diff  
        }
    }
    
    return grades
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gradesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var grades []int32

    for i := 0; i < int(gradesCount); i++ {
        gradesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        gradesItem := int32(gradesItemTemp)
        grades = append(grades, gradesItem)
    }

    result := gradingStudents(grades)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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

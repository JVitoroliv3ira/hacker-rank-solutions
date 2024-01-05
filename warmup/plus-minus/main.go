package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var positive float32 = 0
    var negative float32 = 0
    var zeros float32 = 0
    quantity := len(arr)
    
    for i := 0; i < quantity; i++ {
        value := arr[i]
        switch {
            case value > 0:
                positive += 1
            case value < 0:
                negative += 1
            default:
                zeros += 1
        }
    }
    
    divider := float32(quantity)
    
    positiveProportion := positive / divider
    negativeProportion := negative / divider
    zerosProportion := zeros / divider
    
    fmt.Printf("%.6f\n", positiveProportion)
    fmt.Printf("%.6f\n", negativeProportion)
    fmt.Printf("%.6f\n", zerosProportion)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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

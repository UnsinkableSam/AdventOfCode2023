package main

import (

    "fmt"
    "strconv"
    "strings"
    "os"
)

func main()  {
     b, err := os.ReadFile("./input.txt")
     if err != nil {
         panic(err)
     }
    input := string(b)
    array := strings.Split(input, "\n")

    resultArray := array
    totalsum := 0
    for i, item := range array {
        if (item == "") {
            break
        }
        fmt.Println("first" + item)
        resultArray[i] = ""
        lettersReverse := strings.Split(reverseSting(item),  "")
        letters := strings.Split(item, "")


        resultArray[i] = findDigit(letters)
        fmt.Println(resultArray[i])
        resultArray[i] = resultArray[i] +  findDigit(lettersReverse)

        fmt.Println(resultArray[i])
        totalsum =  sum(resultArray[i]) + totalsum


    }
        fmt.Println(totalsum)
}

func  findDigit(letters []string) string {

    for _, letter := range letters {
        _, err :=  strconv.Atoi(letter)
        if err ==  nil {
            return letter
        }
    }
    return "";
}

func reverseSting(letters string) string{
    chars := []rune(letters)

    for i, j := 0, len(chars)-1; i < j; i, j= i+1, j-1 {
        chars[i], chars[j] = chars[j], chars[i]
    }
    return string(chars)

}

func sum(strNumber  string) int{
    num, err  := strconv.Atoi(strNumber)
    if err != nil {
        panic(err)
    }
    return num
}

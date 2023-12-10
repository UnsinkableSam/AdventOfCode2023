package main

import (

    "fmt"
    "strconv"
    "strings"
)

func main()  {
    array :=  [4]string{"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet"}
    resultArray := array
    totalsum := 0
    for i , item := range array  {
        resultArray[i] = ""
        lettersReverse := strings.Split(reverseSting(item),  "")
        letters := strings.Split(item, "")

        resultArray[i] = resultArray[i] + findDigit(letters)
        resultArray[i] = resultArray[i] + findDigit(lettersReverse)

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

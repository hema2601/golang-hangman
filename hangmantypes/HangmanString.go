package hangmantypes

import(
    "fmt"
    "strings"
)



type HangmanString struct{
    word string
    revealed [] bool
    l_arr LetterArray
}

func (h_str *HangmanString) Init(str string){
    
    h_str.word = str

    for _, char := range str {
        if ((char <= 'z' && char >= 'a') || (char <= 'Z' && char >= 'A')){
            h_str.revealed = append(h_str.revealed, false)
        } else{
            h_str.revealed = append(h_str.revealed, true)
        }
    }

    h_str.l_arr.Init()
}

func (h_str *HangmanString) ValidateChar(c uint8) uint{
 
    var count uint = 0
    for pos, char := range h_str.word{
        if strings.ToLower(string(c)) == strings.ToLower(string(char)) && h_str.revealed[pos] == false {
            h_str.revealed[pos] = true;
            count++
        }
    }

    if count > 0 {
        h_str.l_arr.Update(c, true)
    }else{
        h_str.l_arr.Update(c, false)
    }

    return count

}


func (h_str *HangmanString) ValidateGuess(str string) uint{

    if len(str) != 1 && len(str) != len(h_str.word){
        return 0
    }

    var count uint = 0

    for _, char := range str{
        count += h_str.ValidateChar(uint8(char))
    }

    return count

}

func (h_str HangmanString)IsRevealed() bool{

    for _, val := range h_str.revealed{
        if val == false{
            return false
        }
    }

    return true
}

func (h_str HangmanString)PrintStr(){
    for pos, char := range h_str.word{
        if h_str.revealed[pos] == true{
            fmt.Print(string(char))
        }else{
            fmt.Print("_")
        }
    }
    fmt.Println("")
}

func (h_str HangmanString)PrintAll(){

    h_str.PrintStr()

    h_str.l_arr.PrintArr()

}

package hangmantypes

import(
    "fmt"
    "errors"
)



type LetterArray struct{
    letters string 
    status [26] uint8
}

func (l_arr *LetterArray)Init(){
    l_arr.letters = "abcdefghijklmnopqrstuvwxyz"
}

func (l_arr *LetterArray)Update(pos uint8, found bool) error{


    if(pos < 'a' || pos > 'z'){
        return errors.New("LetterArray.update: Tried to update with a character not in range [a, z]")
    }

    if found == true{
        l_arr.status[pos-'a'] = 1
    }else{
        l_arr.status[pos-'a'] = 2
    }

    return nil
}

func (l_arr LetterArray)PrintArr(){
    for pos, char := range l_arr.letters{
        if(l_arr.status[pos] == 0){
            fmt.Print(string(char))
        }else{
            fmt.Print("-")
        }
    }

    fmt.Println("")
}

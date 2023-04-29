package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
    "os/exec"
    "time"
    "strconv"
    "errors"
    "unicode"
)



type letter_array struct{
    letters string 
    status [26] uint8
}

func (l_arr *letter_array)init(){
    l_arr.letters = "abcdefghijklmnopqrstuvwxyz"
}

func (l_arr *letter_array)update(pos uint8, found bool) error{


    if(pos < 'a' || pos > 'z'){
        return errors.New("letter_array.update: Tried to update with a character not in range [a, z]")
    }

    if found == true{
        l_arr.status[pos-'a'] = 1
    }else{
        l_arr.status[pos-'a'] = 2
    }

    return nil
}

func (l_arr letter_array)print_arr(){
    for pos, char := range l_arr.letters{
        if(l_arr.status[pos] == 0){
            fmt.Print(string(char))
        }else{
            fmt.Print("-")
        }
    }

    fmt.Println("")
}

type hangman_string struct{
    word string
    revealed [] bool
    l_arr letter_array
}

func (h_str *hangman_string) init(str string){
    
    h_str.word = str

    for _, char := range str {
        if ((char <= 'z' && char >= 'a') || (char <= 'Z' && char >= 'A')){
            h_str.revealed = append(h_str.revealed, false)
        } else{
            h_str.revealed = append(h_str.revealed, true)
        }
    }

    h_str.l_arr.init()
}

func (h_str *hangman_string) validate_char(c uint8) uint{
 
    var count uint = 0
    for pos, char := range h_str.word{
        if strings.ToLower(string(c)) == strings.ToLower(string(char)) && h_str.revealed[pos] == false {
            h_str.revealed[pos] = true;
            count++
        }
    }

    if count > 0 {
        h_str.l_arr.update(c, true)
    }else{
        h_str.l_arr.update(c, false)
    }

    return count

}


func (h_str *hangman_string) validate_guess(str string) uint{

    if len(str) != 1 && len(str) != len(h_str.word){
        return 0
    }

    var count uint = 0

    for _, char := range str{
        count += h_str.validate_char(uint8(char))
    }

    return count

}

func (h_str hangman_string)isRevealed() bool{

    for _, val := range h_str.revealed{
        if val == false{
            return false
        }
    }

    return true
}

func (h_str hangman_string)print_str(){
    for pos, char := range h_str.word{
        if h_str.revealed[pos] == true{
            fmt.Print(string(char))
        }else{
            fmt.Print("_")
        }
    }
    fmt.Println("")
}

func (h_str hangman_string)print_all(){

    h_str.print_str()

    h_str.l_arr.print_arr()

}

type hangman_game struct {
    stage uint8
    symbol [10] string
    log [3] string
    h_str hangman_string
}

func (h_game *hangman_game)init(){
    h_game.stage = 0

    h_game.symbol[0] = "       \n        \n        \n        \n        \n" 
    h_game.symbol[1] = "       \n  |     \n  |     \n  |     \n  |     \n" 
    h_game.symbol[2] = "   __  \n  |     \n  |     \n  |     \n  |     \n" 
    h_game.symbol[3] = "   __  \n  |  |  \n  |     \n  |     \n  |     \n"
    h_game.symbol[4] = "   __  \n  |  |  \n  |  O  \n  |     \n  |     \n"
    h_game.symbol[5] = "   __  \n  |  |  \n  |  O  \n  |  |  \n  |     \n"
    h_game.symbol[6] = "   __  \n  |  |  \n  | \\O  \n  |  |  \n  |     \n"
    h_game.symbol[7] = "   __  \n  |  |  \n  | \\O/ \n  |  |  \n  |     \n"
    h_game.symbol[8] = "   __  \n  |  |  \n  | \\O/ \n  |  |  \n  | /   \n"
    h_game.symbol[9] = "   __  \n  |  |  \n  | \\O/ \n  |  |  \n  | / \\ \n"

    h_game.log[0] = ""
    h_game.log[1] = ""
    h_game.log[2] = ""

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Enter your string: ")

    scanner.Scan()
    
    h_game.h_str.init(scanner.Text())

}




func get_guess () string{

    var letter string
    
    fmt.Print("Your guess (one letter, or whole word):  ")

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()

    letter = scanner.Text()


    return letter
}

func clear(){
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func print_start(){
    fmt.Print("Starting your game")

    i := 0

    for i < 3{
        time.Sleep(time.Millisecond * 400)
        fmt.Print(".")
        i++

    }
       
    time.Sleep(time.Millisecond * 700)

    fmt.Println("");
}

func print_char_by_char(milli time.Duration, str string){

    for _, char := range str{
        fmt.Print(string(char))
        time.Sleep(time.Millisecond * milli)
    }
    fmt.Println("")

}

func scanning_stdout(quit chan bool){
    scanner := bufio.NewScanner(os.Stdout)
    
    for{
        select {
            case <- quit:
                return
            default:
                scanner.Scan()
                fmt.Println("Scanned:", scanner.Text())
        }
    }
}


func (h_game *hangman_game) insert_log(str string){

    

    i := 0

    const max int = 3

    for i < max{
        if i == max-1 { h_game.log[max-1] = str 
        }else          { h_game.log[i] = h_game.log[i+1] }
        i++
    }

}

func (h_game hangman_game) print_state(){
    
    clear()

    fmt.Println("----HANGMAN----")

    fmt.Println(h_game.symbol[h_game.stage])

    fmt.Println("")

    h_game.h_str.print_all()

    fmt.Println("")

    fmt.Println("=============================")
    i := 0
    for i < 3{
        fmt.Println(h_game.log[i])
        i++
    }
    fmt.Println("=============================")

    fmt.Println("")

}

func (h_game *hangman_game) do_turn() {
        
        guess := get_guess()

        if len(guess) != 1 && len(guess) != len(h_game.h_str.word) {
            h_game.insert_log( "Invalid guess! Choose either one letter, or the entire word.")
            return
        }

        if(len(guess) == 1 && !unicode.IsLetter(rune(guess[0]))){
            h_game.insert_log( "Invalid guess! Only alphabetic characters as single guess.")
            return
        }

        count := h_game.h_str.validate_guess(guess)

        if(count == 0){
            h_game.stage++
            h_game.insert_log( "'" + guess + "' is not included...")
            return
        }

        h_game.insert_log("Your guess '" + guess + "' revealed " + strconv.Itoa(int(count)) + " new letter(s)!")     

        return
}

func (h_game *hangman_game) over() bool{
    if h_game.stage == 9{
        h_game.insert_log("Game Over...")
        return true
    }

    if h_game.h_str.isRevealed(){

        h_game.insert_log("Congrats! You guessed the right word!")
        return true
    }

    return false
}

func main(){

    clear()

    var h_game hangman_game 

    h_game.init()

    print_start()

    for !h_game.over()  {
        
        h_game.print_state()

        h_game.do_turn()

    }

    clear()

    h_game.print_state()


}

package hangmantypes

import(
    "fmt"
    "bufio"
    "os"
    "strconv"
    "unicode"
    "hema/hangman/hangmandisplay"
)



type HangmanGame struct {
    stage uint8
    symbol [10] string
    log [3] string
    h_str HangmanString
}

func (h_game *HangmanGame)Init(){
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
    
    h_game.h_str.Init(scanner.Text())

}


func (h_game *HangmanGame) InsertLog(str string){

    

    i := 0

    const max int = 3

    for i < max{
        if i == max-1 { h_game.log[max-1] = str 
        }else          { h_game.log[i] = h_game.log[i+1] }
        i++
    }

}

func (h_game HangmanGame) PrintState(){
    
    hangmandisplay.Clear()

    fmt.Println("----HANGMAN----")

    fmt.Println(h_game.symbol[h_game.stage])

    fmt.Println("")

    h_game.h_str.PrintAll()

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

func (h_game *HangmanGame) DoTurn() {
    
        fmt.Println("Your guess (one letter, or whole word): ")

        scanner := bufio.NewScanner(os.Stdin)

        scanner.Scan()
        
        guess := scanner.Text()

        if len(guess) != 1 && len(guess) != len(h_game.h_str.word) {
            h_game.InsertLog( "Invalid guess! Choose either one letter, or the entire word.")
            return
        }

        if(len(guess) == 1 && !unicode.IsLetter(rune(guess[0]))){
            h_game.InsertLog( "Invalid guess! Only alphabetic characters as single guess.")
            return
        }

        count := h_game.h_str.ValidateGuess(guess)

        if(count == 0){
            h_game.stage++
            h_game.InsertLog( "'" + guess + "' is not included...")
            return
        }

        h_game.InsertLog("Your guess '" + guess + "' revealed " + strconv.Itoa(int(count)) + " new letter(s)!")     

        return
}

func (h_game *HangmanGame) Over() bool{
    if h_game.stage == 9{
        h_game.InsertLog("Game Over...")
        return true
    }

    if h_game.h_str.IsRevealed(){

        h_game.InsertLog("Congrats! You guessed the right word!")
        return true
    }

    return false
}


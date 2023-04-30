package main

import(
    "hema/hangman/hangmantypes"
    "hema/hangman/hangmandisplay"
)


func main(){

    hangmandisplay.Clear()

    var h_game hangmantypes.HangmanGame 

    h_game.Init()

    hangmandisplay.PrintStart()

    for !h_game.Over()  {
        
        h_game.PrintState()

        h_game.DoTurn()

    }

    hangmandisplay.Clear()

    h_game.PrintState()


}

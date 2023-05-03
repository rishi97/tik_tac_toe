//Simple main.go file for Tik Tac Toe
package main

import (
    "fmt"
    "net/http"
    "strconv"
)

// Define the Tic Tac Toe game state
type gameState struct {
    board       [3][3]string
    playerTurn  string
    gameOver    bool
    winner      string
}

// Initialize a new Tic Tac Toe game state
func newGameState() *gameState {
    return &gameState{
        board:      [3][3]string{},
        playerTurn: "X",
        gameOver:   false,
        winner:     "",
    }
}

// Handle requests to the Tic Tac Toe game
func handleTicTacToe(w http.ResponseWriter, r *http.Request) {
    // Load the HTML page for the Tic Tac Toe game
    html := `
        <html>
        <head>
            <title>Tic Tac Toe</title>
            <script>
                function makeMove(row, col) {
                    // Send a request to the server to make a move
                    var xhr = new XMLHttpRequest();
                    xhr.open("POST", "/move?row=" + row + "&col=" + col);
                    xhr.send();
                }
            </script>
        </head>
        <body>
            <table>
                {{range $row := .Board}}
                <tr>
                    {{range $col := $row}}
                    <td><button onclick="makeMove({{$col.Row}}, {{$col.Col}})">{{$col.Value}}</button></td>
                    {{end}}
                </tr>
                {{end}}
            </table>
            <p>{{.Status}}</p>
        </body>
        </html>
    `

    // Parse the HTML template
    tmpl, err := template.New("ticTacToe").Parse(html)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Create a new Tic Tac Toe game state if necessary
    session, err := r.Cookie("session")
    var game *gameState
    if err != nil {
        game = newGameState()
    } else {
        // Load the game state from the session cookie
        b, err := base64.StdEncoding.DecodeString(session.Value)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        err = gob.NewDecoder(bytes.NewReader(b)).Decode(&game)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Check if the game is over
    checkGameOver(game)

    // Render the HTML template with the Tic Tac Toe game state
    err = tmpl.Execute(w, struct{
        Board  [3][3]square
        Status string
    }{
        Board:  convertToSquare(game.board),
        Status: getGameStatus(game),
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Save the game state to the session cookie
    var buf bytes.Buffer
    err = gob.NewEncoder(&buf).Encode(game)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
    http.SetCookie(w, &http.Cookie{
        Name:  "session",
        Value: encoded,
    })
}

// Handle requests to make a move in the Tic Tac Toe game
func handleMove(w http.ResponseWriter, r *http.Request) {
    // Parse the row and column of the move from the request parameters
    row,

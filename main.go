//Code in Go and HTML
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Define the Tic Tac Toe game state
type gameState struct {
	board     [3][3]string
	player    string
	gameOver  bool
	winner    string
}

// Initialize a new Tic Tac Toe game state
func newGameState() *gameState {
	return &gameState{
		board:     [3][3]string{},
		player:    "X",
		gameOver:  false,
		winner:    "",
	}
}

// Check if the Tic Tac Toe game is over
func checkGameOver(game *gameState) {
	// Check rows for a win
	for row := 0; row < 3; row++ {
		if game.board[row][0] != "" && game.board[row][0] == game.board[row][1] && game.board[row][1] == game.board[row][2] {
			game.gameOver = true
			game.winner = game.board[row][0]
			return
		}
	}

	// Check columns for a win
	for col := 0; col < 3; col++ {
		if game.board[0][col] != "" && game.board[0][col] == game.board[1][col] && game.board[1][col] == game.board[2][col] {
			game.gameOver = true
			game.winner = game.board[0][col]
			return
		}
	}

	// Check diagonals for a win
	if game.board[0][0] != "" && game.board[0][0] == game.board[1][1] && game.board[1][1] == game.board[2][2] {
		game.gameOver = true
		game.winner = game.board[0][0]
		return
	}
	if game.board[0][2] != "" && game.board[0][2] == game.board[1][1] && game.board[1][1] == game.board[2][0] {
		game.gameOver = true
		game.winner = game.board[0][2]
		return
	}

	// Check for a tie
	tie := true
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if game.board[row][col] == "" {
				tie = false
				break
			}
		}
	}
	if tie {
		game.gameOver = true
		game.winner = "Tie"
		return
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
					<td><button onclick="makeMove({{$col.Row}}, {{$col.Col}})">{{$col.Value

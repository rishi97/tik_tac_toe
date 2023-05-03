# Tic Tac Toe in Go

This is a simple implementation of Tic Tac Toe in Go, with an HTML interface that allows users to play the game in a web browser.

## Prerequisites
To run this program, you will need to have the following software installed on your machine:

1. Go (version 1.13 or higher)
2. A web browser (e.g. Chrome, Firefox)

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/tic-tac-toe-go.git
```
2. Navigate to the project directory:
```bash
cd tic-tac-toe-go
```
3. Install the required dependencies:
```go
go get github.com/gorilla/mux
```

## Usage
1. Start the Tic Tac Toe game server:
```go
go run main.go
```
2. Open a web browser and navigate to http://localhost:8080.

3. Play the game by clicking on the buttons in the Tic Tac Toe board. Each button corresponds to a row and column on the board, and clicking a button will make a move in that position.

4. The game will end when one player wins, the game is tied, or you close the web browser.

5. To start a new game, simply refresh the web page.

## License

[MIT](https://choosealicense.com/licenses/mit/)

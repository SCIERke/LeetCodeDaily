package main

import "fmt"

var target_word string

func bfs(board [][]byte, x int, y int, wordIndex int, visited [][]bool) bool {
    // If we have matched the entire word, return true
    if wordIndex == len(target_word) {
        return true
    }

    // Check boundaries and if the cell is already visited or doesn't match the current character
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] || board[x][y] != target_word[wordIndex] {
        return false
    }

    // Mark the cell as visited
    visited[x][y] = true

    // Directions for moving up, down, left, right
    dx := [4]int{0, 0, -1, 1}
    dy := [4]int{-1, 1, 0, 0}

    // Explore all 4 directions
    for i := 0; i < 4; i++ {
        new_x := x + dx[i]
        new_y := y + dy[i]

        if bfs(board, new_x, new_y, wordIndex+1, visited) {
            return true
        }
    }

    // Backtrack: unmark the cell as visited
    visited[x][y] = false

    return false
}

func exist(board [][]byte , word string) bool {
		target_word = word;
    // Initialize the visited matrix
    visited := make([][]bool, len(board))
    for i := range visited {
        visited[i] = make([]bool, len(board[0]))
    }

    // Start the search from each cell that matches the first character of the word
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] {
                if bfs(board, i, j, 0, visited) {
                    return true
                }
            }
        }
    }

    return false
}

func main() {
    tmpBoard := [][]byte{
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'},
    }

    word = "SEE"

    fmt.Print(exist(tmpBoard))
}
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution79 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    char[][] board = {
      {'a','a','a'},{'A','A','A'},{'a','a','a'},
    };
    String word = "aAaaaAaaA";
    System.out.println(sol.exist(board, word));
  }
}

class Solution {
  private Set<String> visited;

  public boolean exist(char[][] board, String word) {
      visited = new HashSet<>();
      for (int row = 0; row < board.length; row++) {
          for (int col = 0; col < board[0].length; col++) {
              boolean exists = backtrack(board, word, row, col, 0, new ArrayList<>());
              if (exists) return true;
          }
      }
      return false;
  }

  private boolean backtrack(char[][] board, String word, int currentRow, int currentCol, int wordIdx, List<String> path) {
      int maxRows = board.length;
      int maxCols = board[0].length;
      // if we reached to the end of the word without errors,
      // it is a valid path
      if (wordIdx == word.length()) return true;
      // check if out of boundaries
      // or current word idx doesn't 
      // match in the board
      if (currentRow >= maxRows || 
          currentCol >= maxCols || 
          currentCol < 0 ||
          currentRow < 0 ||
          word.charAt(wordIdx) != board[currentRow][currentCol] ||
          isVisited(currentRow, currentCol)
      ) return false;
      // add current row and col to visited set
      // so we don't visit them again
      createCell(currentRow, currentCol);
      // define allowed directions
      int[][] directions = {
          {0, 1}, // upward
          {0, -1}, // downward
          {1, 0}, // right
          {-1, 0}, // left
      };
      // move to all directions
      for (int[] dir : directions) {
          int newRow = currentRow + dir[0];
          int newCol = currentCol + dir[1];
          boolean exists = backtrack(board, word, newRow, newCol, wordIdx + 1, path);
          if (exists) return true;
      }
      // remove current row and col from visited set
      removeCell(currentRow, currentCol);
      // if path doesn't exist, return false
      return false;
  }

  private boolean isVisited(int row, int col) {
      return visited.contains(getCell(row, col));
  }

  private void createCell(int row, int col) {
      visited.add(getCell(row, col));
  }

  private void removeCell(int row, int col) {
      visited.remove(getCell(row, col));
  }
  
  private String getCell(int row, int col) {
      return row + "," + col;
  }
}
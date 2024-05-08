import { Piece } from "../Piece/Piece";
import { Coordinates, getBoardCoordinates } from "./Grid";

// Define the board state
const boardState = new Map<string, Piece>();

export type BoardState = typeof boardState;

// Helper function to create a map key from coordinates
function coordinatesKey(coordinates: Coordinates): string {
  return `${coordinates.x},${coordinates.y}`;
}

// Example pieces
const piece1: Piece = {
  color: "WHITE",
  type: "KING",
};

const piece2: Piece = {
  color: "BLACK",
  type: "QUEEN",
};

// Adding pieces to the board
boardState.set(coordinatesKey({ x: 4, y: 0 }), piece1);
boardState.set(coordinatesKey({ x: 3, y: 7 }), piece2);

// Function to move a piece
function movePiece(from: Coordinates, to: Coordinates) {
  const piece = boardState.get(coordinatesKey(from));
  if (piece) {
    boardState.delete(coordinatesKey(from));
    boardState.set(coordinatesKey(to), piece);
  }
}

export function getDefaultBoardState(): BoardState {
  const boardState = new Map<string, Piece>();
  const coordinates = getBoardCoordinates(8);
  // coordinates.forEach()
  return boardState;
}

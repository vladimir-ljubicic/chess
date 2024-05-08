import { Coordinates } from "../Board/Grid";
import WhitePawn from "../assets/w_pawn_svg_withShadow.svg";
import BlackPawn from "../assets/b_pawn_svg_withShadow.svg";
import WhiteKnight from "../assets/w_knight_svg_withShadow.svg";
import BlackKnight from "../assets/b_knight_svg_withShadow.svg";
import WhiteBishop from "../assets/w_bishop_svg_withShadow.svg";
import BlackBishop from "../assets/b_bishop_svg_withShadow.svg";
import WhiteRook from "../assets/w_rook_svg_withShadow.svg";
import BlackRook from "../assets/b_rook_svg_withShadow.svg";
import WhiteQueen from "../assets/w_queen_svg_withShadow.svg";
import BlackQueen from "../assets/b_queen_svg_withShadow.svg";
import WhiteKing from "../assets/w_king_svg_withShadow.svg";
import BlackKing from "../assets/b_king_svg_withShadow.svg";

const COLOR = {
  WHITE: "WHITE",
  BLACK: "BLACK",
} as const;

export type Color = keyof typeof COLOR;

const PIECE_TYPE = {
  PAWN: "PAWN",
  KNIGHT: "KNIGHT",
  BISHOP: "BISHOP",
  ROOK: "ROOK",
  QUEEN: "QUEEN",
  KING: "KING",
} as const;

export type PieceType = keyof typeof PIECE_TYPE;

export type Piece = {
  color: Color;
  type: PieceType;
  legalMoves?: Coordinates[];
};

export const pieceToSvgMap: { [key in PieceType]: { [key in Color]: string } } =
  {
    PAWN: { WHITE: WhitePawn, BLACK: BlackPawn },
    KNIGHT: { WHITE: WhiteKnight, BLACK: BlackKnight },
    BISHOP: { WHITE: WhiteBishop, BLACK: BlackBishop },
    ROOK: { WHITE: WhiteRook, BLACK: BlackRook },
    QUEEN: { WHITE: WhiteQueen, BLACK: BlackQueen },
    KING: { WHITE: WhiteKing, BLACK: BlackKing },
  };

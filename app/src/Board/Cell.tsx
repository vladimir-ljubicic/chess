import { Piece, pieceToSvgMap } from "../Piece/Piece";
import "./Board.css";
import { Coordinates } from "./Grid";

export type CellProps = {
  coordinates: Coordinates;
  piece?: Piece;
};

export const Cell: React.FC<CellProps> = ({
  coordinates,
  piece,
}: CellProps) => {
  let svg: string = "";
  if (piece) {
    svg = pieceToSvgMap[piece.type][piece.color];
  }

  return (
    <div className="container">
      <div
        className={
          (coordinates.x + coordinates.y) % 2 ? "cell-white" : "cell-black"
        }
      />
      {svg && <img className="piece-svg" src={svg} alt="Whoops"></img>}
    </div>
  );
};

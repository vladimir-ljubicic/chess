import React from "react";
import { Cell } from "./Cell";
import "./Board.css";
import { getDefaultBoardState } from "./BoardState";

export function Board() {
  const boardState = getDefaultBoardState();
  let cells: React.ReactNode[] = [];

  boardState.forEach((value, key) => {
    const coords = key.split(",");
    cells.push(
      Cell({
        coordinates: { x: parseInt(coords[0]), y: parseInt(coords[1]) },
        piece: value,
      })
    );
  });

  return (
    <div className="grid-container">
      {cells.map((cell, index) => {
        return <div key={index}>{cell}</div>;
      })}
    </div>
  );
}

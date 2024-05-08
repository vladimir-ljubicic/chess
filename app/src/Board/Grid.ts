export type Coordinates = {
  x: number;
  y: number;
};

export function getBoardCoordinates(dimension: number): Coordinates[] {
  let result: Coordinates[] = [];
  let y: number = 0;
  while (y < dimension) {
    let x: number = 0;
    while (x < dimension) {
      result.push({ x, y });
      x++;
    }
    y++;
  }

  return result;
}

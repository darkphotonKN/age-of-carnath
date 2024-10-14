import { GridState, Position } from "./types";
import {
  XDirection,
  YDirection,
  XDirectionEnum,
  YDirectionEnum,
} from "@/constants/enums";

/**
 * Handles highlighting action between the player and the board.
 *
 *  Algorithm from x1, y1 -> x2, y2
 *  1) Check if already at the target, if so just return the same point.
 *  2) Find if the target is on the same row or column, if so we just iterate towards it and
 *     store the coordinates along the way.
 *  3) Iterate *diagonally* checking 1 and 2, storing coordinates along the way.
 *     3 a) Find direction of coordinate:
 *     3 b) Up / Down: if y1 < y2 then its down and vice versa
 *     3 c) Left / Right:  if x1 < x2 then its right and vice versa.
 *     3 d) iterate in direction until on point, row, or col, then check 1 or 2.
 **/
export function highlightPath(
  x2: number,
  y2: number,
  { x: x1, y: y1 }: Position,
  gridState: GridState,
): Position[] {
  if (x2 > gridState.length || y2 >= gridState[0].length) {
    console.error("Index passed for x or y is out of bounds.");
    return [];
  }

  const coordinates: Position[] = [];

  // always start with default
  coordinates.push({ x: x1, y: y1 });

  // 1 -- check if already at target --
  if (x1 == x2 && y1 == y2) {
    coordinates.push({ x: x2, y: y2 });
    return coordinates;
  }

  // 2 -- check if same row or col --

  // --- same column ---
  if (x1 === x2) {
    iterateRows(x2, y1, y2, coordinates);
    return coordinates;
  }

  // --- same row ---
  if (y1 == y2) {
    iterateColumns(y2, x1, x2, coordinates);
    return coordinates;
  }

  // 3 -- traverse diagonally --

  // find direction
  const yDirection: YDirectionEnum = y1 < y2 ? YDirection.DOWN : YDirection.UP; // direction is down if y1 is smaller
  const xDirection: XDirectionEnum =
    x1 < x2 ? XDirection.RIGHT : XDirection.LEFT; // direction is right if x1 is smaller

  // declare starting vars at the start position
  let x3 = x1;
  let y3 = y1;

  // Target: BOTTOM - RIGHT
  if (yDirection === YDirection.DOWN && xDirection === XDirection.RIGHT) {
    do {
      // iterate diagonally
      x3++;
      y3++;

      // add current coordinate
      coordinates.push({ x: x3, y: y3 });

      // check if on the same row or col as target
      if (y3 === y2) {
        // add remaining coordinates
        iterateColumns(y3, x3, x2, coordinates);
        return coordinates;
      }

      if (x3 === x2) {
        // add remaining coordinates
        iterateRows(x3, y3, y2, coordinates);
        return coordinates;
      }
    } while (y3 < y2 && x3 < x2);
  }

  return coordinates;
}

// updates row coordinates via pass by reference
function iterateRows(
  x: number,
  y1: number,
  y2: number,
  coordinates: Position[],
) {
  const direction = y1 < y2 ? YDirection.DOWN : YDirection.UP;
  // --- same column ---
  // iterate towards row
  if (direction === YDirection.DOWN) {
    for (let yi = y1 + 1; yi < y2; yi++) {
      coordinates.push({ y: yi, x });
    }
    return;
  }
  if (direction === YDirection.UP) {
    for (let yi = y1 - 1; yi > y2; yi--) {
      coordinates.push({ y: yi, x });
    }
    return;
  }
}

// updates column coordinates via pass by reference
function iterateColumns(
  y: number,
  x1: number,
  x2: number,
  coordinates: Position[],
) {
  const direction = x1 < x2 ? XDirection.RIGHT : XDirection.LEFT;

  // --- same row ---
  // iterate towards col
  if (direction === XDirection.RIGHT) {
    for (let xi = x1 + 1; xi < x2; xi++) {
      coordinates.push({ y, x: xi });
    }
    return;
  }
  if (direction === XDirection.LEFT) {
    for (let xi = x1 - 1; xi > x2; xi--) {
      coordinates.push({ y, x: xi });
    }
    return;
  }
}

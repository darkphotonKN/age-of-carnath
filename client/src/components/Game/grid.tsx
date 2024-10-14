import { useState, useEffect } from "react";
import GameBlock from "./Block";
import { GridState } from "@/game/types";
import { ContentType } from "@/constants/enums";

/**
 * GameGrid Component
 *
 * Manages the game grid render by keeping a game grid state locally and
 * syncing it with the game server's grid state.
 *
 * Local state grid is used to render things for temporary visual purposes.
 **/
function GameGrid() {
  const [gridState, setGridState] = useState<GridState>();

  useEffect(() => {
    const COL_SIZE = 24;
    const ROW_SIZE = 16;

    const mockGrid: GridState = [];

    // creating mock data
    for (let row = 0; row < ROW_SIZE; row++) {
      // initialize new row
      mockGrid[row] = [];

      for (let col = 0; col < COL_SIZE; col++) {
        mockGrid[row][col] = {
          contentType: ContentType.EMPTY,
          position: { x: col, y: row },
        };
      }
    }

    setGridState(mockGrid);

    // load server grid state into local grid state
  }, []);

  console.log("Mock GridState:", gridState);

  function highlightAction(rowIndex: number, colIndex: number) {
    console.log(`Highlighting index: x (${rowIndex}) y (${colIndex})`);
  }

  // TODO: use grid from game server
  return (
    <div className="mt-5 flex flex-col justify-center items-center">
      {gridState?.map((gridRow, rowIndex) => {
        return (
          <div key={rowIndex} className="flex">
            {gridState?.[rowIndex].map((gridBlock, colIndex) => (
              <GameBlock
                key={rowIndex + " " + colIndex}
                onMouseEnter={() => highlightAction(rowIndex, colIndex)}
              />
            ))}
          </div>
        );
      })}
    </div>
  );
}

export default GameGrid;

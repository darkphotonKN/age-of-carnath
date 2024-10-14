import { useState, useEffect } from "react";
import GameBlock from "./Block";
import { GridState } from "@/game/types";
import { ContentType } from "@/constants/enums";
import { highlightPath } from "@/game/gameLogic";

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

  // NOTE: Testing - Remove after test
  useEffect(() => {
    const testPlayer = {
      name: "Test player",
      position: {
        x: 17,
        y: 0,
      },
    };
    if (!gridState) return;
    const coords = highlightPath(3, 0, testPlayer.position, gridState);
    console.log("[@highlightPath]: Coords:", coords);
  }, [gridState]);

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
    if (!gridState) return;
    console.log(`Highlighting index: x (${rowIndex}) y (${colIndex})`);
    const newGridState = [...gridState];

    newGridState?.forEach((row) => {
      row.forEach((item) => {
        if (item.position.x == colIndex && item.position.y == rowIndex) {
          item.highlight = true;
        }
      });
    });

    setGridState(newGridState);
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
                highlight={gridBlock.highlight}
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

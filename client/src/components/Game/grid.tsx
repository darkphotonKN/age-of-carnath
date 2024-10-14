import { useState, useEffect } from "react";
import GameBlock from "./Block";
import { GridState } from "@/game/types";
import { ContentType } from "@/constants/enums";
import { clearGridStateHighlighting, highlightPath } from "@/game/gameLogic";

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

  useEffect(() => {}, [gridState]);

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

  function highlightPathPreview(rowIndex: number, colIndex: number) {
    if (!gridState) return;

    const newGridState = [...gridState];

    // clear all old highlight state
    clearGridStateHighlighting(newGridState);

    const testPlayer = {
      name: "Test player",
      position: {
        x: 0,
        y: 0,
      },
    };

    // highlights preview of possible paths
    highlightPath(colIndex, rowIndex, testPlayer.position, newGridState);

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
                onMouseEnter={() => highlightPathPreview(rowIndex, colIndex)}
              />
            ))}
          </div>
        );
      })}
    </div>
  );
}

export default GameGrid;

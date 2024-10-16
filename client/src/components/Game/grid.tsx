"use client";
import { useState, useEffect } from "react";
import GameBlock from "./Block";
import { GridBlock, GridState, Player, Position } from "@/game/types";
import { ContentType } from "@/constants/enums";
import { clearGridStateHighlighting, highlightPath } from "@/game/gameLogic";
import Image from "next/image";

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
  const [currentPlayer, setCurrentPlayer] = useState<Player>();

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

    // adding mock player
    const mockPlayer = {
      id: "123",
      name: "Mock player",
    };
    const mockXCoord = 12;
    const mockYCoord = 6;

    mockGrid[mockYCoord][mockXCoord] = {
      contentType: ContentType.PLAYER,
      position: {
        x: mockXCoord,
        y: mockYCoord,
      },
      content: mockPlayer,
    };

    setGridState(mockGrid);

    // TODO: load player's character as the current state
    setCurrentPlayer(mockPlayer);

    // TODO: load server grid state into local grid state
  }, []);

  /**
   * Find the position of the current client's own character player based on ID
   **/
  function findCurrentPlayerPos(): Position | undefined {
    if (!currentPlayer || !gridState) return;

    // search current gridState for their own player

    // searching y (rows)
    for (let y = 0; y < gridState.length; y++) {
      // searching x (columns)
      for (let x = 0; x < gridState[y].length; x++) {
        if (gridState[y][x]?.content?.id === currentPlayer.id) {
          // found player
          return {
            x,
            y,
          };
        }
      }
    }
  }

  function highlightPathPreview(targetY: number, targetX: number) {
    if (!gridState) return;

    const newGridState = [...gridState];

    // clear all old highlight state
    clearGridStateHighlighting(newGridState);

    // get current client's player's position
    const currentPlayerPos = findCurrentPlayerPos();

    console.log("current player pos:", currentPlayerPos);

    if (!currentPlayerPos) return;

    // highlights preview of possible paths
    // NOTE: player position is the current turn player's position
    highlightPath(targetX, targetY, currentPlayerPos, newGridState);

    setGridState(newGridState);
  }

  // TODO: use grid from game server
  return (
    <div className="mt-5 flex flex-col justify-center items-center">
      {gridState?.map((gridRow, y) => {
        return (
          <div key={y} className="flex">
            {gridState?.[y].map((gridBlock, x) => (
              <GameBlock
                key={y + " " + x}
                highlight={gridBlock.highlight}
                contentType={gridBlock?.contentType}
                coords={gridBlock.position}
                onMouseEnter={() => highlightPathPreview(y, x)}
              />
            ))}
          </div>
        );
      })}
    </div>
  );
}

export default GameGrid;

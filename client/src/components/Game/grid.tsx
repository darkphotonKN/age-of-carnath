"use client";
import { useState, useEffect } from "react";
import GameBlock from "./Block";
import { GridState, Player, Position } from "@/game/types";
import { ContentType } from "@/constants/enums";
import { clearGridStateHighlighting, highlightPath } from "@/game/gameLogic";
import { TooltipProps } from "../Tooltip";
import { useWebsocketStore } from "@/stores/websocketStore";

/**
 * GameGrid Component
 *
 * Manages the game grid render by keeping a game grid state locally and
 * syncing it with the game server's grid state.
 *
 * Local state grid is used to render things for temporary visual purposes.
 **/
function GameGrid() {
  const { gameState } = useWebsocketStore();
  const [gridState, setGridState] = useState<GridState>();
  const [currentPlayer, setCurrentPlayer] = useState<Player>();
  const [currentTarget, setCurrentTarget] = useState<TooltipProps>();

  useEffect(() => {
    const { gridState, players } = gameState || {};
    console.log("@WS gridState from server:", gameState);
    console.log("@WS players in match from server:", gameState);

    // TODO: Make grid state conform to the frontend version.
    setGridState(gridState);

    //TODO: Update to real authenticated player.
    const player = {
      id: "5f77878d-a770-4729-b8d2-90ac1b6296d3",
      name: "Next Client Player",
    };

    setCurrentPlayer(player);
  }, []);

  console.log("@WS gridState client:", gridState);

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
        if (gridState[y][x]?.content?.player?.id === currentPlayer.id) {
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

    // set current target, for Tooltip and target tracking
    setCurrentTarget({ position: { x: targetX, y: targetY } });

    // highlights preview of possible paths
    // NOTE: player position is the current turn player's position
    highlightPath(targetX, targetY, currentPlayerPos, newGridState);

    setGridState(newGridState);
  }

  // TODO: use grid from game server
  return (
    <div className="border border-customBorderGray overflow-hidden mt-5 flex flex-col justify-center items-center">
      {gridState?.map((gridRow, y) => {
        return (
          <div key={y} className="flex">
            {gridState?.[y].map((gridBlock, x) => (
              <GameBlock
                key={y + " " + x}
                highlight={gridBlock.highlight}
                contentType={gridBlock?.contentType}
                coords={gridBlock.position}
                tooltipProps={currentTarget}
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

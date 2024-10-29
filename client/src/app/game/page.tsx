"use client";
import GameGrid from "@/components/Game/Grid";
import GameOptions from "@/components/Game/Options";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useEffect } from "react";

export default function Game() {
  // const { closeConnection } = useWebsocketStore();

  useEffect(() => {
    // TODO: clean up when leaving game in any way
    return () => {
      // closeConnection();
    };
  }, []);

  return (
    <div className="h-full">
      {/* -- Main Game View -- */}
      <GameGrid />

      {/* -- Options Menu -- */}
      <GameOptions />
    </div>
  );
}

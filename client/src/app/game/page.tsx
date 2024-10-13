"use client";
import GameGrid from "@/components/Game/Grid";
import GameOptions from "@/components/Game/Options";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useEffect } from "react";

export default function Game() {
  // connect to websocket state store
  const { ws } = useWebsocketStore();

  useEffect(() => {
    console.log("[@Game page] ws:", ws);
    // clean up when leaving game in any way
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

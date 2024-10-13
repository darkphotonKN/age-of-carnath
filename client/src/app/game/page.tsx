"use client";
import GameGrid from "@/components/Game/grid";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useEffect } from "react";

export default function Game() {
  // connect to websocket state store
  const { ws, closeConnection } = useWebsocketStore();

  useEffect(() => {
    console.log("[@Game page] ws:", ws);
    // clean up when leaving game in any way
    return () => {
      if (ws) {
        console.log(
          "[@Game page] Cleaning up connection due to leaving the page or dismount",
        );
        closeConnection();
      }
    };
  }, []);

  return (
    <div className="h-full">
      <GameGrid />
    </div>
  );
}

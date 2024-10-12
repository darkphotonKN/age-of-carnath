"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useWebsocketStore } from "@/stores/websocketStore";

export default function MainMenu() {
  const router = useRouter();

  // connect to websocket state store
  const { ws, setupWebSocket, startMatchmaking } = useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  function handleFindMatch() {
    setupWebSocket();
  }

  useEffect(() => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      // route to the match page
      router.push("/game");

      // init matchmaking
      const player = {
        id: uuidv4(),
        name: "test first ever player",
      };

      startMatchmaking(player);
    }
  }, [ws, ws?.readyState]);

  return (
    <div className="flex flex-col justify-center items-center h-full">
      <Button variant="default" size="default" onClick={handleFindMatch}>
        Find Match
      </Button>
    </div>
  );
}

"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useState } from "react";

export default function MainMenu() {
  const [matchStart, setMatchStart] = useState(false);
  // connect to websocket state store
  const { ws, setupWebSocket, startMatchmaking, findingMatch } =
    useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  function handleInitFindMatch() {
    setupWebSocket();
  }

  if (!matchStart && ws && ws.readyState === WebSocket.OPEN) {
    const id = uuidv4();
    console.log("re-initializing id:", id);
    // init matchmaking
    const player = {
      id,
      name: "test first ever player",
    };

    startMatchmaking(player);

    setMatchStart(true);
  }

  return (
    <div className="flex flex-col justify-center items-center h-full">
      {findingMatch ? (
        <div>Searching for a Match...</div>
      ) : (
        <Button variant="default" size="default" onClick={handleInitFindMatch}>
          Find Match
        </Button>
      )}
    </div>
  );
}

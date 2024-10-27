"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useWebsocketStore } from "@/stores/websocketStore";

export default function MainMenu() {
  // connect to websocket state store
  const { ws, setupWebSocket, startMatchmaking, findingMatch } =
    useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  function handleInitFindMatch() {
    setupWebSocket();
  }

  if (ws && ws.readyState === WebSocket.OPEN) {
    // init matchmaking
    const player = {
      id: uuidv4(),
      name: "test first ever player",
    };

    startMatchmaking(player);
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

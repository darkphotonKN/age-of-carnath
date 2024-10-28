"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useEffect, useState } from "react";

export default function MainMenu() {
  const [matchStart, setMatchStart] = useState(false);

  // connect to websocket state store
  const { ws, setupWebSocket, startMatchmaking, findingMatch, isConnected } =
    useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  useEffect(() => {
    // NOTE: useEffect is warranted here alright u snoopers, useEffect
    // is designed for interacting with EXTERNAL systems - in this case our
    // websocket server.

    if (matchStart && isConnected) {
      const id = uuidv4();
      // init matchmaking
      // TODO: remove test player
      const player = {
        id,
        name: "test first ever player",
      };
      startMatchmaking(player);
    }
  }, [matchStart, startMatchmaking, isConnected]);

  function handleInitFindMatch() {
    if (!ws) {
      setupWebSocket();
    }

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

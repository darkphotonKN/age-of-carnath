"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useWebsocketStore } from "@/stores/websocketStore";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function MainMenu() {
  const router = useRouter();
  const [matchStart, setMatchStart] = useState(false);

  // connect to websocket state store
  const {
    ws,
    setupWebSocket,
    startMatchmaking,
    findingMatch,
    isConnected,
    matchInitiated,
  } = useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  useEffect(() => {
    console.log(`@WS matchStart: ${matchStart} isConnected: ${isConnected}`);
    if (matchStart && isConnected) {
      const id = uuidv4();
      // init matchmaking
      // TODO: Add actual authenticated player.
      const player = {
        id,
        name: "test first ever player",
      };

      startMatchmaking(player);
    }
  }, [matchStart, startMatchmaking, isConnected]);

  useEffect(() => {
    // route to game page after game successfully inits.
    if (matchInitiated) {
      router.push("/game");
    }
  }, [router, matchInitiated]);

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

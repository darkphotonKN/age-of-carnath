"use client";
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

  // -- Handle Finding a Match --
  useEffect(() => {
    console.log(`@WS matchStart: ${matchStart} isConnected: ${isConnected}`);
    if (matchStart && isConnected) {
      // init matchmaking
      // TODO: Add actual authenticated player.
      const player = {
        id: "5f77878d-a770-4729-b8d2-90ac1b6296d3",
        name: "test first ever player",
      };

      startMatchmaking(player);
    }
  }, [matchStart, isConnected]);

  // route to game page after game successfully inits.
  useEffect(() => {
    if (matchInitiated) {
      router.push("/game");
    }
  }, [matchInitiated]);

  function handleInitFindMatch() {
    if (!ws) {
      setupWebSocket();
    }

    setMatchStart(true);
  }

  console.log("@WS findingMatch:", findingMatch);

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

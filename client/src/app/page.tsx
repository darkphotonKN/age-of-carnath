"use client";
import { Button } from "@/components/Button";
import { GameAction } from "@/constants/enums";
import { GamePayload, Player } from "@/game/types";
import useWebSocketServer from "@/hooks/useWebsocketServer";
import { useEffect, useState } from "react";

export default function Home() {
  const [connect, setConnect] = useState(false);
  // connect to websocket
  const { ws } = useWebSocketServer({
    connectToWebSocket: connect,
    playerId: "1",
    gameId: "1",
  });

  // -- Handle Finding a Match --
  function handleFindMatch() {
    setConnect(true);
  }
  useEffect(() => {
    if (ws) {
      const messagePayload: GamePayload<Player> = {
        action: GameAction.FIND_MATCH,
        payload: {
          id: "123",
          name: "test first ever player",
        },
      };
      // start matchmaking
      setTimeout(() => {
        ws.send(JSON.stringify(messagePayload));
      }, 2000);
    }
  }, [ws]);

  return (
    <div className="flex flex-col justify-center content-center h-full">
      <Button variant="default" size="default" onClick={handleFindMatch}>
        Find Match
      </Button>
    </div>
  );
}

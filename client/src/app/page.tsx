"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { GameAction } from "@/constants/enums";
import { GamePayload, Player } from "@/game/types";
import useWebSocketServer from "@/hooks/useWebsocketServer";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function MainMenu() {
  const router = useRouter();

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
    if (ws && ws.readyState === WebSocket.OPEN) {
      const messagePayload: GamePayload<Player> = {
        action: GameAction.FIND_MATCH,
        payload: {
          id: uuidv4(),
          name: "test first ever player",
        },
      };
      // start matchmaking
      ws.send(JSON.stringify(messagePayload));

      // route to the match page
      router.push("/game");
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

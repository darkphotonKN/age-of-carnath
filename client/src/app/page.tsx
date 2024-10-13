"use client";
import { v4 as uuidv4 } from "uuid";
import { Button } from "@/components/Button";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useWebsocketStore } from "@/stores/websocketStore";

export default function MainMenu() {
  const router = useRouter();

  // for handling init connection
  const [initFindMatch, setFindMatch] = useState(false);

  // connect to websocket state store
  const { ws, setupWebSocket, startMatchmaking } = useWebsocketStore();

  // -- Handle Finding a Match function and useEffect --
  function handleInitFindMatch() {
    setupWebSocket();
    setFindMatch(true);
  }

  useEffect(() => {
    console.log("[@Home page] ws:", ws);
    if (initFindMatch && ws && ws.readyState === WebSocket.OPEN) {
      // init matchmaking
      const player = {
        id: uuidv4(),
        name: "test first ever player",
      };

      startMatchmaking(player);

      // route to the match page
      router.push("/game");
    }
  }, [ws, ws?.readyState]);

  return (
    <div className="flex flex-col justify-center items-center h-full">
      <Button variant="default" size="default" onClick={handleInitFindMatch}>
        Find Match
      </Button>
    </div>
  );
}

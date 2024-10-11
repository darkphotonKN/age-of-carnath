import { useEffect, useState } from "react";

type WebSocketOptions = {
  connectToWebSocket: boolean;
  playerId?: string;
  gameId?: string;
};

export default function useWebSocketServer({
  connectToWebSocket,
  playerId,
  gameId,
}: WebSocketOptions) {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [systemMsg, setSystemMsg] = useState("");
  const [systemMsgPopup, setSystemMsgPopup] = useState(false);

  // -- Setup WebSocket --
  useEffect(() => {
    if (!connectToWebSocket) return;

    // Connect to WebSocket server (adjust URL for your actual server)
    const socket = new WebSocket(
      `ws://localhost:4111/ws?playerId=${playerId}&gameId=${gameId}`,
    );

    socket.onopen = () => {
      console.log("Connected to WebSocket server!");
    };

    socket.onerror = (error) => {
      console.log("WebSocket error:", error);
    };

    socket.onmessage = (event) => {
      // TODO: remove after testing
      console.log(event.data);

      if (typeof event.data === "string") {
        try {
          const message = JSON.parse(event.data);
          handleIncomingMessage(message);
        } catch (error) {
          console.log("Failed to parse incoming message:", error);
        }
      }
    };

    socket.onclose = () => {
      console.log("Disconnected from WebSocket server.");
    };

    // Set WebSocket to state for persistent usage
    setWs(socket);

    // Cleanup when component unmounts
    return () => {
      socket.close(1000, "Client disconnected");
    };
  }, [connectToWebSocket, gameId, playerId]);

  // Handle incoming WebSocket messages
  function handleIncomingMessage(message: {
    action: string;
    payload: unknown;
  }) {
    switch (message.action) {
      case "player_joined":
        const playerJoinedPayload = message.payload as { name: string };

        setSystemMsgPopup(true);
        setSystemMsg(`${playerJoinedPayload.name} has joined the game.`);
        break;
      default:
        console.log("Unhandled message action:", message.action);
    }
  }

  // -- System Message Handling --
  useEffect(() => {
    if (systemMsg) {
      setTimeout(() => {
        setSystemMsgPopup(false);
      }, 2000);

      setTimeout(() => {
        setSystemMsg("");
      }, 2200);
    }
  }, [systemMsg]);

  return { ws, systemMsg, systemMsgPopup };
}

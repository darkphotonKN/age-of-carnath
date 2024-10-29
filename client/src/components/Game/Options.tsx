import { useWebsocketStore } from "@/stores/websocketStore";
import { Button } from "../Button";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

function GameOptions() {
  // connect to websocket state store
  const { ws, isConnected, closeConnection } = useWebsocketStore();

  const router = useRouter();

  function handleCloseConnection() {
    // only close conneiton if ws is even instantiated
    console.log(
      "[@Game page] Cleaning up connection due to leaving the page or dismount",
    );
    closeConnection();
  }

  useEffect(() => {
    // only route back to home when both isConnected is set to false and websocket connection has been set to null
    if (!isConnected && !ws) {
      router.push("/");
    }
  }, [isConnected, ws, router]);

  return (
    <div className="flex justify-center mt-5 gap-3">
      <Button variant="default">Settings</Button>
      <Button variant="default" onClick={handleCloseConnection}>
        Exit Match
      </Button>
    </div>
  );
}

export default GameOptions;

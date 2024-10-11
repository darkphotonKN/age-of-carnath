import { GameActionEnum } from "@/constants/enums";

// controls game payloads between client and game server.
export type GamePayload<T> = {
  action: GameActionEnum;
  payload: T;
};

export type Player = {
  id: string;
  name: string;
};

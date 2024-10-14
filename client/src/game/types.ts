import { ContentTypeEnum, GameActionEnum } from "@/constants/enums";

// controls game payloads between client and game server.
export type GamePayload<T> = {
  action: GameActionEnum;
  payload: T;
};

export type Player = {
  id: string;
  name: string;
};

export type Position = {
  x: number;
  y: number;
};

export type GridBlock = {
  contentType: ContentTypeEnum;
  position: Position;
};

export type GridState = GridBlock[][];

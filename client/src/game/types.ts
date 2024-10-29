import { ContentTypeEnum, GameActionEnum } from "@/constants/enums";

// controls game payloads between client and game server.
export type GamePayload<T> = {
  action: GameActionEnum;
  payload: T;
};

// general game information, matches server's
export type GameState = {
  gridState: GridBlock[][];
  players: Player[];
};

export type Player = {
  id: string;
  name: string;
};

export type Item = {
  id: string;
  label: string;
};

export type Content = {
  player?: Player;
  item?: Item;
};

export type Position = {
  x: number;
  y: number;
};

export type GridBlock = {
  contentType: ContentTypeEnum;
  position: Position;
  content?: Content;
  highlight?: boolean;
};

export type GridState = GridBlock[][];

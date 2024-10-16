type EnumValue<T> = T[keyof T];

export const GameAction = {
  FIND_MATCH: "find_match",
  MOVE: "move",
  ATTACK: "attack",
} as const;
export type GameActionEnum = EnumValue<typeof GameAction>;

export const ContentType = {
  EMPTY: "empty",
  PLAYER: "player",
  ITEM: "item",
} as const;
export type ContentTypeEnum = EnumValue<typeof ContentType>;

export const YDirection = {
  UP: "up",
  DOWN: "down",
} as const;
export type YDirectionEnum = EnumValue<typeof YDirection>;

export const XDirection = {
  LEFT: "left",
  RIGHT: "right",
} as const;
export type XDirectionEnum = EnumValue<typeof XDirection>;

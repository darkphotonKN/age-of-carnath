type EnumValue<T> = T[keyof T];

export const GameAction = {
  FIND_MATCH: "find_match",
  MOVE: "move",
  ATTACK: "attack",
} as const;
export type GameActionEnum = EnumValue<typeof GameAction>;

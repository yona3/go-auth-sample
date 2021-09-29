import { atom } from "recoil";
import type { User } from "src/model";

export const accessTokenState = atom<string | null>({
  key: "accessTokenState",
  default: null,
});

export const meState = atom<User | null>({
  key: "meState",
  default: null,
});

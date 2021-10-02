import type { User } from "src/model";

type WouldBe<T> = { [P in keyof T]: T[P] };

const isObj = <T>(obj: unknown): obj is WouldBe<T> =>
  typeof obj === "object" && obj !== null;

export const isMe = (me: unknown): me is User =>
  isObj<User>(me) &&
  typeof me.id === "number" &&
  typeof me.uuid === "string" &&
  typeof me.name === "string" &&
  typeof me.email === "string" &&
  (me.password !== undefined ? typeof me.password === "string" : true) &&
  typeof me.createdAt === "string" &&
  typeof me.updatedAt === "string" &&
  typeof me.loggedOutAt === "string";

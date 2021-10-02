export type User = {
  id: number;
  uuid: string;
  name: string;
  email: string;
  password?: string;
  signinWith: "email" | "google" | "twitter";
  createdAt: string;
  updatedAt: string;
  loggedOutAt: string;
};

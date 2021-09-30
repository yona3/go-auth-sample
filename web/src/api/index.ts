import { apiPaths } from "src/utils/apiPaths";

const fetcher = (url: string, init?: RequestInit) =>
  fetch(url, {
    ...init,
    mode: "cors",
  });

export const getIndex = () =>
  fetcher(apiPaths.root(), {
    method: "GET",
  });

export const fetchAccessToken = () =>
  fetcher(apiPaths.token(), {
    method: "POST",
    credentials: "include",
  });

export const fetchMe = (accessToken: string) =>
  fetcher(apiPaths.users.me(), {
    method: "GET",
    headers: {
      // eslint-disable-next-line @typescript-eslint/naming-convention
      Authorization: `Bearer ${accessToken}`,
    },
  });

export const signInWithGoogle = () =>
  fetcher(apiPaths.google.oauth2(), { method: "GET" });

export const signOut = (accessToken: string) =>
  fetcher(apiPaths.refreshToken(), {
    method: "DELETE",
    headers: {
      // eslint-disable-next-line @typescript-eslint/naming-convention
      Authorization: `Bearer ${accessToken}`,
    },
    credentials: "include",
  });

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

export const signInWithGoogle = () =>
  fetcher(apiPaths.google.oauth2(), { method: "GET" });

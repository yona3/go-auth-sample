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

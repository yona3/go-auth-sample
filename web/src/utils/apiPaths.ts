const root = "http://localhost:8080";

export const apiPaths = {
  root: () => root,
  google: {
    root: () => `${root}/google`,
    oauth2: () => `${root}/google/oauth2`,
  },
};

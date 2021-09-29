import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { fetchAccessToken } from "src/api";
import { accessTokenState } from "src/state";

export const useAccessToken = () => {
  const [accessToken, setAccessToken] = useRecoilState(accessTokenState);

  const handleRevokeAccessToken = () => {
    setAccessToken(null);
  };

  useEffect(() => {
    if (accessToken) return;

    // fetch access token
    (async () => {
      try {
        const res = await fetchAccessToken();
        const data = await res.json();

        if (!data.ok || !data.access_token)
          throw new Error("Failed to fetch access token");

        setAccessToken(data.access_token);
      } catch (err) {
        console.error(err);
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return { accessToken, handleRevokeAccessToken };
};

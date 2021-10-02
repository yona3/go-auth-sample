import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { fetchMe } from "src/api";
import { meState } from "src/state";
import { toCamelCaseObjectKey } from "src/utils/toCamelCase";
import { isMe } from "src/utils/typeGuard";

import { useAccessToken } from "./useAccessToken";

export const useMe = () => {
  const [me, setMe] = useRecoilState(meState);
  const { accessToken } = useAccessToken();

  // set me on mount
  useEffect(() => {
    if (!accessToken) return setMe(null); // logout
    if (me) return;

    // fetch me
    (async () => {
      const response = await fetchMe(accessToken);
      const data = await response.json();

      const me = toCamelCaseObjectKey(data);
      if (!isMe(me)) throw new Error("me doc is invalid");

      console.log("me: ", me);
      setMe(me);
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [accessToken]);

  return { me };
};

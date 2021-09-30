import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { fetchMe } from "src/api";
import { meState } from "src/state";

import { useAccessToken } from "./useAccessToken";

export const useMe = () => {
  const [me, setMe] = useRecoilState(meState);
  const { accessToken } = useAccessToken();

  // todo: add handle logout

  // set me on mount
  useEffect(() => {
    if (!accessToken) return setMe(null);
    if (me) return;

    // fetch me
    (async () => {
      const response = await fetchMe(accessToken);
      const data = await response.json();

      // todo: convert snack_case to camelCase
      const me = {
        ...data,
        signinWith: data.signin_with,
        createdAt: data.created_at,
        updatedAt: data.updated_at,
        loggedOutAt: data.logged_out_at,
      };

      delete me.signin_with;
      delete me.created_at;
      delete me.updated_at;
      delete me.logged_out_at;

      // todo: fix any
      setMe(me);
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [accessToken]);

  return { me };
};

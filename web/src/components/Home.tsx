import type { VFC } from "react";
import { signOut } from "src/api";
import { useAccessToken } from "src/hooks/useAccessToken";
import type { User } from "src/model";

import { Button } from "./uiParts/Button";

type Props = {
  me: User;
};

export const Home: VFC<Props> = ({ me }) => {
  const { accessToken, handleRevokeAccessToken } = useAccessToken();

  const handleSignOut = async () => {
    try {
      if (!accessToken) throw new Error("accessToken is not null");

      const res = await signOut(accessToken);
      const data = await res.json();

      if (!data.ok) throw new Error(data.message);

      handleRevokeAccessToken();
      console.log(data.message);
      console.log("Logged out at:", data.logged_out_at);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="pt-12 text-center">
      <h1 className="font-mono text-2xl">
        Hi, {me.name}!
        <span role="img" aria-label="hi">
          👋
        </span>
      </h1>
      <p className="mt-4">{me.email}</p>
      <div className="mt-6">
        <Button className="px-4 text-sm" onClick={handleSignOut}>
          Logout
        </Button>
      </div>
    </div>
  );
};

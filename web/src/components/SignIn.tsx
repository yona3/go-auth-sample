import { useRouter } from "next/dist/client/router";
import type { VFC } from "react";
import { signInWithGoogle } from "src/api";

import { Button } from "./uiParts/Button";

export const SignIn: VFC = () => {
  const router = useRouter();

  const handleSignin = async () => {
    const res = await signInWithGoogle();
    const data = await res.json();
    router.push(data.url);
  };

  return (
    <div className="pt-12 text-center">
      <h1 className="font-mono text-2xl">Sign In</h1>
      <div className="pt-12">
        <div>
          <Button onClick={handleSignin}>Google</Button>
        </div>
      </div>
    </div>
  );
};

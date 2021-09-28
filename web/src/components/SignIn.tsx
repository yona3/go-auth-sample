import { useRouter } from "next/dist/client/router";
import type { VFC } from "react";
import { useEffect } from "react";
import { signInWithGoogle } from "src/api";
import { useError } from "src/hooks/useError";

import { Button } from "./uiParts/Button";

export const SignIn: VFC = () => {
  const { error } = useError();
  const router = useRouter();

  useEffect(() => {
    if (error) alert(error);
  }, [error]);

  const handleSignin = async () => {
    try {
      const res = await signInWithGoogle();
      const data = await res.json();

      if (!data.ok) throw new Error("failed to signin");
      router.push(data.url);
    } catch (err) {
      console.error(err);
    }
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

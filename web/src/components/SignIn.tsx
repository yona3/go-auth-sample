import type { VFC } from "react";

import { Button } from "./uiParts/Button";

export const SignIn: VFC = () => {
  return (
    <div className="pt-12 text-center">
      <h1 className="font-mono text-2xl">Sign In</h1>
      <div className="pt-12">
        <div>
          <Button>Google</Button>
        </div>
      </div>
    </div>
  );
};

import type { VFC } from "react";

export const Home: VFC = () => {
  return (
    <div className="pt-12 text-center">
      <h1 className="font-mono text-2xl">
        Hi, yona!
        <span role="img" aria-label="hi">
          👋
        </span>
      </h1>
    </div>
  );
};

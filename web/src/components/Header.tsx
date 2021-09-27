import type { VFC } from "react";

export const Header: VFC = () => {
  return (
    <header className="py-4 text-center bg-gray-700">
      <h1 className="text-xl font-semibold text-white">Go Auth Sample</h1>
    </header>
  );
};

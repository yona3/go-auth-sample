import type { ReactNode, VFC } from "react";

import { Header } from "./Header";

type Props = {
  children: ReactNode;
};

export const Layout: VFC<Props> = ({ children }) => {
  return (
    <div className="min-h-screen text-white bg-gray-800">
      <Header />
      <div className="px-4">{children}</div>
    </div>
  );
};

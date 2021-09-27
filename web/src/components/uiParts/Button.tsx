import type { VFC } from "react";

type Props = {
  children: React.ReactNode;
};

export const Button: VFC<Props> = ({ children }) => {
  return (
    <button
      className="
        py-2 w-52 text-lg hover:bg-gray-700 
        rounded border transition
      "
    >
      {children}
    </button>
  );
};

import type { ReactNode, VFC } from "react";

type Props = {
  children: ReactNode;
  onClick?: () => Promise<void>;
};

export const Button: VFC<Props> = ({ children, onClick }) => {
  const handleSignIn = onClick;

  return (
    <button
      className="
        py-2 w-52 text-lg hover:bg-gray-700 
        rounded border transition
      "
      onClick={handleSignIn}
    >
      {children}
    </button>
  );
};

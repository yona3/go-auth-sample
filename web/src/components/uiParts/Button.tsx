import type { ReactNode, VFC } from "react";

type Props = {
  children: ReactNode;
  className?: string;
  onClick?: () => Promise<void>;
};

export const Button: VFC<Props> = ({ children, className, onClick }) => {
  const handleOnClick = onClick;

  return (
    <button
      className={`
        py-2 hover:bg-gray-700 
        rounded border transition ${className}
      `}
      onClick={handleOnClick}
    >
      {children}
    </button>
  );
};

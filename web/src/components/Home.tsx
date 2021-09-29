import type { VFC } from "react";
import type { User } from "src/model";

type Props = {
  me: User;
}

export const Home: VFC<Props> = ({ me }) => {
  console.log(me);
  
  return (
    <div className="pt-12 text-center">
      <h1 className="font-mono text-2xl">
        Hi, {me.name}!
        <span role="img" aria-label="hi">
          👋
        </span>
      </h1>
      <p className="mt-2">{me.email}</p>
    </div>
  );
};

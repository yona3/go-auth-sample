import "tailwindcss/tailwind.css";

import type { AppProps } from "next/dist/shared/lib/router/router";
import { RecoilRoot } from "recoil";

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <RecoilRoot>
      <Component {...pageProps} />
    </RecoilRoot>
  );
};

export default App;

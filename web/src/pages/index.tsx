import type { NextPage } from "next";
import Head from "next/head";
import { useEffect } from "react";
import { getIndex } from "src/api";
// import { Home } from "src/components/Home";
import { SignIn } from "src/components/SignIn";

import { Layout } from "../components/Layout";

const Index: NextPage = () => {
  useEffect(() => {
    (async () => {
      try {
        const res = await getIndex();
        const data = await res.json();

        console.log(data);
      } catch (err) {
        console.error(err);
      }
    })();
  }, []);

  return (
    <Layout>
      <Head>
        <title>Home</title>
        <meta
          name="description"
          content="This is my Next.js + TypeScript + Tailwind CSS starter template."
        />
      </Head>

      <SignIn />
      {/* <Home /> */}
    </Layout>
  );
};

export default Index;

import type { NextPage } from "next";
import Head from "next/head";
import { SignIn } from "src/components/SignIn";

import { Layout } from "../components/Layout";

const Index: NextPage = () => {
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
    </Layout>
  );
};

export default Index;

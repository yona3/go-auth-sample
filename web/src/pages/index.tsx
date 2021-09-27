import type { NextPage } from "next";
import Head from "next/head";

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

      <div className="pt-10 text-center">Hello, world!</div>
    </Layout>
  );
};

export default Index;

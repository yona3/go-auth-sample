import type { NextPage } from "next";
import Head from "next/head";
import Link from "next/link";
import { useEffect } from "react";
import { getIndex } from "src/api";
import { Home } from "src/components/Home";
import { SignIn } from "src/components/SignIn";
import { useMe } from "src/hooks/useMe";

import { Layout } from "../components/Layout";

const Index: NextPage = () => {
  const { me } = useMe();

  // fetch access token
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

      {me ? <Home me={me} /> : <SignIn />}

      <div className="mt-12 text-center">
        <Link href="/about">
          <a className="text-blue-400 underline">About</a>
        </Link>
      </div>
    </Layout>
  );
};

export default Index;

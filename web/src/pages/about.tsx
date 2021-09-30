import type { NextPage } from "next";
import Head from "next/head";
import Link from "next/link";
import { Layout } from "src/components/Layout";

const About: NextPage = () => {
  return (
    <Layout>
      <Head>
        <title>About</title>
        <meta name="description" content="This is the about page." />
      </Head>

      <div className="mt-12 text-center">
        <h1 className="text-lg">About</h1>
        <p className="mt-4">This is the about page.</p>

        <div className="mt-5">
          <Link href="/">
            <a className="text-blue-400 underline">Top</a>
          </Link>
        </div>
      </div>
    </Layout>
  );
};

export default About;

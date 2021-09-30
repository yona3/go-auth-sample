import type { NextPage } from "next";
import Link from "next/link";
import { Layout } from "src/components/Layout";

const About: NextPage = () => {
  return (
    <Layout>
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

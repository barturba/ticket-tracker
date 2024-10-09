import { Metadata } from "next";
import Link from "next/link";

export const metadata: Metadata = {
  title: "Ticket Tracker",
};

export default function Page() {
  return (
    <div className="bg-white">
      <div className="px-6 py-24 mx-auto max-w-7xl sm:py-32 lg:px-8">
        <h2 className="text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
          Ticket Tracker
        </h2>
        <div className="flex items-center mt-10 gap-x-6">
          <Link
            href="/login"
            className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Log in
          </Link>
        </div>
      </div>
    </div>
  );
}

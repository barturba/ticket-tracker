import { Metadata } from "next";
import { Heading } from "./components/heading";
import { Button } from "./components/button";

export const metadata: Metadata = {
  title: "Ticket Tracker",
};

export default function Page() {
  return (
    <>
      <Heading>Welcome to Ticket Tracker</Heading>
      <div className="mt-8 flex items-end justify-between">
        <Button className="-my-0.5" href="/dashboard">
          Dashboard
        </Button>
      </div>
    </>
  );
}

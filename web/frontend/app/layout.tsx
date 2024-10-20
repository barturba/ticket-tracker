export const dynamic = "force-dynamic";
import "@/styles/tailwind.css";
import type { Metadata } from "next";
import type React from "react";
import { ApplicationLayout } from "@/app/application-layout";
import NextTopLoader from "nextjs-toploader";
import { SessionProvider } from "next-auth/react";

export const metadata: Metadata = {
  title: {
    template: "%s - Ticket Tracker",
    default: "Ticket Tracker",
  },
  description: "Ticket Tracker for IT incident tracking.",
  metadataBase: new URL("https://bartas.co/ticket-tracker"),
};

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html
      lang="en"
      className="text-zinc-950 antialiased lg:bg-zinc-100 dark:bg-zinc-900 dark:text-white dark:lg:bg-zinc-950"
    >
      <head>
        <link rel="preconnect" href="https://rsms.me/" />
        <link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
      </head>
      <body>
        <NextTopLoader showSpinner={false} />

        <SessionProvider>
          <ApplicationLayout>{children}</ApplicationLayout>
        </SessionProvider>
      </body>
    </html>
  );
}

import "@/app/ui/global.css";

import { inter } from "./ui/fonts";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: {
    template: "%s | Ticket Tracker",
    default: "Ticket Tracker",
  },
  description: "Ticket Tracker for IT incident tracking.",
  metadataBase: new URL("https://bartas.co/ticket-tracker"),
};
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="h-full bg-white">
      <body className={`${inter.className} antialiased h-full`}>
        {children}
      </body>
    </html>
  );
}

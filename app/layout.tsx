import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Go Koans",
  description: "120 Go exercises in 30 topics. Edit the code and run the tests.",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}

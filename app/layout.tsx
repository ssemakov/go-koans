import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Go Koans — Learn by questioning",
  description: "Small Go exercises. A little curiosity. One insight at a time.",
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

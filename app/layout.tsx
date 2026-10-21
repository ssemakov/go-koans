import type { Metadata } from "next";
import curriculum from "@/lib/curriculum.json";
import { themeInitScript } from "@/lib/theme";
import "./globals.css";

export const metadata: Metadata = {
  title: "Go Koans",
  description: `${curriculum.length} Go exercises in ${new Set(curriculum.map((exercise) => exercise.topic)).size} topics. Edit the code and run the tests.`,
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        {/* biome-ignore lint/security/noDangerouslySetInnerHtml: This trusted, static script applies the saved theme before first paint. */}
        <script dangerouslySetInnerHTML={{ __html: themeInitScript }} />
      </head>
      <body>{children}</body>
    </html>
  );
}

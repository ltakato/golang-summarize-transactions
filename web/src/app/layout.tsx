import type { Metadata } from "next";
import "./globals.css";
import React from "react";

export const metadata: Metadata = {
  title: "Expenses Summary",
  description: "Application to check expenses summarized by categories",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <div className="flex-1 space-y-4 p-8 gap-1">{children}</div>
      </body>
    </html>
  );
}

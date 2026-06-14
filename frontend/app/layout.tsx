import './globals.css';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Harasyn - Zero Trust Access Platform',
  description: 'Production-grade Zero Trust Access Platform',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className="min-h-screen flex flex-col">{children}</body>
    </html>
  );
}

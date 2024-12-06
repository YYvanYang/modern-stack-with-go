import { AuthProvider } from '@/contexts/AuthContext';
import './globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="zh" suppressHydrationWarning>
      <head />
      <body className="antialiased bg-white dark:bg-gray-900">
        <AuthProvider>{children}</AuthProvider>
      </body>
    </html>
  );
}

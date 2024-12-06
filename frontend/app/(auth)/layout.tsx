import { redirect } from 'next/navigation';
import { cookies } from 'next/headers';

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const cookieStore = cookies();
  const token = cookieStore.get('token');

  if (token) {
    redirect('/');
  }

  return <>{children}</>;
} 
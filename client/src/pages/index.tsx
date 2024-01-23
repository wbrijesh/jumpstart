import MainLayout from "@/layouts/mainLayout";
import { useAuth } from "@clerk/nextjs";

export default function Home() {
  const { isLoaded, userId, sessionId, getToken } = useAuth();

  if (!isLoaded || !userId) {
    return null;
  }
  return (
    <MainLayout>
      <p>
        Hello, {userId} your current active session is {sessionId}
      </p>
    </MainLayout>
  );
}

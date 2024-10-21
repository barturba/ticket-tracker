"use client";
import { useSession } from "next-auth/react";

export default function SessionMetrics() {
  const { data: session } = useSession();
  return (
    <div>
      {session?.user?.role === "admin" ? (
        <p>You are an admin, welcome!</p>
      ) : (
        <p>You&apos;re not an admin</p>
      )}
      <div>Session info: {JSON.stringify(session, null, 2)}</div>
    </div>
  );
}

import { Suspense } from "react";
import LatestIncidents from "@/app/ui/dashboard/latest-incidents";
import { LatestIncidentsSkeleton } from "@/app/ui/skeletons/incidents";
import { Heading, Subheading } from "@/app/components/heading";

export default async function Page() {
  return (
    <main>
      <Heading>Dashboard</Heading>
      <div className="mt-8 flex items-end justify-between">
        <Subheading>Overview</Subheading>
      </div>
      <ul>
        <Suspense fallback={<LatestIncidentsSkeleton />}>
          <LatestIncidents />
        </Suspense>
      </ul>
      <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-4"></div>
      <div className="grid grid-cols-1 gap-6 mt-6 md:grid-cols-4 lg:grid-cols-8"></div>
    </main>
  );
}

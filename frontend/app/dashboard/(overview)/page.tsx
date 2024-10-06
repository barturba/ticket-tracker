import { lusitana } from "@/app/ui/fonts";
import { Suspense } from "react";
import { RevenueChartSkeleton } from "@/app/ui/skeletons";
import IncidentsChart from "../incidents-chart";

export default async function Page() {
  // const incidents = await fetchIncidents();
  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Dashboard
      </h1>
      <ul>
        <Suspense fallback={<RevenueChartSkeleton />}>
          <IncidentsChart />
        </Suspense>
      </ul>
      <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-4"></div>
      <div className="grid grid-cols-1 gap-6 mt-6 md:grid-cols-4 lg:grid-cols-8"></div>
    </main>
  );
}

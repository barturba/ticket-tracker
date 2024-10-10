import { fetchIncidentsPages } from "@/app/lib/actions";
import { lusitana } from "@/app/ui/fonts";
import { CreateIncident } from "@/app/ui/incidents/buttons";
import Pagination from "@/app/ui/utils/pagination";
import Table from "@/app/ui/incidents/table";
import Search from "@/app/ui/search";
import { IncidentsTableSkeleton } from "@/app/ui/skeletons";
import { Metadata } from "next";
import { Suspense } from "react";

export const metadata: Metadata = {
  title: "Incidents",
};

export default async function Page(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const totalPages = (await fetchIncidentsPages(query)) ?? 0;

  return (
    <div className="w-full">
      <div className="flex w-full items-center justify-between">
        <h1 className={`${lusitana.className} text-2xl`}>Incidents</h1>
      </div>

      <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
        <Search placeholder="Search incidents..." />
        <CreateIncident />
      </div>
      <Suspense key={query + currentPage} fallback={<IncidentsTableSkeleton />}>
        <Table query={query} currentPage={currentPage} />
      </Suspense>
      <div className="mt-5 flex w-full justify-center">
        <Pagination totalPages={totalPages} />
      </div>
    </div>
  );
}

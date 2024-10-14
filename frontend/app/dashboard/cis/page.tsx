import { lusitana } from "@/app/ui/fonts";
import Pagination from "@/app/ui/utils/pagination";
import Search from "@/app/ui/search";
import { Metadata } from "next";
import { Suspense } from "react";
import { getCIs } from "@/app/lib/actions/cis";
import { CIsTableSkeleton } from "@/app/ui/skeletons/cis";
import Table from "@/app/ui/cis/table";
import { CreateCI } from "@/app/ui/cis/buttons";

export const metadata: Metadata = {
  title: "Configuration Items",
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

  const ciData: CIData = await getCIs(query, currentPage);

  const totalPages = ciData.metadata.last_page;
  const cis = ciData.cis;

  return (
    <div className="w-full">
      <div className="flex w-full items-center justify-between">
        <h1 className={`${lusitana.className} text-2xl`}>
          Configuration Items
        </h1>
      </div>

      <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
        <Search placeholder="Search cis..." />
        <CreateCI />
      </div>
      <Suspense key={query + currentPage} fallback={<CIsTableSkeleton />}>
        <Table cis={cis} />
      </Suspense>
      <div className="mt-5 flex w-full justify-center">
        <Pagination totalPages={totalPages} />
      </div>
    </div>
  );
}

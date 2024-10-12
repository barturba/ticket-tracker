import { fetchCompanies } from "@/app/lib/actions";
import { CreateCompany } from "@/app/ui/companies/buttons";
import Table from "@/app/ui/companies/table";
import { lusitana } from "@/app/ui/fonts";
import Pagination from "@/app/ui/utils/pagination";
import Search from "@/app/ui/search";
import { CompaniesTableSkeleton } from "@/app/ui/skeletons";
import { Metadata } from "next";
import { Suspense } from "react";
import { CompanyData } from "@/app/lib/definitions";

export const metadata: Metadata = {
  title: "Companies",
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

  const companydata: CompanyData = await fetchCompanies(query, currentPage);
  const totalPages = companydata.metadata.last_page;
  const companies = companydata.companies;

  return (
    <div className="w-full">
      <div className="flex w-full items-center justify-between">
        <h1 className={`${lusitana.className} text-2xl`}>Companies</h1>
      </div>

      <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
        <Search placeholder="Search companies ..." />
        <CreateCompany />
      </div>
      <Suspense key={query + currentPage} fallback={<CompaniesTableSkeleton />}>
        <Table companies={companies} />
      </Suspense>
      <div className="mt-5 flex w-full justify-center">
        <Pagination totalPages={totalPages} />
      </div>
    </div>
  );
}

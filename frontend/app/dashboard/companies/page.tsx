import AppHeading from "@/app/application-components/heading";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getCompanies } from "@/app/api/companies/companies";
import { CompaniesData } from "@/app/lib/definitions/companies";
import { formatDateToLocal } from "@/app/lib/utils";
import PaginationApp from "@/app/ui/utils/pagination-app";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Companies",
};

export default async function Companies(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const companiesData: CompaniesData = await getCompanies(query, currentPage);

  return (
    <>
      <AppHeading
        name="Companies"
        createLabel="Create Company"
        createLink="/dashboard/companies/create"
      />
      <Table className="mt-8 [--gutter:theme(spacing.6)] lg:[--gutter:theme(spacing.10)]">
        <TableHead>
          <TableRow>
            <TableHeader>ID </TableHeader>
            <TableHeader>Updated date</TableHeader>
            <TableHeader>Name</TableHeader>
          </TableRow>
        </TableHead>
        <TableBody>
          {companiesData.companies.map((company) => (
            <TableRow
              key={company.id}
              href={`/dashboard/companies/${company.id}/edit`}
              title={`Company #${company.id}`}
            >
              <TableCell>{company.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(company.updated_at)}
              </TableCell>
              <TableCell>{company.name}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <PaginationApp totalPages={companiesData.metadata.last_page} />
    </>
  );
}

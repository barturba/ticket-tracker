import AppHeading from "@/app/application-components/heading";
import { Badge } from "@/app/components/badge";
import { Button } from "@/app/components/button";
import { Heading } from "@/app/components/heading";
import {
  Pagination,
  PaginationGap,
  PaginationList,
  PaginationNext,
  PaginationPage,
  PaginationPrevious,
} from "@/app/components/pagination";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getCompanies } from "@/app/lib/actions/companies";
import { CompanyData } from "@/app/lib/definitions/companies";
import { formatDateToLocal, truncate } from "@/app/lib/utils";
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

  const companyData: CompanyData = await getCompanies(query, currentPage);

  return (
    <>
      <AppHeading
        name="Company"
        createLabel="Create Company"
        createLink="/dashboard/companies/create"
      />
      <Table className="mt-8 [--gutter:theme(spacing.6)] lg:[--gutter:theme(spacing.10)]">
        <TableHead>
          <TableRow>
            <TableHeader>ID </TableHeader>
            <TableHeader>Updated date</TableHeader>
            <TableHeader>Assigned to</TableHeader>
            <TableHeader>Short description</TableHeader>
            <TableHeader>State</TableHeader>
          </TableRow>
        </TableHead>
        <TableBody>
          {/* {companyData.companies.map((company) => (
            <TableRow
              key={company.id}
              href={`/dashboard/companies/${company.id}/edit`}
              title={`Company #${company.id}`}
            >
              <TableCell>{company.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(company.updated_at)}
              </TableCell>
              <TableCell>{company.assigned_to_name}</TableCell>
              <TableCell>{truncate(company.short_description)}</TableCell>
              <TableCell>
                <Badge className="max-sm:hidden" state={company.state}>
                  {company.state}
                </Badge>
              </TableCell>
            </TableRow>
          ))} */}
        </TableBody>
      </Table>

      <Pagination>
        <PaginationPrevious href="?page=2" />
        <PaginationList>
          <PaginationPage href="?page=1">1</PaginationPage>
          <PaginationPage href="?page=2">2</PaginationPage>
          <PaginationPage href="?page=3" current>
            3
          </PaginationPage>
          <PaginationPage href="?page=4">4</PaginationPage>
          <PaginationGap />
          <PaginationPage href="?page=65">65</PaginationPage>
          <PaginationPage href="?page=66">66</PaginationPage>
        </PaginationList>
        <PaginationNext href="?page=4" />
      </Pagination>
    </>
  );
}

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
import { getCIs } from "@/app/lib/actions/cis";
import { CIsData } from "@/app/lib/definitions/cis";
import { formatDateToLocal, truncate } from "@/app/lib/utils";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "CIs",
};

export default async function CIs(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const cisData: CIsData = await getCIs(query, currentPage);

  return (
    <>
      <AppHeading
        name="CI"
        createLabel="Create CI"
        createLink="/dashboard/cis/create"
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
          {cisData.cis.map((ci) => (
            <TableRow
              key={ci.id}
              href={`/dashboard/cis/${ci.id}/edit`}
              title={`CI #${ci.id}`}
            >
              <TableCell>{ci.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(ci.updated_at)}
              </TableCell>
              <TableCell>{ci.name}</TableCell>
            </TableRow>
          ))}
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

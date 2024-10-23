import AppHeading from "@/app/application-components/heading";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getCIs } from "@/app/api/cis/cis";
import { formatDateToLocal } from "@/app/lib/utils";
import PaginationApp from "@/app/ui/utils/pagination-app";
import type { Metadata } from "next";
import { CIsData } from "@/app/api/cis/cis.d";

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
        name="CIs"
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
              <TableCell>{ci.id.slice(0, 8)}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(ci.updated_at)}
              </TableCell>
              <TableCell>{ci.name}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <PaginationApp totalPages={cisData.metadata.last_page} />
    </>
  );
}

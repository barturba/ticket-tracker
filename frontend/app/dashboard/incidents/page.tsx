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
import { getIncidents } from "@/app/lib/actions/incidents";
import { IncidentData } from "@/app/lib/definitions/incidents";
import { formatDateToLocal, truncate } from "@/app/lib/utils";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Incidents",
};

export default async function Incidents(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const incidentData: IncidentData = await getIncidents(query, currentPage);

  return (
    <>
      <AppHeading name="Incidents" createLink="/dashboard/incidents/create" />
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
          {incidentData.incidents.map((incident) => (
            <TableRow
              key={incident.id}
              href={`/dashboard/incidents/${incident.id}/edit`}
              title={`Incident #${incident.id}`}
            >
              <TableCell>{incident.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(incident.updated_at)}
              </TableCell>
              <TableCell>{incident.assigned_to}</TableCell>
              <TableCell>{truncate(incident.short_description)}</TableCell>
              <TableCell>
                <Badge className="max-sm:hidden" state={incident.state}>
                  {incident.state}
                </Badge>
              </TableCell>
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

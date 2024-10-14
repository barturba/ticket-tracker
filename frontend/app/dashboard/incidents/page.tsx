import AppHeading from "@/app/application-components/heading";
import { Badge } from "@/app/components/badge";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getIncidents } from "@/app/lib/actions/incidents";
import { IncidentsData } from "@/app/lib/definitions/incidents";
import { formatDateToLocal, truncate } from "@/app/lib/utils";
import PaginationApp from "@/app/ui/utils/pagination-app";
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

  const incidentsData: IncidentsData = await getIncidents(query, currentPage);
  console.log(`total records: ${incidentsData.metadata.last_page}`);

  return (
    <>
      <AppHeading
        name="Incident"
        createLabel="Create Incident"
        createLink="/dashboard/incidents/create"
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
          {incidentsData.incidents.map((incident) => (
            <TableRow
              key={incident.id}
              href={`/dashboard/incidents/${incident.id}/edit`}
              title={`Incident #${incident.id}`}
            >
              <TableCell>{incident.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(incident.updated_at)}
              </TableCell>
              <TableCell>{incident.assigned_to_name}</TableCell>
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

      <PaginationApp totalPages={incidentsData.metadata.last_page} />
    </>
  );
}

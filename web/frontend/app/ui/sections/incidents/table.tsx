"use client";
import { IncidentsData } from "@/app/api/incidents/incidents.d";
import { Badge } from "@/app/components/badge";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { formatDateToLocal, truncate } from "@/app/lib/utils";
import { Suspense } from "react";

export default function IncidentsTable({
  incidentsData,
}: {
  incidentsData: IncidentsData;
}) {
  return (
    <>
      <Table className="mt-8 [--gutter:theme(spacing.6)] lg:[--gutter:theme(spacing.10)]">
        <TableHead>
          <TableRow>
            <TableHeader className="sm:pl-0">ID </TableHeader>
            <TableHeader className="hidden lg:table-cell">State</TableHeader>
            <TableHeader className="hidden lg:table-cell">
              Short description
            </TableHeader>
            <TableHeader className="hidden sm:table-cell">
              Assigned to
            </TableHeader>
            <TableHeader className="hidden sm:table-cell">
              Updated date
            </TableHeader>
          </TableRow>
        </TableHead>
        <TableBody>
          <Suspense fallback={<p>Loading...</p>}>
            {incidentsData.incidents.map((incident) => (
              <TableRow
                key={incident.id}
                href={`/dashboard/incidents/${incident.id}/edit`}
                title={`Incident #${incident.id}`}
              >
                <TableCell className="sm:w-auto sm:max-w-none sm:pl-0">
                  {incident.id}
                  <dl className="font-normal lg:hidden">
                    <dt className="sr-only">State</dt>
                    <dd className="mt-1 truncate text-gray-700">
                      <Badge state={incident.state}>{incident.state}</Badge>
                    </dd>
                    <dt className="sr-only">Short description</dt>
                    <dd className="mt-1 truncate text-gray-700">
                      {truncate(incident.short_description, true)}
                    </dd>
                    <dt className="sr-only">Assigned to</dt>
                    <dd className="mt-1 text-gray-500">
                      {incident.assigned_to_name}
                    </dd>
                    <dt className="sr-only">Updated at</dt>
                    <dd className="mt-1 text-zinc-500">
                      {incident.updated_at}
                    </dd>
                  </dl>
                </TableCell>
                <TableCell className="hidden lg:table-cell">
                  <Badge state={incident.state}>{incident.state}</Badge>
                </TableCell>
                <TableCell className="hidden lg:table-cell">
                  {truncate(incident.short_description, true)}
                </TableCell>
                <TableCell className="hidden sm:table-cell">
                  {incident.assigned_to_name}
                </TableCell>
                <TableCell className="text-zinc-500 hidden sm:table-cell">
                  {formatDateToLocal(incident.updated_at)}
                </TableCell>
              </TableRow>
            ))}
          </Suspense>
        </TableBody>
      </Table>
    </>
  );
}

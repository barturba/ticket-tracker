"use client";
import { Badge } from "@/app/components/badge";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { IncidentsData } from "@/app/lib/definitions/incidents";
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
            <TableHeader>ID </TableHeader>
            <TableHeader>Updated date</TableHeader>
            <TableHeader>Assigned to</TableHeader>
            <TableHeader>Short description</TableHeader>
            <TableHeader>State</TableHeader>
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
          </Suspense>
        </TableBody>
      </Table>
    </>
  );
}

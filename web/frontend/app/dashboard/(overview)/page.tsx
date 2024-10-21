import AppHeading from "@/app/application-components/heading";
import { getIncidents } from "@/app/api/incidents/incidents";
import PaginationApp from "@/app/ui/utils/pagination-app";
import type { Metadata } from "next";
import { IncidentsData } from "@/app/api/incidents/incidents.d";
import IncidentsTable from "../../ui/sections/incidents/table";
// import SessionMetrics from "@/app/application-components/session-metrics";

export const metadata: Metadata = {
  title: "Dashboard",
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

  return (
    <>
      {/* <SessionMetrics /> */}
      <AppHeading
        name="Dashboard - Recent Incidents"
        createLabel="Create Incident"
        createLink="/dashboard/incidents/create"
      />
      <IncidentsTable incidentsData={incidentsData} />
      <PaginationApp totalPages={incidentsData.metadata.last_page} />
    </>
  );
}

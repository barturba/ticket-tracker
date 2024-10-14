import AppHeading from "@/app/application-components/heading";
import { getIncidents } from "@/app/api/incidents/incidents";
import PaginationApp from "@/app/ui/utils/pagination-app";
import type { Metadata } from "next";
import IncidentsTable from "./table";
import { IncidentsData } from "@/app/api/incidents/incidents.d";

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

  return (
    <>
      <AppHeading
        name="Incidents"
        createLabel="Create Incident"
        createLink="/dashboard/incidents/create"
      />
      <IncidentsTable incidentsData={incidentsData} />
      <PaginationApp totalPages={incidentsData.metadata.last_page} />
    </>
  );
}

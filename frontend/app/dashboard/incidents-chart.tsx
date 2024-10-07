import { fetchIncidents } from "@/app/lib/actions";
import { Incident } from "@/app/lib/definitions";

export default async function IncidentsChart() {
  const incidents = await fetchIncidents();

  if (!incidents || incidents.length === 0) {
    return <p className="mt-4 text-gray-400">No data available.</p>;
  }
  return (
    <ul>
      {incidents.map((i: Incident) => (
        <li key={i.id}>{i.short_description}</li>
      ))}
    </ul>
  );
}

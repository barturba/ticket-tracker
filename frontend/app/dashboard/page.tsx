import { fetchIncidents } from "../lib/actions";
import { lusitana } from "../ui/fonts";

export default async function Page() {
  const incidents = await fetchIncidents();
  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Dashboard
      </h1>
      <ul>
        {incidents?.map((incident) => {
          return <li key={incident.id}>{incident.short_description}</li>;
        })}
      </ul>
      <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-4"></div>
      <div className="grid grid-cols-1 gap-6 mt-6 md:grid-cols-4 lg:grid-cols-8"></div>
    </main>
  );
}
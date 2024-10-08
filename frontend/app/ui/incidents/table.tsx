import { fetchFilteredIncidents } from "@/app/lib/actions";
import IncidentStatus from "@/app/ui/incidents/status";
import { DeleteIncident, UpdateIncident } from "@/app/ui/incidents/buttons";
import { Incident } from "@/app/lib/definitions";
export default async function IncidentTable({
  query,
  currentPage,
}: {
  query: string;
  currentPage: number;
}) {
  const incidents: Incident[] = await fetchFilteredIncidents(
    query,
    currentPage
  );
  if (!incidents) {
    return <p>No data</p>;
  }
  return (
    <div className="mt-6 flow-root">
      <div className="inline-block min-w-full align-middle">
        <div className="rounded-lg bg-gray-50 p-2 md:pt-0">
          <div className="md:hidden">
            {incidents?.map((incident) => (
              <div
                key={incident.id}
                className="mb-2 w-full rounded-md bg-white p-4"
              >
                <div className="flex items-center justify-between border-b pb-4">
                  <div>
                    <div className="mb-2 flex items-center">
                      <p>{incident.id}</p>
                    </div>
                  </div>
                  <IncidentStatus status={incident.state} />
                </div>
                <div className="flex w-full items-center justify-between pt-4">
                  <p className="text-xl font-medium">{incident.created_at}</p>
                </div>
                <div className="flex justify-end gap-2">
                  <UpdateIncident id={incident.id} />
                  <DeleteIncident id={incident.id} />
                </div>
              </div>
            ))}
          </div>
          <table className="hidden min-w-full text-gray-900 md:table">
            <thead className="rounded-lg text-left text-sm font-normal">
              <tr>
                <th scope="col" className="px-4 py-5 font-medium sm:pl-6">
                  ID
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Short Description
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Assigned To
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Date
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  State
                </th>
                <th scope="col" className="relative py-3 pl-6 pr-3">
                  <span className="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {incidents?.map((incident) => (
                <tr
                  key={incident.id}
                  className="w-full border-b py-3 text-sm last-of-type:border-none [&:first-child>td:first-child]:rounded-tl-lg [&:first-child>td:last-child]:rounded-tr-lg [&:last-child>td:first-child]:rounded-bl-lg [&:last-child>td:last-child]:rounded-br-lg"
                >
                  <td className="whitespace-nowrap py-3 pl-6 pr-3">
                    <div className="flex items-center gap-3">
                      <p>{incident.id}</p>
                    </div>
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {incident.short_description}
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {incident.assigned_to_id}
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {incident.created_at}
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    <IncidentStatus status={incident.state} />
                  </td>
                  <td className="whitespace-nowrap py-3 pl-6 pr-3">
                    <div className="flex justify-end gap-3">
                      <UpdateIncident id={incident.id} />
                      <DeleteIncident id={incident.id} />
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

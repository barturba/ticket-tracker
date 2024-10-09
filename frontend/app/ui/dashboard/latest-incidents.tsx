import { ArrowPathIcon } from "@heroicons/react/24/outline";
import clsx from "clsx";
// import Image from "next/image";
import { lusitana } from "@/app/ui/fonts";
import { fetchLatestIncidents } from "@/app/lib/actions";
import { Incident } from "@/app/lib/definitions";
import Link from "next/link";

import dayjs from "dayjs";
import { truncate } from "@/app/lib/utils";

export default async function LatestIncidents() {
  console.log("LatestIncidents");
  const latestIncidents: Incident[] = await fetchLatestIncidents();

  return (
    <div className="flex w-full flex-col md:col-span-4">
      <h2 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Latest Incidents
      </h2>
      <div className="flex grow flex-col justify-between rounded-xl bg-gray-50 p-4">
        <div className="bg-white px-6">
          {latestIncidents.map((incident, i) => {
            return (
              <div
                key={incident.id}
                className={clsx(
                  "flex flex-row items-center justify-between py-4",
                  {
                    "border-t": i !== 0,
                  }
                )}
              >
                <div className="flex items-center">
                  {/* <Image
                    src={incident.image_url}
                    alt={`${incident.name}'s profile picture`}
                    className="mr-4 rounded-full"
                    width={32}
                    height={32}
                  /> */}
                  <div className="min-w-0">
                    {/* <p className="truncate text-sm font-semibold md:text-base">
                      {incident.short_description}
                    </p> */}
                    <Link
                      href={`/dashboard/incidents/${incident.id}/edit`}
                      className="truncate text-sm font-semibold md:text-base"
                    >
                      {truncate(incident.short_description, 35, true)}
                    </Link>
                    {/* <p className="hidden text-sm text-gray-500 sm:block">
                      {incident.id}
                    </p> */}
                    {/* <Link
                      href={`/dashboard/incidents/${incident.id}/edit`}
                      className="hidden text-sm text-gray-500 sm:block"
                    >
                      {incident.id}
                    </Link> */}
                  </div>
                </div>
                <p
                  className={`${lusitana.className} truncate text-sm font-medium md:text-base`}
                >
                  {dayjs(incident.updated_at).format("MMMM DD, YYYY")}
                </p>
              </div>
            );
          })}
        </div>
        <div className="flex items-center pb-2 pt-6">
          <ArrowPathIcon className="h-5 w-5 text-gray-500" />
          <h3 className="ml-2 text-sm text-gray-500 ">Updated just now</h3>
        </div>
      </div>
    </div>
  );
}

import { ClockIcon, CheckIcon } from "@heroicons/react/24/outline";
import { clsx } from "clsx";

export default function IncidentStatus({ status }: { status: string }) {
  return (
    <span
      className={clsx(
        "inline-flex items-center rounded-full px-2 py-1 text-xs",
        {
          "bg-green-500 text-white": status === "New",
          "bg-gray-100 text-gray-500":
            status == "Assigned" ||
            status === "In Progress" ||
            status === "On Hold" ||
            status === "Resolved",
        }
      )}
    >
      {status === "New" ? (
        <>
          New
          <CheckIcon className="ml-1 w-4 text-white" />
        </>
      ) : null}
      {status === "Assigned" ? (
        <>
          Assigned
          <ClockIcon className="ml-1 w-4 text-gray-500" />
        </>
      ) : null}
      {status === "In Progress" ? (
        <>
          In Progress
          <ClockIcon className="ml-1 w-4 text-gray-500" />
        </>
      ) : null}
      {status === "On Hold" ? (
        <>
          On Hold
          <ClockIcon className="ml-1 w-4 text-gray-500" />
        </>
      ) : null}
      {status === "Resolved" ? (
        <>
          Resolved
          <ClockIcon className="ml-1 w-4 text-gray-500" />
        </>
      ) : null}
    </span>
  );
}

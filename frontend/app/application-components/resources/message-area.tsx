import { CompanyState } from "@/app/lib/actions/companies";
import { IncidentState } from "@/app/lib/actions/incidents";

export default function MessageArea({
  state,
}: {
  state: IncidentState | CompanyState;
}) {
  return (
    <div aria-live="polite" aria-atomic="true">
      {state.errors && state.message ? (
        <p className="mt-2 text-sm text-red-500">{state.message}</p>
      ) : null}
      {!state.errors && state.message ? (
        <p className="mt-2 text-sm text-green-500">{state.message}</p>
      ) : null}
    </div>
  );
}

import { IncidentState } from "@/app/lib/actions/incidents";

export default function MessageArea({ state }: { state: IncidentState }) {
  return (
    <div aria-live="polite" aria-atomic="true">
      {state.message ? (
        <p className="mt-2 text-sm text-red-500">{state.message}</p>
      ) : null}
    </div>
  );
}

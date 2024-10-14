import { IncidentState } from "@/app/lib/actions/incidents";

export default function MessageArea({ state }: { state: IncidentState }) {
  console.log(`MessageArea state: ${JSON.stringify(state, null, 2)}`);
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

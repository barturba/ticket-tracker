"use client";
import { useFormStatus } from "react-dom";

export default function SubmitButton() {
  const { pending } = useFormStatus();

  return (
    <>
      <button
        type="submit"
        className="rounded-md bg-zinc-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-zinc-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-zinc-500 disabled:opacity-50"
        disabled={pending}
        aria-disabled={pending}
      >
        Save
      </button>
    </>
  );
}

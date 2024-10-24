"use client";

export default function SubmitButton({ isPending }: { isPending: boolean }) {
  return (
    <>
      <button
        type="submit"
        className="rounded-md mt-6 bg-zinc-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-zinc-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-zinc-500 disabled:opacity-50"
        disabled={isPending}
      >
        {isPending ? "Saving..." : "Save Changes"}
      </button>
    </>
  );
}

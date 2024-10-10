"use client";
import { createCompany, State } from "@/app/lib/actions";
import { BuildingOfficeIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import { Button } from "../button";
import { useFormState } from "react-dom";

export default function Form() {
  const initialState: State = { message: null, errors: {} };
  const [state, formAction] = useFormState(createCompany, initialState);

  return (
    <form action={formAction}>
      <div className="rounded-md bg-gray-50 p-4 md:p-6">
        {/* Company Name */}
        <div className="mb-4">
          <label
            htmlFor="company-name"
            className="mb-2 block text-sm font-medium"
          >
            Enter a company name
          </label>
          <div className="relative mt-2 rounded-md">
            <div className="relative">
              <input
                id="company-name"
                name="name"
                type="text"
                placeholder="Enter a company name"
                className="peer block w-full rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
                aria-describedby="name-error"
              />
            </div>
            <BuildingOfficeIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500" />
          </div>
        </div>
      </div>
      <div className="mt-6 flex justify-end gap-4">
        <Link
          href="/dashboard/companies"
          className="flex h-10 items-center rounded-lg bg-gray-100 px-4 text-sm font-medium text-gray-600 transition-colors hover:bg-gray-200"
        >
          Cancel
        </Link>
        <Button type="submit">Create Company</Button>
      </div>
    </form>
  );
}

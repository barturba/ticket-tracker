"use client";
import {
  fetchUsers,
  fetchUsersByCompany,
  State,
  updateIncident,
} from "@/app/lib/actions";
import {
  CompaniesField,
  ConfigurationItemsField,
  UsersField,
} from "@/app/lib/definitions";
import {
  BuildingOffice2Icon,
  CheckIcon,
  ClockIcon,
  CpuChipIcon,
  UserCircleIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { Button } from "../button";
import { IncidentForm } from "@/app/lib/definitions";
import { useFormState, useFormStatus } from "react-dom";
import { useActionState, useEffect, useState } from "react";

function SubmitButton() {
  const { pending } = useFormStatus();

  return (
    <Button type="submit" aria-disabled={pending}>
      Update incident
    </Button>
  );
}

export default function EditForm({
  incident,
  initialUsers,
  companies,
  configurationItems,
}: {
  incident: IncidentForm;
  companies: CompaniesField[];
  initialUsers: UsersField[];
  configurationItems: ConfigurationItemsField[];
}) {
  const initialState: State = { message: null, errors: {} };
  const updateIncidentWithId = updateIncident.bind(null, incident.id);
  const [state, formAction] = useActionState(
    updateIncidentWithId,
    initialState
  );
  const [users, setUsers] = useState(initialUsers);
  const [selectedCompany, setSelectedCompany] = useState(incident.company_id);

  const handleChange = (event) => {
    setSelectedCompany(event.target.value);
  };

  useEffect(() => {
    console.log(`selectedCompany: ${selectedCompany}`);
    const fetchData = async () => {
      const data = await fetchUsersByCompany(selectedCompany);
      // console.log(`data.length: ${data.length}`);
      setUsers(data);
    };
    fetchData()
      // make sure to catch any error
      .catch(console.error);
  }, [selectedCompany]);

  // TODO [ ]: Get users and configuration items for selected company

  // const handleChange = async (e): Promise<void> => {
  //   const { name, value } = e.target;
  //   const newUsers = await fetchUsersByCompany(value);
  //   setUsers(newUsers);
  //   console.log(`name: ${name} value: ${value}`);
  //   console.log(`newUsers: ${newUsers}`);
  // };
  // const [users, setUsers] = useState([]);

  // useEffect(() => {
  //   const fetchData = async () => {
  //     const newUsers = await fetchUsersByCompany(value);
  //     setUsers(newUsers);
  //   };
  //   fetchData();
  // }, []);
  console.log(`got incident: ${JSON.stringify(incident)}`);

  return (
    <form action={formAction}>
      <div className="rounded-md bg-gray-50 p-4 md:p-6">
        {/* Company Name */}
        <div className="mb-4">
          <label htmlFor="company" className="mb-2 block text-sm font-medium">
            Choose company
          </label>
          <div className="relative">
            <select
              id="companyId"
              name="company_id"
              className="peer block w-full cursor-pointer rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
              // defaultValue={incident.company_id}
              // defaultValue={selectedCompany}
              aria-describedby="company-error"
              onChange={handleChange}
            >
              <option value="" disabled>
                Select a company
              </option>
              {companies.map((company) => (
                <option key={company.id} value={company.id}>
                  {company.name}
                </option>
              ))}
            </select>
            <BuildingOffice2Icon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500" />
          </div>

          <div id="company-error" aria-live="polite" aria-atomic="true">
            {state.errors?.companyId &&
              state.errors.companyId.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </div>

        {/* Assigned To Name */}
        <div className="mb-4">
          <label
            htmlFor="assignedToId"
            className="mb-2 block text-sm font-medium"
          >
            Choose who this incident is assigned to
          </label>
          <div className="relative">
            <select
              id="assignedToId"
              name="assigned_to_id"
              className="peer block w-full cursor-pointer rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
              defaultValue={incident.assigned_to_id}
              aria-describedby="assigned-to-error"
            >
              <option value="" disabled>
                Select a user
              </option>
              {users.map((user) => (
                <option key={user.id} value={user.id}>
                  {user.name}
                </option>
              ))}
            </select>
            <UserCircleIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500" />
          </div>

          <div id="assigned-to-error" aria-live="polite" aria-atomic="true">
            {state.errors?.assignedToId &&
              state.errors.assignedToId.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </div>

        {/* Configuration Item Name */}
        <div className="mb-4">
          <label
            htmlFor="configurationItemId"
            className="mb-2 block text-sm font-medium"
          >
            Choose which configuration item this incident is for
          </label>
          <div className="relative">
            <select
              id="configurationItemId"
              name="configuration_item_id"
              className="peer block w-full cursor-pointer rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
              defaultValue={incident.configuration_item_id}
              aria-describedby="configuration-item-error"
            >
              <option value="" disabled>
                Select a configuration item
              </option>
              {configurationItems.map((ci) => (
                <option key={ci.id} value={ci.id}>
                  {ci.name}
                </option>
              ))}
            </select>
            <CpuChipIcon className="pointer-events-none absolute left-3 top-1/2 h-[18px] w-[18px] -translate-y-1/2 text-gray-500" />
          </div>

          <div id="company-error" aria-live="polite" aria-atomic="true">
            {state.errors?.configurationItemId &&
              state.errors.configurationItemId.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </div>

        {/* Short Description */}
        <div className="mb-4">
          <label
            htmlFor="short-description"
            className="mb-2 block text-sm font-medium"
          >
            Enter a short description
          </label>
          <div className="relative mt-2 rounded-md">
            <div className="relative">
              <input
                id="shortDescription"
                name="short_description"
                type="text"
                placeholder="Enter short description"
                className="peer block w-full rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
                aria-describedby="amount-error"
                defaultValue={incident.short_description}
              />
            </div>
          </div>

          <div
            id="short-description-error"
            aria-live="polite"
            aria-atomic="true"
          >
            {state.errors?.shortDescription &&
              state.errors.shortDescription.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </div>

        {/* Description */}
        <div className="mb-4">
          <label
            htmlFor="description"
            className="mb-2 block text-sm font-medium"
          >
            Enter a description
          </label>
          <div className="relative mt-2 rounded-md">
            <div className="relative">
              <textarea
                rows={3}
                id="description"
                name="description"
                placeholder="Enter a description"
                className="peer block w-full rounded-md border border-gray-200 py-2 pl-10 text-sm outline-2 placeholder:text-gray-500"
                aria-describedby="amount-error"
                defaultValue={incident.description}
              />
            </div>
          </div>

          <div
            id="short-description-error"
            aria-live="polite"
            aria-atomic="true"
          >
            {state.errors?.shortDescription &&
              state.errors.shortDescription.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </div>

        {/* Incident State */}
        <fieldset>
          <legend className="mb-2 block text-sm font-medium">
            Set the incident state
          </legend>
          <div className="rounded-md border border-gray-200 bg-white px-[14px] py-3">
            <div className="flex gap-4">
              <div className="flex items-center">
                <input
                  id="new"
                  name="state"
                  type="radio"
                  value="New"
                  defaultChecked={incident.state === "New"}
                  className="text-white-600 h-4 w-4 cursor-pointer border-gray-300 bg-gray-100 focus:ring-2"
                />
                <label
                  htmlFor="new"
                  className="ml-2 flex cursor-pointer items-center gap-1.5 rounded-full bg-gray-100 px-3 py-1.5 text-xs font-medium text-gray-600"
                >
                  New <ClockIcon className="h-4 w-4" />
                </label>
              </div>
              <div className="flex items-center">
                <input
                  id="In Progress"
                  name="state"
                  type="radio"
                  value="In Progress"
                  defaultChecked={incident.state === "In Progress"}
                  className="h-4 w-4 cursor-pointer border-gray-300 bg-gray-100 text-gray-600 focus:ring-2"
                />
                <label
                  htmlFor="inProgress"
                  className="ml-2 flex cursor-pointer items-center gap-1.5 rounded-full bg-green-500 px-3 py-1.5 text-xs font-medium text-white"
                >
                  In Progress
                  <CheckIcon className="h-4 w-4" />
                </label>
              </div>
            </div>
          </div>
          <div id="state-error" aria-live="polite" aria-atomic="true">
            {state.errors?.state &&
              state.errors.state.map((error: string) => (
                <p className="mt-2 text-sm text-red-500" key={error}>
                  {error}
                </p>
              ))}
          </div>
        </fieldset>
        <div aria-live="polite" aria-atomic="true">
          <div>
            {state.message ? (
              <p className="mt-2 text-sm text-red-500">{state.message}</p>
            ) : null}
          </div>
        </div>
      </div>
      <div className="mt-6 flex justify-end gap-4">
        <Link
          href="/dashboard/incidents"
          className="flex h-10 items-center rounded-lg bg-gray-100 px-4 text-sm font-medium text-gray-600 transition-colors hover:bg-gray-200"
        >
          Cancel
        </Link>
        <SubmitButton />
      </div>
    </form>
  );
}

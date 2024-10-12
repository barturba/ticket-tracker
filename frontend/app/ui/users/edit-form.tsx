"use client";
import { State } from "@/app/lib/actions";
import Link from "next/link";
import { Button } from "../button";
import { useFormStatus } from "react-dom";
import { useActionState, useState } from "react";
import { updateIncident } from "@/app/lib/actions/incidents";
import { CompaniesField } from "@/app/lib/definitions/companies";
import { IncidentForm } from "@/app/lib/definitions/incidents";
import { UsersField } from "@/app/lib/definitions/users";
import { CIsField } from "@/app/lib/definitions/cis";

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
  cis,
}: {
  incident: IncidentForm;
  companies: CompaniesField[];
  initialUsers: UsersField[];
  cis: CIsField;
}) {
  const initialState: State = { message: null, errors: {} };
  const updateIncidentWithId = updateIncident.bind(null, incident.id);
  const [state, formAction] = useActionState(
    updateIncidentWithId,
    initialState
  );
  const [users] = useState(initialUsers);
  const [setSelectedCompany] = useState(incident.company_id);

  // const handleChange = (event) => {
  //   setSelectedCompany(event.target.value);
  // };

  // useEffect(() => {
  //   console.log(`selectedCompany: ${selectedCompany}`);
  //   const fetchData = async () => {
  //     const data = await fetchUsersByCompany(selectedCompany);
  //     // console.log(`data.length: ${data.length}`);
  //     setUsers(data);
  //   };
  //   fetchData()
  //     // make sure to catch any error
  //     .catch(console.error);
  // }, [selectedCompany]);

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
  // console.log(`got incident: ${JSON.stringify(incident)}`);

  return (
    <form action={formAction}>
      <div className="rounded-md bg-gray-50 p-4 md:p-6">
        {/* Company Name */}
        {/* <div className="mb-4">
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
        </div> */}

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

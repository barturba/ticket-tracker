"use client";
import { Button } from "@/app/components/button";
import { useActionState, useState } from "react";
import { updateIncident } from "@/app/lib/actions/incidents";
import { IncidentForm } from "@/app/lib/definitions/incidents";
import { CI, CIField } from "@/app/lib/definitions/cis";
import { User, UserField } from "@/app/lib/definitions/users";
import { Company, CompanyField } from "@/app/lib/definitions/companies";
import { Subheading } from "@/app/components/heading";
import { Divider } from "@/app/components/divider";
import { Input } from "@/app/components/input";
import { Textarea } from "@/app/components/textarea";
import {
  Description,
  ErrorMessage,
  Field,
  FieldGroup,
  Fieldset,
  Label,
  Legend,
} from "@/app/components/fieldset";
import { Select } from "@/app/components/select";
import { Listbox, ListboxLabel, ListboxOption } from "@/app/components/listbox";
import { useFormStatus } from "react-dom";
import CompanyInput from "@/app/application-components/incident/form-input";
import FormInput from "@/app/application-components/incident/form-input";

type State = {
  message: string;
  errors: {
    shortDescription?: string[];
    description?: string[];
    incidentId?: string[];
    companyID?: string[];
    assignedToId?: string[];
    configurationItemId?: string[];
    state?: string[];
  };
};

export default function EditForm({
  incident,
  initialUsers,
  companies,
  cis,
}: {
  incident: IncidentForm;
  initialUsers: UserField[];
  companies: CompanyField[];
  cis: CIField[];
}) {
  const initialState: State = { message: "", errors: {} };
  const updateIncidentWithId = updateIncident.bind(null, incident.id);
  const [state, formAction] = useActionState(
    updateIncidentWithId,
    initialState
  );
  console.log(`EditForm incident: ${JSON.stringify(incident, null, 2)}`);

  return (
    <>
      <Subheading>Summary</Subheading>
      <Divider className="mt-4" />

      <form action={formAction}>
        {/* Incident Details*/}

        <Fieldset aria-label="Incident details">
          <FieldGroup>
            {/* Company */}
            <FormInput
              label="Company"
              id="companyID"
              name="company_id"
              placeholder="Select a company"
              inputs={companies}
              defaultValue={incident.company_id}
              invalid={
                state.errors?.companyId && state.errors.companyId.length > 0
              }
              errorMessage={state.errors?.shortDescription?.join(", ")}
            />
            {/* <Field>
              <Label>Company</Label>
              <Select id="companyId" name="company_id">
                <option value="" disabled>
                  Select a company
                </option>
                {companies.map((company) => (
                  <option key={company.id} value={company.id}>
                    {company.name}
                  </option>
                ))}
              </Select>
              {state.errors.has('full_name') && <ErrorMessage>{errors.get('full_name')}</ErrorMessage>
              )}
            </Field> */}

            {/* Assigned To*/}
            <Field>
              <Label>Assigned To</Label>
              <Select
                id="assignedToId"
                name="assigned_to_id"
                aria-describedby="assigned-to-error"
              >
                <option value="" disabled>
                  Select a user
                </option>
                {initialUsers.map((user) => (
                  <option key={user.id} value={user.id}>
                    {`${user.first_name} ${user.last_name}`}
                  </option>
                ))}
              </Select>
            </Field>

            {/* CI */}
            <Field>
              <Label>CI</Label>
              <Select
                id="configurationItemId"
                name="configuration_item_id"
                aria-describedby="configuration-item-error"
              >
                <option value="" disabled>
                  Select a CI
                </option>
                {cis.map((ci) => (
                  <option key={ci.id} value={ci.id}>
                    {ci.name}
                  </option>
                ))}
              </Select>
            </Field>

            {/* Short Description */}
            <Field>
              <Label>Short description</Label>
              <Input
                name="short_description"
                defaultValue={incident.short_description}
                invalid={
                  state.errors?.shortDescription &&
                  state.errors.shortDescription.length > 0
                }
              />

              {state.errors?.shortDescription &&
                state.errors?.shortDescription.map((error: string) => (
                  <ErrorMessage key={error}>{error}</ErrorMessage>
                ))}
            </Field>

            {/* Description */}
            <Field>
              <Label>Description</Label>
              <Textarea
                name="description"
                defaultValue={incident.description.String}
              />
              <Description>
                Provide a detailed description of the incident
              </Description>
            </Field>

            {/* Incident State */}

            <Field>
              <Label>State</Label>
              <Listbox name="state" defaultValue="New">
                <ListboxOption value="New">
                  <ListboxLabel>New</ListboxLabel>
                </ListboxOption>
                <ListboxOption value="In Progress">
                  <ListboxLabel>In Progress</ListboxLabel>
                </ListboxOption>
                <ListboxOption value="Assigned">
                  <ListboxLabel>Assigned</ListboxLabel>
                </ListboxOption>
                <ListboxOption value="On Hold">
                  <ListboxLabel>On Hold</ListboxLabel>
                </ListboxOption>
                <ListboxOption value="Resolved">
                  <ListboxLabel>Resolved</ListboxLabel>
                </ListboxOption>
              </Listbox>
            </Field>
          </FieldGroup>
        </Fieldset>

        <Divider className="my-10" soft />

        <div className="flex justify-end gap-4">
          <Button type="reset" plain>
            Reset
          </Button>
          <button type="submit">Send</button>
        </div>
      </form>
    </>
  );
}

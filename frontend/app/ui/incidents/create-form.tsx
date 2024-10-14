"use client";
import Link from "next/link";
import { Button } from "@/app/components/button";
import { useFormState } from "react-dom";
import { createIncident, IncidentState } from "@/app/lib/actions/incidents";
import { UserField } from "@/app/lib/definitions/users";
import { CIField } from "@/app/lib/definitions/cis";
import { CompanyField } from "@/app/lib/definitions/companies";
import FormWrapper from "@/app/application-components/form-wrapper";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import FormInput from "@/app/application-components/incident/form-input";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
import DescriptionTextarea from "@/app/application-components/incident/description-textarea";
import StateListbox from "@/app/application-components/incident/state-listbox";
import { Divider } from "@/app/components/divider";
import { useActionState } from "react";
import MessageArea from "@/app/application-components/incident/message-area";

export default function CreateIncidentForm({
  companies,
  users,
  cis,
}: {
  companies: CompanyField[];
  users: UserField[];
  cis: CIField[];
}) {
  const initialState: IncidentState = { message: "", errors: {} };
  const [state, formAction] = useActionState(createIncident, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="Incident details">
          <FieldGroup>
            {/* Company */}
            <FormInput
              label="Company"
              id="companyID"
              name="company_id"
              placeholder="Select a company"
              inputs={companies}
              invalid={
                state.errors?.companyId && state.errors.companyId.length > 0
              }
              errorMessage={state.errors?.companyId?.join(", ")}
            />

            {/* Assigned To*/}
            <FormInput
              label="Assigned To"
              id="assignedToId"
              name="assigned_to_id"
              placeholder="Select a user"
              inputs={users.map((user) => ({
                id: user.id,
                name: `${user.first_name} ${user.last_name}`,
              }))}
              invalid={
                !!state.errors?.assignedToId &&
                state.errors.assignedToId.length > 0
              }
              errorMessage={state.errors?.assignedToId?.join(", ")}
            />

            {/* CI */}
            <FormInput
              label="CI"
              id="configurationItemId"
              name="configuration_item_id"
              placeholder="Select a CI"
              inputs={cis}
              invalid={
                !!state.errors?.configurationItemId &&
                state.errors.configurationItemId.length > 0
              }
              errorMessage={state.errors?.configurationItemId?.join(", ")}
            />

            {/* Short Description */}
            <ShortDescriptionInput
              invalid={
                !!state.errors?.shortDescription &&
                state.errors.shortDescription.length > 0
              }
              errorMessage={state.errors?.shortDescription?.join(", ") || ""}
            />

            {/* Description */}
            <DescriptionTextarea
              invalid={
                !!state.errors?.shortDescription &&
                state.errors.shortDescription.length > 0
              }
              errorMessage={state.errors?.shortDescription?.join(", ") || ""}
            />

            {/* Incident State */}
            <StateListbox
              invalid={!!state.errors?.state && state.errors.state.length > 0}
              errorMessage={state.errors?.state?.join(", ") || ""}
            />
          </FieldGroup>

          {/* Message Area */}
          <MessageArea state={state} />
        </Fieldset>

        <Divider className="my-10" soft />

        <div className="flex justify-end gap-4">
          <Button type="reset" plain>
            Reset
          </Button>
          <Button type="submit">Save</Button>
        </div>
      </form>
    </FormWrapper>
  );
}

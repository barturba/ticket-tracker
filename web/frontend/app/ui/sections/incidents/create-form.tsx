"use client";
import { Button } from "@/app/components/button";
import { createIncident, IncidentState } from "@/app/api/incidents/incidents";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import FormInput from "@/app/application-components/resources/form-input";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
import DescriptionTextarea from "@/app/application-components/incident/description-textarea";
import StateListbox from "@/app/application-components/incident/state-listbox";
import { Divider } from "@/app/components/divider";
import { useActionState } from "react";
// import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";
import { CIField } from "@/app/api/cis/cis.d";
import { CompanyField } from "@/app/api/companies/companies.d";
import { User } from "@/types/users/base";

export default function CreateIncidentForm({
  companies,
  users,
  cis,
}: {
  companies: CompanyField[];
  users: User[];
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
              inputs={cis}
              invalid={
                !!state.errors?.configurationItemId &&
                state.errors.configurationItemId.length > 0
              }
              errorMessage={state.errors?.configurationItemId?.join(", ")}
            />

            {/* Short Description */}
            <ShortDescriptionInput
              label="Short Description"
              name="short_description"
              invalid={
                !!state.errors?.shortDescription &&
                state.errors.shortDescription.length > 0
              }
              errorMessage={state.errors?.shortDescription?.join(", ") || ""}
            />

            {/* Description */}
            <DescriptionTextarea
              label="Description"
              name="description"
              description="Provide a detailed description of the incident"
              invalid={
                !!state.errors?.description &&
                state.errors.description.length > 0
              }
              errorMessage={state.errors?.description?.join(", ") || ""}
              defaultValue={""}
            />

            {/* Incident State */}
            <StateListbox
              invalid={!!state.errors?.state && state.errors.state.length > 0}
              errorMessage={state.errors?.state?.join(", ") || ""}
            />
          </FieldGroup>

          {/* Message Area */}
          {/* <MessageArea state={state} /> */}
        </Fieldset>

        <Divider className="my-10" soft />

        <div className="flex justify-end gap-4">
          <Button type="reset" plain>
            Reset
          </Button>
          <SubmitButton isPending={false} />
        </div>
      </form>
    </FormWrapper>
  );
}

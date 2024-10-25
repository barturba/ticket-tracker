"use client";
import { Button } from "@/app/components/button";
import { useActionState } from "react";
import { CompanyState, updateCompany } from "@/app/api/companies/companies";
import { Divider } from "@/app/components/divider";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
// import MessageArea from "@/app/application-components/resources/message-area";
import { CompanyForm } from "@/app/api/companies/companies.d";

export default function EditCompanyForm({ company }: { company: CompanyForm }) {
  const initialState: CompanyState = { message: "", errors: {} };

  const updateCompanyWithId = updateCompany.bind(null, company.id);
  const [state, formAction] = useActionState(updateCompanyWithId, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="Company details">
          <FieldGroup>
            {/* Name */}
            <ShortDescriptionInput
              label="Name"
              name="name"
              defaultValue={company.name}
              invalid={!!state.errors?.name && state.errors.name.length > 0}
              errorMessage={state.errors?.name?.join(", ") || ""}
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

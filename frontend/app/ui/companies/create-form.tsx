"use client";
import { Button } from "@/app/components/button";
import { CompanyState, createCompany } from "@/app/lib/actions/companies";
import { UserField } from "@/app/lib/definitions/users";
import { CIField } from "@/app/lib/definitions/cis";
import { CompanyField } from "@/app/lib/definitions/companies";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import { Divider } from "@/app/components/divider";
import { useActionState } from "react";
import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";

export default function CreateCompanyForm({
  companies,
  users,
  cis,
}: {
  companies: CompanyField[];
  users: UserField[];
  cis: CIField[];
}) {
  const initialState: CompanyState = { message: "", errors: {} };
  const [state, formAction] = useActionState(createCompany, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="Company details">
          <FieldGroup>
            {/* Name */}
            <ShortDescriptionInput
              label="Name"
              name="name"
              invalid={
                !!state.errors?.shortDescription &&
                state.errors.shortDescription.length > 0
              }
              errorMessage={state.errors?.shortDescription?.join(", ") || ""}
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
          <SubmitButton />
        </div>
      </form>
    </FormWrapper>
  );
}

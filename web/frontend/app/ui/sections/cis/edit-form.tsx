"use client";
import { Button } from "@/app/components/button";
import { useActionState } from "react";
import { updateCI } from "@/app/api/cis/cis";
import { Divider } from "@/app/components/divider";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import { CIState } from "@/app/api/cis/cis";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
import { CIForm } from "@/app/api/cis/cis.d";

export default function EditCIForm({ ci }: { ci: CIForm }) {
  const initialState: CIState = { message: "", errors: {} };
  const updateCIWithId = updateCI.bind(null, ci.id);
  const [state, formAction] = useActionState(updateCIWithId, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="CI details">
          <FieldGroup>
            {/* Name */}
            <ShortDescriptionInput
              label="Name"
              name="name"
              defaultValue={ci.name}
              invalid={!!state.errors?.name && state.errors.name.length > 0}
              errorMessage={state.errors?.name?.join(", ") || ""}
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

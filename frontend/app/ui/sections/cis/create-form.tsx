"use client";
import { Button } from "@/app/components/button";
import { createCI, CIState } from "@/app/lib/actions/cis";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import { Divider } from "@/app/components/divider";
import { useActionState } from "react";
import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";

export default function CreateCIForm() {
  const initialState: CIState = { message: "", errors: {} };
  const [state, formAction] = useActionState(createCI, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="CI details">
          <FieldGroup>
            {/* Name */}
            <ShortDescriptionInput
              label="Name"
              name="name"
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

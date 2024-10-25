"use client";
import { Button } from "@/app/components/button";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import { Divider } from "@/app/components/divider";
import { useActionState } from "react";
// import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
import { createUser } from "@/app/api/users/mutations";
import { UserFormState } from "@/types/users/base";

export default function CreateUserForm() {
  const initialState: UserFormState = { message: "", errors: {} };
  const [state, formAction] = useActionState(createUser, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="User details">
          <FieldGroup>
            {/* First Name*/}
            <ShortDescriptionInput
              label="First Name"
              name="first_name"
              invalid={
                !!state.errors?.first_name && state.errors.first_name.length > 0
              }
              errorMessage={state.errors?.first_name?.join(", ") || ""}
            />
            {/* Last Name*/}
            <ShortDescriptionInput
              label="Last Name"
              name="last_name"
              invalid={
                !!state.errors?.last_name && state.errors.last_name.length > 0
              }
              errorMessage={state.errors?.last_name?.join(", ") || ""}
            />
            {/* Email */}
            <ShortDescriptionInput
              label="Email"
              name="email"
              invalid={!!state.errors?.email && state.errors.email.length > 0}
              errorMessage={state.errors?.email?.join(", ") || ""}
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

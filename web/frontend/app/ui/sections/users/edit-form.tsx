"use client";
import { Button } from "@/app/components/button";
import { useActionState } from "react";
import { updateUser } from "@/app/api/users/users";
import { Divider } from "@/app/components/divider";
import { UserForm } from "@/app/api/users/users.d";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import { UserState } from "@/app/api/users/users";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import MessageArea from "@/app/application-components/resources/message-area";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";

export default function EditUserForm({ user }: { user: UserForm }) {
  const initialState: UserState = { message: "", errors: {} };
  const updateUserWithId = updateUser.bind(null, user.id);
  const [state, formAction] = useActionState(updateUserWithId, initialState);

  return (
    <FormWrapper subheading="Summary">
      <form action={formAction}>
        <Fieldset aria-label="User details">
          <FieldGroup>
            {/* First Name*/}
            <ShortDescriptionInput
              label="First Name"
              name="first_name"
              defaultValue={user.first_name}
              invalid={
                !!state.errors?.first_name && state.errors.first_name.length > 0
              }
              errorMessage={state.errors?.first_name?.join(", ") || ""}
            />
            {/* Last Name*/}
            <ShortDescriptionInput
              label="Last Name"
              name="last_name"
              defaultValue={user.last_name}
              invalid={
                !!state.errors?.last_name && state.errors.last_name.length > 0
              }
              errorMessage={state.errors?.last_name?.join(", ") || ""}
            />
            {/* Email */}
            <ShortDescriptionInput
              label="Email"
              name="email"
              defaultValue={user.email}
              invalid={!!state.errors?.email && state.errors.email.length > 0}
              errorMessage={state.errors?.email?.join(", ") || ""}
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

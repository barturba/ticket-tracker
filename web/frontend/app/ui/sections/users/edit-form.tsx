"use client";
import { Button } from "@/app/components/button";
import { useState, useTransition } from "react";
import { Divider } from "@/app/components/divider";
import { FieldGroup, Fieldset } from "@/app/components/fieldset";
import FormWrapper from "@/app/application-components/resources/form-wrapper";
import SubmitButton from "@/app/application-components/resources/button-submit";
import ShortDescriptionInput from "@/app/application-components/incident/short-description-input";
import { User } from "@/types/users/base";
import { updateUser } from "@/app/api/users/mutations";

interface EditUserFormProps {
  user: User;
}

export default function EditUserForm({ user }: { user: EditUserFormProps }) {
  const [isPending, startTransition] = useTransition();
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (formData: FormData) => {
    setError(null);
    startTransition(async () => {
      try {
        await updateUser(user.user.id, {}, formData);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "An unexpected error occurred."
        );
      }
    });
  };

  return (
    <FormWrapper subheading="Summary">
      <form action={handleSubmit} space-y-6>
        {error && (
          <div className="rounded-md bg-red-50 p-4">
            <p className="text-sm text-red-700">{error}</p>
          </div>
        )}
        <Fieldset aria-label="User details">
          <FieldGroup>
            {/* First Name*/}
            <ShortDescriptionInput
              disabled={isPending}
              label="First Name"
              name="first_name"
              defaultValue={user.first_name}
              invalid={false}
              errorMessage={""}
            />
            {/* Last Name*/}
            <ShortDescriptionInput
              disabled={isPending}
              label="Last Name"
              name="last_name"
              defaultValue={user.last_name}
              invalid={false}
              errorMessage={""}
            />
            {/* Email */}
            <ShortDescriptionInput
              disabled={isPending}
              label="Email"
              name="email"
              defaultValue={user.email}
              invalid={false}
              errorMessage={""}
            />
          </FieldGroup>
        </Fieldset>

        <Divider className="my-10" soft />

        <div className="flex justify-end gap-4">
          <Button type="reset" plain>
            Reset
          </Button>
          <SubmitButton isPending={isPending} />
        </div>
      </form>
    </FormWrapper>
  );
}

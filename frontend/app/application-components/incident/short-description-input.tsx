import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Input } from "@/app/components/input";

export default function ShortDescriptionInput({
  defaultValue,
  invalid,
  errorMessage,
}: {
  defaultValue?: string;
  invalid: boolean;
  errorMessage: string;
}) {
  return (
    <Field>
      <Label>Short description</Label>
      <Input
        name="short_description"
        defaultValue={defaultValue}
        invalid={!!invalid}
      />

      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

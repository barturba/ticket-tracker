import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Input } from "@/app/components/input";

export default function ShortDescriptionInput({
  label,
  name,
  defaultValue,
  invalid,
  errorMessage,
}: {
  label: string;
  name: string;
  defaultValue?: string;
  invalid: boolean;
  errorMessage: string;
}) {
  return (
    <Field>
      <Label>{label}</Label>
      <Input name={name} defaultValue={defaultValue} invalid={!!invalid} />

      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

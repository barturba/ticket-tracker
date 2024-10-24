import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Input } from "@/app/components/input";

export default function ShortDescriptionInput({
  disabled,
  label,
  name,
  defaultValue,
  invalid,
  errorMessage,
}: {
  disabled?: boolean;
  label: string;
  name: string;
  defaultValue?: string;
  invalid: boolean;
  errorMessage: string;
}) {
  return (
    <Field>
      <Label>{label}</Label>
      <Input
        disabled={disabled}
        name={name}
        defaultValue={defaultValue}
        invalid={!!invalid}
      />

      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

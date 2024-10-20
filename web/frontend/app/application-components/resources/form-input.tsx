import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Select } from "@/app/components/select";

export default function FormInput({
  label,
  id,
  name,
  inputs,
  defaultValue,
  invalid,
  errorMessage,
}: {
  label: string;
  id: string;
  name: string;
  inputs: { id: string; name: string }[];
  defaultValue?: string;
  invalid?: boolean;
  errorMessage?: string;
}) {
  return (
    <Field>
      <Label>{label}</Label>
      <Select
        id={id}
        name={name}
        defaultValue={defaultValue}
        invalid={!!invalid}
      >
        {inputs.map((input) => (
          <option key={input.id} value={input.id}>
            {input.name}
          </option>
        ))}
      </Select>
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

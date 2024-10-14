import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Select } from "@/app/components/select";
import { State } from "@/app/lib/actions";

export default function FormInput({
  label,
  id,
  name,
  placeholder,
  inputs,
  defaultValue,
  invalid,
  errorMessage,
}: {
  label: string;
  id: string;
  name: string;
  placeholder: string;
  inputs: any[];
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
        <option value="" disabled>
          {placeholder}&hellip;
        </option>
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

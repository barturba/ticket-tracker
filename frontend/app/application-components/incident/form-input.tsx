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
  state,
}: {
  label: string;
  id: string;
  name: string;
  placeholder: string;
  inputs: any[];
  defaultValue?: string;
  state: State;
}) {
  return (
    <Field>
      <Label>{label}</Label>
      <Select
        id={id}
        name={name}
        invalid={
          state.errors?.shortDescription &&
          state.errors.shortDescription.length > 0
        }
      >
        <option value="" disabled>
          {placeholder}
        </option>
        {inputs.map((input) => (
          <option key={input.id} value={input.id}>
            {input.name}
          </option>
        ))}
      </Select>
    </Field>
  );
}

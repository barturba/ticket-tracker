import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Select } from "@/app/components/select";

export default function StateListbox({
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
      <Label>State</Label>
      <Select name="state" defaultValue={defaultValue}>
        <option value="New">New</option>
        <option value="Assigned">Assigned</option>
        <option value="In Progress">In Progress</option>
        <option value="On Hold">On Hold</option>
        <option value="Resolved">Resolved</option>
      </Select>
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

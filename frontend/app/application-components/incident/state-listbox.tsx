import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Listbox, ListboxLabel, ListboxOption } from "@/app/components/listbox";
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
      {/* <Listbox name="state" defaultValue={defaultValue}>
        <ListboxOption value="New">
          <ListboxLabel>New</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="Assigned">
          <ListboxLabel>Assigned</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="In Progress">
          <ListboxLabel>In Progress</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="On Hold">
          <ListboxLabel>On Hold</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="Resolved">
          <ListboxLabel>Resolved</ListboxLabel>
        </ListboxOption>
      </Listbox> */}
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

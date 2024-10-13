import { ErrorMessage, Field, Label } from "@/app/components/fieldset";
import { Listbox, ListboxLabel, ListboxOption } from "@/app/components/listbox";

export default function StateListbox({
  defaultValue,
  invalid,
  errorMessage,
}: {
  defaultValue: string;
  invalid: boolean;
  errorMessage: string;
}) {
  return (
    <Field>
      <Label>State</Label>
      <Listbox name="state" defaultValue={defaultValue}>
        <ListboxOption value="New">
          <ListboxLabel>New</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="In Progress">
          <ListboxLabel>In Progress</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="Assigned">
          <ListboxLabel>Assigned</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="On Hold">
          <ListboxLabel>On Hold</ListboxLabel>
        </ListboxOption>
        <ListboxOption value="Resolved">
          <ListboxLabel>Resolved</ListboxLabel>
        </ListboxOption>
      </Listbox>
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

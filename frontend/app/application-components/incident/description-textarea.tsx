import {
  Description,
  ErrorMessage,
  Field,
  Label,
} from "@/app/components/fieldset";
import { Textarea } from "@/app/components/textarea";

export default function DescriptionTextarea({
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
      <Label>Description</Label>
      <Textarea name="description" defaultValue={defaultValue} />
      <Description>Provide a detailed description of the incident</Description>
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

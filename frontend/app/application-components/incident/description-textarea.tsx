import {
  Description,
  ErrorMessage,
  Field,
  Label,
} from "@/app/components/fieldset";
import { Textarea } from "@/app/components/textarea";

export default function DescriptionTextarea({
  label,
  name,
  description,
  defaultValue,
  invalid,
  errorMessage,
}: {
  label: string;
  name: string;
  description: string;
  defaultValue: string;
  invalid: boolean;
  errorMessage: string;
}) {
  return (
    <Field>
      <Label>{label}</Label>
      <Textarea name={name} defaultValue={defaultValue} />
      <Description>{description}</Description>
      {!!invalid && <ErrorMessage>{errorMessage}</ErrorMessage>}
    </Field>
  );
}

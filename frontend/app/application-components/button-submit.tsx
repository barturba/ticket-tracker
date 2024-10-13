import { useFormStatus } from "react-dom";
import { Button } from "@/app/components/button";

export default function SubmitButton() {
  const { pending } = useFormStatus();

  return (
    <Button type="submit" aria-disabled={pending}>
      Save Changes
    </Button>
  );
}

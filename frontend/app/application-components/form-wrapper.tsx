import { Divider } from "@/app/components/divider";
import { Subheading } from "@/app/components/heading";

export default function FormWrapper({
  children,
  subheading,
}: {
  children: React.ReactNode;
  subheading: string;
}) {
  return (
    <div className="mt-12">
      <Subheading>{subheading}</Subheading>
      <Divider className="mt-4" />
      {children}
    </div>
  );
}

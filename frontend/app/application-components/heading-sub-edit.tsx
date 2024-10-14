import { Badge } from "@/app/components/badge";
import { Heading } from "@/app/components/heading";

export default function HeadingSubEdit({
  name,
  badgeState,
  badgeText,
}: {
  name: string;
  badgeState:
    | "New"
    | "Assigned"
    | "In Progress"
    | "On Hold"
    | "Resolved"
    | undefined;
  badgeText: string | undefined;
}) {
  return (
    <div className="mt-4 lg:mt-8">
      <div className="flex items-center gap-4">
        <Heading>{name}</Heading>
        {badgeState && badgeText && (
          <Badge state={badgeState}>{badgeText}</Badge>
        )}
      </div>
    </div>
  );
}

import { ChevronLeftIcon } from "@heroicons/react/24/outline";
import { Link } from "../components/link";

export default function HeadingEdit({
  name,
  backLink,
}: {
  name: string;
  backLink: string;
}) {
  return (
    <div className="max-lg:hidden">
      <Link
        href={backLink}
        className="inline-flex items-center gap-2 text-sm/6 text-zinc-500 dark:text-zinc-400"
      >
        <ChevronLeftIcon className="size-4 fill-zinc-400 dark:fill-zinc-500" />
        {name}
      </Link>
    </div>
  );
}

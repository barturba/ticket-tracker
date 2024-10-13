import { Heading } from "@/app/components/heading";
import { lusitana } from "@/app/ui/fonts";
import { ChevronRightIcon, HomeIcon } from "@heroicons/react/24/outline";
import clsx from "clsx";
import Link from "next/link";

interface Breadcrumb {
  label: string;
  href: string;
  active?: boolean;
}

export default function Breadcrumbs({
  breadcrumbs,
}: {
  breadcrumbs: Breadcrumb[];
}) {
  return (
    <nav aria-label="Breadcrumb" className="flex">
      <ol role="list" className="flex items-center space-x-4">
        <li>
          <div>
            <a href="#" className="text-gray-400 hover:text-gray-500">
              <HomeIcon aria-hidden="true" className="h-5 w-5 flex-shrink-0" />
              <span className="sr-only">Home</span>
            </a>
          </div>
        </li>
        {breadcrumbs.map((breadcrumb, index) => (
          <li
            key={breadcrumb.label}
            aria-current={breadcrumb.active}
            className={clsx(breadcrumb.active ? "text-gray-400" : "text-white")}
          >
            <div className="flex items-center">
              <ChevronRightIcon
                aria-hidden="true"
                className="h-5 w-5 flex-shrink-0 text-gray-400"
              />
              <Heading level={2} className="text-gray-400">
                {breadcrumb.label}
              </Heading>
            </div>
          </li>
        ))}
      </ol>
    </nav>

    // <nav aria-label="Breadcrumb" className="flex">
    //   <ol className={clsx("flex text-xl md:text-2xl")}>
    //     {breadcrumbs.map((breadcrumb, index) => (
    //       <li
    //         key={breadcrumb.href}
    //         aria-current={breadcrumb.active}
    //         className={clsx(breadcrumb.active ? "text-zinc-950" : "text-white")}
    //       >
    //         <Navbar href={breadcrumb.href}>{breadcrumb.label}</>
    //         {index < breadcrumbs.length - 1 ? (
    //           <span className="mx-3 inline-block">/</span>
    //         ) : null}
    //       </li>
    //     ))}
    //   </ol>
    // </nav>
  );
}

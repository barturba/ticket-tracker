"use client";
import {
  BriefcaseIcon,
  CpuChipIcon,
  DocumentDuplicateIcon,
  HomeIcon,
  UserGroupIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { usePathname } from "next/navigation";
import clsx from "clsx";
import { classNames } from "@/app/lib/utils";
import { Links } from "@/app/lib/definitions";

export default function NavLinksMobile() {
  const pathname = usePathname();
  return (
    <nav className="flex flex-col flex-1">
      <ul role="list" className="flex flex-col flex-1 gap-y-7">
        <li>
          <ul role="list" className="-mx-2 space-y-1">
            {Links.map((item) => (
              <li key={item.name}>
                <a
                  href={item.href}
                  className={classNames(
                    pathname === item.href
                      ? "bg-gray-50 text-indigo-600"
                      : "text-gray-700 hover:bg-gray-50 hover:text-indigo-600",
                    "group flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6"
                  )}
                >
                  <item.icon
                    aria-hidden="true"
                    className={classNames(
                      pathname === item.href
                        ? "text-indigo-600"
                        : "text-gray-400 group-hover:text-indigo-600",
                      "h-6 w-6 shrink-0"
                    )}
                  />
                  {item.name}
                </a>
              </li>
            ))}
          </ul>
        </li>
      </ul>
    </nav>
  );
}

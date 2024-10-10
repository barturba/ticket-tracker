"use client";
import {
  BriefcaseIcon,
  CpuChipIcon,
  DocumentDuplicateIcon,
  HomeIcon,
  UserGroupIcon,
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { Links } from "@/app/lib/definitions";
import { usePathname } from "next/navigation";
import clsx from "clsx";
import { classNames } from "@/app/lib/utils";

export default function NavLinks() {
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

        {/* <li className="mt-auto -mx-6">
          <a
            href="#"
            className="flex items-center px-6 py-3 text-sm font-semibold leading-6 text-gray-900 gap-x-4 hover:bg-gray-50"
          >
            <img
              alt=""
              src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
              className="w-8 h-8 rounded-full bg-gray-50"
            />
            <span className="sr-only">Your profile</span>
            <span aria-hidden="true">Tom Cook</span>
          </a>
        </li> */}
      </ul>
    </nav>
  );
}

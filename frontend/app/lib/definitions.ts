import {
  BriefcaseIcon,
  CpuChipIcon,
  DocumentDuplicateIcon,
  HomeIcon,
  UserGroupIcon,
} from "@heroicons/react/24/outline";

export const Links = [
  {
    name: "Dashboard",
    href: "/dashboard",
    icon: HomeIcon,
  },
  {
    name: "Incidents",
    href: "/dashboard/incidents",
    icon: DocumentDuplicateIcon,
  },
  { name: "Companies", href: "/dashboard/companies", icon: BriefcaseIcon },
  { name: "Users", href: "/dashboard/users", icon: UserGroupIcon },
  {
    name: "Configuration Items",
    href: "/dashboard/cis",
    icon: CpuChipIcon,
  },
];

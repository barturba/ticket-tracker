import { MagnifyingGlassIcon } from "@heroicons/react/24/outline";
import { Input, InputGroup } from "@/app/components/input";
import { Heading } from "@/app/components/heading";
import { Select } from "@/app/components/select";
import { Button } from "@/app/components/button";
import { html } from "framer-motion/client";

export default function AppHeading({
  name,
  createLink,
}: {
  name: string;
  createLink: string;
}) {
  return (
    <div className="flex flex-wrap items-end justify-between gap-4">
      <div className="max-sm:w-full sm:flex-1">
        <Heading>{name}</Heading>
        {/* <div className="mt-4 flex max-w-xl gap-4">
          <div className="flex-1">
            <InputGroup>
              <MagnifyingGlassIcon />
              <Input name="search" placeholder={`Search ${name}...`} />
            </InputGroup>
          </div>
          <div>
            <Select name="sort_by">
              <option value="name">Sort by name</option>
              <option value="date">Sort by date</option>
              <option value="status">Sort by status</option>
            </Select>
          </div>
        </div> */}
      </div>
      <Button href={createLink}>Create {name}</Button>
    </div>
  );
}

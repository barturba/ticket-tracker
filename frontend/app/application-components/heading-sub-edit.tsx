import { BanknotesIcon, CalendarIcon } from "@heroicons/react/24/outline";
import { Badge } from "../components/badge";
import { Heading } from "../components/heading";
import { CreditCardIcon } from "@heroicons/react/24/outline";
import { Button } from "../components/button";
import { MenuHeadingProps } from "@headlessui/react";

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
      {/* <div className="isolate mt-2.5 flex flex-wrap justify-between gap-x-6 gap-y-4">
        <div className="flex flex-wrap gap-x-10 gap-y-4 py-1.5">
          <span className="flex items-center gap-3 text-base/6 text-zinc-950 sm:text-sm/6 dark:text-white">
            <BanknotesIcon className="size-4 shrink-0 fill-zinc-400 dark:fill-zinc-500" />
            <span>US{order.amount.usd}</span>
          </span>
          <span className="flex items-center gap-3 text-base/6 text-zinc-950 sm:text-sm/6 dark:text-white">
            <CreditCardIcon className="size-4 shrink-0 fill-zinc-400 dark:fill-zinc-500" />
            <span className="inline-flex gap-3">
              {order.payment.card.type}{" "}
              <span>
                <span aria-hidden="true">••••</span> {order.payment.card.number}
              </span>
            </span>
          </span>
          <span className="flex items-center gap-3 text-base/6 text-zinc-950 sm:text-sm/6 dark:text-white">
            <CalendarIcon className="size-4 shrink-0 fill-zinc-400 dark:fill-zinc-500" />
            <span>{order.date}</span>
          </span>
        </div>
        <div className="flex gap-4">
           <RefundOrder outline amount={order.amount.usd}>
            Refund
          </RefundOrder> 
          <Button>Resend Invoice</Button>
        </div>
      </div> */}
    </div>
  );
}

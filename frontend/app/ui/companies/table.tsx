import { fetchFilteredCompanies } from "@/app/lib/actions";
// import CompanyStatus from "@/app/ui/companies/status";
import { DeleteCompany, UpdateCompany } from "@/app/ui/companies/buttons";
import { Company } from "@/app/lib/definitions";
import { truncate } from "@/app/lib/utils";
export default async function CompanyTable({
  query,
  currentPage,
}: {
  query: string;
  currentPage: number;
}) {
  const companies: Company[] = await fetchFilteredCompanies(query, currentPage);
  return (
    <div className="mt-6 flow-root">
      <div className="inline-block min-w-full align-middle">
        <div className="rounded-lg bg-gray-50 p-2 md:pt-0">
          <div className="md:hidden">
            {companies?.map((company) => (
              <div
                key={company.id}
                className="mb-2 w-full rounded-md bg-white p-4"
              >
                <div className="flex items-center justify-between border-b pb-4">
                  <div>
                    <div className="mb-2 flex items-center">
                      <p>{company.id}</p>
                    </div>
                  </div>
                  {/* <CompanyStatus status={company.state} /> */}
                </div>
                <div className="flex w-full items-center justify-between pt-4">
                  <p className="text-xl font-medium">{company.created_at}</p>
                </div>
                <div className="flex justify-end gap-2">
                  <UpdateCompany id={company.id} />
                  <DeleteCompany id={company.id} />
                </div>
              </div>
            ))}
          </div>
          <table className="hidden min-w-full text-gray-900 md:table">
            <thead className="rounded-lg text-left text-sm font-normal">
              <tr>
                <th scope="col" className="px-4 py-5 font-medium sm:pl-6">
                  ID
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Name
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Date
                </th>
                <th scope="col" className="relative py-3 pl-6 pr-3">
                  <span className="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {companies?.map((company) => (
                <tr
                  key={company.id}
                  className="w-full border-b py-3 text-sm last-of-type:border-none [&:first-child>td:first-child]:rounded-tl-lg [&:first-child>td:last-child]:rounded-tr-lg [&:last-child>td:first-child]:rounded-bl-lg [&:last-child>td:last-child]:rounded-br-lg"
                >
                  <td className="whitespace-nowrap py-3 pl-6 pr-3">
                    <div className="flex items-center gap-3">
                      <p>{company.id}</p>
                    </div>
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {truncate(company.name, true)}
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {company.created_at}
                  </td>
                  <td className="whitespace-nowrap py-3 pl-6 pr-3">
                    <div className="flex justify-end gap-3">
                      <UpdateCompany id={company.id} />
                      <DeleteCompany id={company.id} />
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
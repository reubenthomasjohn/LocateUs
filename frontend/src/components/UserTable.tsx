import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { format } from "date-fns";
import { Pencil, Trash2 } from "lucide-react";
import { User } from "../store/atoms/users";

interface UserTableProps {
  users: User[];
  onEdit: (id: number) => void;
  onDelete: (id: number) => void;
}

const columnHelper = createColumnHelper<User>();

export function UserTable({ users, onEdit, onDelete }: UserTableProps) {
  const columns = [
    columnHelper.accessor("full_name", {
      header: "Name",
      cell: (info) => info.getValue(),
    }),
    // columnHelper.accessor("email", {
    //   header: "Email",
    //   cell: (info) => info.getValue(),
    // }),
    // columnHelper.accessor("date_of_birth", {
    //   header: "Date of Birth",
    //   cell: (info) => format(new Date(info.getValue()), "PP"),
    // }),
    columnHelper.accessor("phone_number", {
      header: "Phone",
      cell: (info) => info.getValue(),
    }),
    columnHelper.accessor("created_at", {
      header: "Joined",
      cell: (info) => format(new Date(info.getValue()), "PP"),
    }),
    columnHelper.display({
      id: "actions",
      header: "Actions",
      cell: (props) => (
        <div className="flex space-x-2">
          <button
            onClick={() => onEdit(props.row.original.id)}
            className="p-1 text-blue-600 hover:text-blue-800 transition-colors"
            title="Edit user"
          >
            <Pencil className="h-4 w-4" />
          </button>
          <button
            onClick={() => onDelete(Number(props.row.original.id))}
            className="p-1 text-red-600 hover:text-red-800 transition-colors"
            title="Delete user"
          >
            <Trash2 className="h-4 w-4" />
          </button>
        </div>
      ),
    }),
  ];

  const table = useReactTable({
    data: users,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div className="mt-8 flow-root">
      <div className="overflow-x-auto">
        <div className="inline-block min-w-full py-2 align-middle">
          <table className="min-w-full divide-y divide-gray-300">
            <thead>
              {table.getHeaderGroups().map((headerGroup) => (
                <tr key={headerGroup.id}>
                  {headerGroup.headers.map((header) => (
                    <th
                      key={header.id}
                      className="px-6 py-3 text-left text-sm font-semibold text-gray-900 bg-gray-50"
                    >
                      {flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                    </th>
                  ))}
                </tr>
              ))}
            </thead>
            <tbody className="divide-y divide-gray-200 bg-white">
              {table.getRowModel().rows.map((row) => (
                <tr key={row.id} className="hover:bg-gray-50">
                  {row.getVisibleCells().map((cell) => (
                    <td
                      key={cell.id}
                      className="whitespace-nowrap px-6 py-4 text-sm text-gray-500"
                    >
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </td>
                  ))}
                </tr>
              ))}
            </tbody>
          </table>
          {users.length === 0 && (
            <div className="text-center py-8 text-gray-500">No users found</div>
          )}
        </div>
      </div>
    </div>
  );
}

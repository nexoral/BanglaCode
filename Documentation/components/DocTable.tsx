import React from "react";

interface Column {
  header: string;
  key: string;
  render?: (value: unknown, row: Record<string, unknown>) => React.ReactNode;
}

interface DocTableProps {
  columns: Column[];
  data: Record<string, unknown>[];
}

export default function DocTable({ columns, data }: DocTableProps) {
  return (
    <div className="overflow-x-auto my-6 rounded-lg border border-border">
      <table className="min-w-full">
        <thead className="bg-secondary/80">
          <tr>
            {columns.map((col) => (
              <th
                key={col.key}
                className="px-4 py-3 text-left font-semibold text-foreground whitespace-nowrap"
              >
                {col.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="divide-y divide-border">
          {data.map((row, idx) => (
            <tr key={idx} className="hover:bg-secondary/30 transition-colors">
              {columns.map((col) => (
                <td key={col.key} className="px-4 py-3 text-muted-foreground">
                  {col.render ? col.render(row[col.key], row) : String(row[col.key] ?? "")}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

// Simple table wrapper for manually written tables
export function TableWrapper({ children }: { children: React.ReactNode }) {
  return (
    <div className="overflow-x-auto my-6 rounded-lg border border-border">
      {children}
    </div>
  );
}

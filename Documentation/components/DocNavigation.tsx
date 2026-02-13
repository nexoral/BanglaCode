"use client";

import Link from "next/link";
import { ArrowLeft, ArrowRight } from "lucide-react";
import { getNavigation } from "@/lib/docs-config";

interface DocNavigationProps {
  currentPath: string;
}

export default function DocNavigation({ currentPath }: DocNavigationProps) {
  const { prev, next } = getNavigation(currentPath);

  if (!prev && !next) return null;

  return (
    <div className="mt-16 pt-8 border-t border-border">
      <nav className="flex justify-between gap-4">
        {prev ? (
          <Link
            href={prev.href}
            className="group flex-1 flex flex-col gap-1 p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-secondary/30 transition-colors"
          >
            <span className="text-xs text-muted-foreground flex items-center gap-1">
              <ArrowLeft className="w-3 h-3" /> Previous
            </span>
            <span className="font-medium group-hover:text-primary transition-colors">
              {prev.name}
            </span>
            {prev.description && (
              <span className="text-sm text-muted-foreground">{prev.description}</span>
            )}
          </Link>
        ) : (
          <div className="flex-1" />
        )}

        {next ? (
          <Link
            href={next.href}
            className="group flex-1 flex flex-col gap-1 p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-secondary/30 transition-colors text-right"
          >
            <span className="text-xs text-muted-foreground flex items-center gap-1 justify-end">
              Next <ArrowRight className="w-3 h-3" />
            </span>
            <span className="font-medium group-hover:text-primary transition-colors">
              {next.name}
            </span>
            {next.description && (
              <span className="text-sm text-muted-foreground">{next.description}</span>
            )}
          </Link>
        ) : (
          <div className="flex-1" />
        )}
      </nav>
    </div>
  );
}

"use client";

import { useState } from "react";
import Sidebar, { MobileSidebarToggle } from "@/components/Sidebar";

export default function DocsLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  return (
    <div className="flex flex-col lg:flex-row min-h-[calc(100vh-4rem)]">
      {/* Sidebar */}
      <Sidebar
        isMobileOpen={isMobileMenuOpen}
        onMobileClose={() => setIsMobileMenuOpen(false)}
      />

      {/* Main content */}
      <div className="lg:pl-64 flex-1">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 py-8 sm:py-12">
          <article className="prose prose-invert prose-purple max-w-none
            prose-headings:scroll-mt-20
            prose-h1:text-3xl prose-h1:font-bold prose-h1:mb-4
            prose-h2:text-2xl prose-h2:font-semibold prose-h2:mt-10 prose-h2:mb-4 prose-h2:border-b prose-h2:border-border prose-h2:pb-2
            prose-h3:text-xl prose-h3:font-medium prose-h3:mt-8 prose-h3:mb-3
            prose-p:text-muted-foreground prose-p:leading-7
            prose-a:text-primary prose-a:no-underline hover:prose-a:underline
            prose-code:text-primary prose-code:bg-secondary/50 prose-code:px-1.5 prose-code:py-0.5 prose-code:rounded prose-code:before:content-none prose-code:after:content-none
            prose-pre:bg-[#1e1e1e] prose-pre:border prose-pre:border-border
            prose-li:text-muted-foreground
            prose-strong:text-foreground
            [&_table]:w-full [&_table]:border-collapse [&_table]:my-6 [&_table]:rounded-lg [&_table]:overflow-hidden [&_table]:border [&_table]:border-border
            [&_thead]:bg-secondary/80
            [&_th]:px-4 [&_th]:py-3 [&_th]:text-left [&_th]:font-semibold [&_th]:text-foreground [&_th]:border-b [&_th]:border-border
            [&_td]:px-4 [&_td]:py-3 [&_td]:text-muted-foreground [&_td]:border-b [&_td]:border-border
            [&_tr]:transition-colors [&_tbody_tr:hover]:bg-secondary/30
            [&_tr:last-child_td]:border-b-0
            [&_td_code]:text-sm [&_td_code]:bg-secondary/50 [&_td_code]:px-1.5 [&_td_code]:py-0.5 [&_td_code]:rounded [&_td_code]:text-primary
          ">
            {children}
          </article>
        </div>
      </div>

      {/* Mobile menu toggle */}
      <MobileSidebarToggle onClick={() => setIsMobileMenuOpen(true)} />
    </div>
  );
}

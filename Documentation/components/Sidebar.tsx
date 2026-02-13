"use client";

import { useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import clsx from "clsx";
import { Menu, X, ChevronDown, ChevronRight } from "lucide-react";
import { DOCS_CONFIG } from "@/lib/docs-config";

interface SidebarProps {
  isMobileOpen?: boolean;
  onMobileClose?: () => void;
}

export default function Sidebar({ isMobileOpen, onMobileClose }: SidebarProps) {
  const pathname = usePathname();
  const [expandedSections, setExpandedSections] = useState<Record<string, boolean>>(() => {
    // Initially expand the section containing the current page
    const expanded: Record<string, boolean> = {};
    DOCS_CONFIG.forEach(section => {
      const hasCurrentPage = section.items.some(item => item.href === pathname);
      expanded[section.section] = hasCurrentPage || section.section === "Getting Started";
    });
    return expanded;
  });

  const toggleSection = (section: string) => {
    setExpandedSections(prev => ({
      ...prev,
      [section]: !prev[section]
    }));
  };

  const sidebarContent = (
    <div className="p-4 space-y-2">
      {DOCS_CONFIG.map((section) => {
        const isExpanded = expandedSections[section.section];
        const hasActivePage = section.items.some(item => item.href === pathname);

        return (
          <div key={section.section} className="mb-1">
            <button
              onClick={() => toggleSection(section.section)}
              className={clsx(
                "w-full flex items-center justify-between gap-2 px-3 py-2 rounded-lg text-sm font-medium transition-colors",
                hasActivePage
                  ? "bg-primary/10 text-primary"
                  : "text-muted-foreground hover:bg-secondary hover:text-foreground"
              )}
            >
              <div className="flex items-center gap-2">
                <section.icon className="w-4 h-4" />
                <span>{section.section}</span>
              </div>
              {isExpanded ? (
                <ChevronDown className="w-4 h-4" />
              ) : (
                <ChevronRight className="w-4 h-4" />
              )}
            </button>

            {isExpanded && (
              <ul className="mt-1 ml-4 space-y-1 border-l border-border pl-3">
                {section.items.map((item) => {
                  const isActive = pathname === item.href;
                  return (
                    <li key={item.href}>
                      <Link
                        href={item.href}
                        onClick={onMobileClose}
                        className={clsx(
                          "block text-sm py-1.5 px-2 rounded-md transition-all",
                          isActive
                            ? "bg-primary/10 text-primary font-medium"
                            : "text-muted-foreground hover:text-foreground hover:bg-secondary/50"
                        )}
                      >
                        {item.name}
                      </Link>
                    </li>
                  );
                })}
              </ul>
            )}
          </div>
        );
      })}
    </div>
  );

  return (
    <>
      {/* Desktop Sidebar */}
      <aside className="fixed top-16 left-0 w-64 h-[calc(100vh-4rem)] border-r border-border bg-background/95 backdrop-blur-sm hidden lg:block overflow-y-auto scrollbar-thin">
        {sidebarContent}
      </aside>

      {/* Mobile Sidebar Overlay */}
      {isMobileOpen && (
        <div className="fixed inset-0 z-50 lg:hidden">
          {/* Backdrop */}
          <div
            className="fixed inset-0 bg-black/50 backdrop-blur-sm"
            onClick={onMobileClose}
          />

          {/* Sidebar */}
          <aside className="fixed top-0 left-0 w-72 h-full bg-background border-r border-border overflow-y-auto z-50">
            <div className="flex items-center justify-between p-4 border-b border-border">
              <span className="font-semibold text-lg">Documentation</span>
              <button
                onClick={onMobileClose}
                className="p-2 rounded-lg hover:bg-secondary transition-colors"
              >
                <X className="w-5 h-5" />
              </button>
            </div>
            {sidebarContent}
          </aside>
        </div>
      )}
    </>
  );
}

// Mobile toggle button component
export function MobileSidebarToggle({ onClick }: { onClick: () => void }) {
  return (
    <button
      onClick={onClick}
      className="lg:hidden fixed bottom-4 right-4 z-40 p-3 bg-primary text-white rounded-full shadow-lg hover:bg-primary/90 transition-colors"
      aria-label="Toggle navigation"
    >
      <Menu className="w-6 h-6" />
    </button>
  );
}

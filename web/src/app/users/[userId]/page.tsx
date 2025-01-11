"use client";

import { useStore } from "@/store/store";
import Categories from "@/app/categories";
import {
  useCategories,
  useCategoryTransactions,
} from "@/hooks/categories-hooks";
import { useSummary } from "@/hooks/summary-hooks";
import DatesCombobox from "@/components/dates-combobox/dates-combobox";
import CategoriesPieChart from "@/components/categories-pie-chart/categories-pie-chart";
import CategoryTransactions from "@/components/category-transactions/category-transactions";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Badge } from "@/components/ui/badge";
import { Bell, BellDot } from "lucide-react";

import React from "react";
import NotificationBell from "@/components/notification-bell/notification-bell";

export default function UserPage() {
  const { setCurrentDate, currentCategory } = useStore();
  const { isPending, error, data: summary } = useSummary();
  const {
    isFetched: isCategoryFetched,
    error: isCategoryError,
    data: categories = [],
  } = useCategories();
  const categoryTransactions = useCategoryTransactions();

  if (isPending) return "loading...";

  if (error || isCategoryError) return "error!";

  return (
    <div>
      <div className="flex justify-between items-center h-16 p-4 border-b">
        <h2 className="text-2xl font-bold tracking-tight">Expenses Summary</h2>
        <div className="flex items-center space-x-4">
          <DropdownMenu>
            <DropdownMenuTrigger>
              <NotificationBell count={10} />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>Profile</DropdownMenuItem>
              <DropdownMenuItem>Billing</DropdownMenuItem>
              <DropdownMenuItem>Team</DropdownMenuItem>
              <DropdownMenuItem>Subscription</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          <p>{summary.userInfo?.email}</p>
        </div>
      </div>
      <div className="flex flex-col space-y-8 p-4">
        <DatesCombobox summary={summary} onChange={setCurrentDate} />
        <div className="grid grid-cols-[60%_40%] gap-4">
          {isCategoryFetched && (
            <div className="space-y-4">
              <Categories />
              <CategoriesPieChart categories={categories} />
            </div>
          )}
          {categoryTransactions.isLoading && "loading category transactions..."}
          {categoryTransactions.isFetched &&
            categoryTransactions.data &&
            currentCategory && (
              <Card>
                <CardHeader>
                  Transactions of category: {currentCategory.name}
                </CardHeader>
                <CardContent>
                  <CategoryTransactions data={categoryTransactions.data} />
                </CardContent>
              </Card>
            )}
        </div>
      </div>
    </div>
  );
}

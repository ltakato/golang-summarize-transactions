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
      <div className="flex justify-between h-16 items-center p-4 border-b">
        <h2 className="text-2xl font-bold tracking-tight">Expenses Summary</h2>
        <p>{summary.userInfo?.email}</p>
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

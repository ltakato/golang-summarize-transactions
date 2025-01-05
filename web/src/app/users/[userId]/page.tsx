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
    <div className="space-y-8">
      <div className="space-y-4">
        <h1>Welcome back, {summary.userInfo?.email}</h1>
        <DatesCombobox summary={summary} onChange={setCurrentDate} />
      </div>
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
  );
}

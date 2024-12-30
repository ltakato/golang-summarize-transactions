"use client";

import { useStore } from "@/store/store";
import Categories from "@/app/categories";
import { useCategories } from "@/hooks/categories-hooks";
import { useSummary } from "@/hooks/summary-hooks";
import DatesCombobox from "@/components/dates-combobox/dates-combobox";
import CategoriesPieChart from "@/components/categories-pie-chart/categories-pie-chart";

export default function UserPage() {
  const { setCurrentDate } = useStore();
  const { isPending, error, data: summary } = useSummary();
  const {
    isFetched: isCategoryFetched,
    error: isCategoryError,
    data: categories = [],
  } = useCategories();

  if (isPending) return "loading...";

  if (error || isCategoryError) return "error!";

  return (
    <>
      <DatesCombobox summary={summary} onChange={setCurrentDate} />
      {isCategoryFetched && (
        <>
          <Categories />
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <CategoriesPieChart categories={categories} />
          </div>
        </>
      )}
    </>
  );
}

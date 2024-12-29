"use client";

import { Combobox } from "@/components/ui/combobox";
import { useStore } from "@/store/store";
import Categories from "@/app/categories";
import { CharDataItem, PieChart } from "@/components/ui/pie-chart";
import { useCategories } from "@/hooks/categories-hooks";
import { useSummary } from "@/hooks/summary-hooks";
import { toBRLCurrencyString } from "@/lib/currency-helper";

export default function UserPage() {
  const { setCurrentDate } = useStore();
  const { isFetched: isCategoryFetched, data: categories = [] } =
    useCategories();
  const { isPending, error, data: summaryData } = useSummary();

  if (isPending) return "loading...";

  if (error) return "error!";

  const mappedData = summaryData.availableDates.map((date) => {
    return { label: date, value: date };
  });
  const mappedChartData: CharDataItem[] = categories.map((category) => ({
    label: category.name,
    value: category.totalAmount,
  }));

  return (
    <>
      <Combobox options={mappedData} onChange={setCurrentDate} />
      {isCategoryFetched && (
        <>
          <Categories />
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <PieChart
              title="Expenses by category"
              chartData={mappedChartData}
              valueFormatter={toBRLCurrencyString}
            />
          </div>
        </>
      )}
    </>
  );
}

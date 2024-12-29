import { Combobox } from "@/components/ui/combobox";
import { useStore } from "@/store/store";
import Categories from "@/app/categories";
import { CharDataItem, PieChart } from "@/components/ui/pie-chart";
import { useCategories } from "@/hooks/categories-hooks";
import { useSummary } from "@/hooks/summary-hooks";
import { useEffect } from "react";

export default function MainPage() {
  const { availableDates, setAvailableDates, currentDate, setCurrentDate } =
    useStore();
  const { data: categories = [] } = useCategories();
  const { isPending, error, data: summaryData } = useSummary();

  useEffect(() => {
    if (!summaryData) return;
    setAvailableDates(summaryData.availableDates);
  }, [summaryData, setAvailableDates]);

  if (isPending) return "loading...";

  if (error) return "error!";

  const mappedData = availableDates.map((date) => {
    return { label: date, value: date };
  });
  const chartValueFormatter = (value: number) =>
    value.toLocaleString("pt-BR", {
      style: "currency",
      currency: "BRL",
    });
  const mappedChartData: CharDataItem[] = categories.map((category) => ({
    label: category.name,
    value: category.totalAmount,
  }));

  return (
    <>
      <Combobox options={mappedData} onChange={setCurrentDate} />
      {currentDate && (
        <>
          <Categories />
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <PieChart
              title="Expenses by category"
              chartData={mappedChartData}
              valueFormatter={chartValueFormatter}
            />
          </div>
        </>
      )}
    </>
  );
}

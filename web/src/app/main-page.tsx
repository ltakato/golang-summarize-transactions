import { useQuery } from "@tanstack/react-query";
import { Summary } from "@/app/app-types";
import { Combobox } from "@/components/ui/combobox";
import { useStore } from "@/store/store";
import Categories from "@/app/categories";
import { CharDataItem, PieChart } from "@/components/ui/pie-chart";
import { useCategories } from "@/hooks/categories-hooks";

export default function MainPage() {
  const { availableDates, setAvailableDates, currentDate, setCurrentDate } =
    useStore();
  const { data: categories = [] } = useCategories();
  const { isPending, error } = useQuery<Summary>({
    queryKey: ["summary"],
    queryFn: async () => {
      const res = await fetch("http://localhost:8080/api/summary");
      const json: Summary = await res.json();
      setAvailableDates(json.availableDates);
      return json;
    },
  });

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

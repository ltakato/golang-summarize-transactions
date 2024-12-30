import { CharDataItem, PieChart } from "@/components/ui/pie-chart";
import { toBRLCurrencyString } from "@/lib/currency-helper";
import { Category } from "@/app/app-types";

type Props = {
  categories: Category[];
};

export default function CategoriesPieChart({ categories }: Props) {
  const mappedChartData: CharDataItem[] = categories.map((category) => ({
    label: category.name,
    value: category.totalAmount,
  }));

  return (
    <PieChart
      title="Expenses by category"
      chartData={mappedChartData}
      valueFormatter={toBRLCurrencyString}
    />
  );
}

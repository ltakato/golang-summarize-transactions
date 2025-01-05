import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Category } from "@/app/app-types";
import { toBRLCurrencyString } from "@/lib/currency-helper";

type Props = {
  category: Category;
  onClick?: (category: Category) => void;
};

export default function CategoryCard({ onClick, category }: Props) {
  const { name, totalAmount } = category;
  const handleClick = (): void => {
    if (!onClick) return;
    onClick(category);
  };

  return (
    <Card onClick={() => handleClick()}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">{name}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="font-bold">{toBRLCurrencyString(totalAmount)}</div>
      </CardContent>
    </Card>
  );
}

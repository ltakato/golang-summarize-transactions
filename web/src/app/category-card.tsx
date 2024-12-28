import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Category } from "@/app/app-types";

type Props = {
  category: Category;
};

export default function CategoryCard({
  category: { name, totalAmount },
}: Props) {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">{name}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">R$ {totalAmount}</div>
      </CardContent>
    </Card>
  );
}

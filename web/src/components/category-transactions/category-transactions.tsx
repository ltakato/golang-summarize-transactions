import { CategoryTransaction } from "@/app/app-types";
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { toBRLCurrencyString } from "@/lib/currency-helper";
import { DateFormats, formatDate } from "@/lib/date-helper";
import { sumArray } from "@/lib/math-helper";

type Props = {
  data: CategoryTransaction[];
};

export default function CategoryTransactions(props: Props) {
  const { data = [] } = props;
  const transactionsAmount = sumArray(data, (category) => category.amount);
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">Date</TableHead>
          <TableHead>Title</TableHead>
          <TableHead className="text-right">Amount</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {data.map((transaction, index) => (
          <TableRow key={index}>
            <TableCell className="font-medium">
              {formatDate(transaction.date, DateFormats.YYYYMMDD)}
            </TableCell>
            <TableCell>{transaction.title}</TableCell>
            <TableCell className="text-right">
              {toBRLCurrencyString(transaction.amount)}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={2}>Total</TableCell>
          <TableCell className="text-right">
            {toBRLCurrencyString(transactionsAmount)}
          </TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
}

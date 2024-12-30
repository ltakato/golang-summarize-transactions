import { Combobox } from "@/components/ui/combobox";
import { Summary } from "@/app/app-types";

type Props = {
  summary: Summary;
  onChange: (value: string | null) => void;
};

export default function DatesCombobox(props: Props) {
  const { summary, onChange } = props;
  const mappedData = summary.availableDates.map((date) => ({
    label: date,
    value: date,
  }));

  return <Combobox options={mappedData} onChange={onChange} />;
}

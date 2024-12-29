import { useQuery } from "@tanstack/react-query";
import { Summary } from "@/app/app-types";
import { Combobox } from "@/components/ui/combobox";
import { useStore } from "@/store/store";
import Categories from "@/app/categories";

export default function MainPage() {
  const { availableDates, setAvailableDates, currentDate, setCurrentDate } =
    useStore();
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

  return (
    <>
      <Combobox options={mappedData} onChange={setCurrentDate} />
      {currentDate && <Categories />}
    </>
  );
}

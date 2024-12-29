import { useQuery } from "@tanstack/react-query";
import { Summary } from "@/app/app-types";

export function useSummary() {
  return useQuery<Summary>({
    queryKey: ["summary"],
    queryFn: async () => {
      const res = await fetch(`http://${process.env.API_URL}/api/summary`);
      return res.json();
    },
  });
}

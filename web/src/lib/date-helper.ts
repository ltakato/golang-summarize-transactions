import { format } from "date-fns";

export enum DateFormats {
  YYYYMMDD = "yyyy-MM-dd",
}

export function formatDate(date: Date, formatStr: DateFormats) {
  return format(date, formatStr);
}

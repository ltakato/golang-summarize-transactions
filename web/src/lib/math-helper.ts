export function sumArray<T>(array: T[], numberGetter: (item: T) => number) {
  return array.reduce(
    (curr, transaction) => curr + numberGetter(transaction),
    0,
  );
}

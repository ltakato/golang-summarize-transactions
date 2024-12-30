export class TransactionsSummaryApiClient {
  private readonly baseUrlApi = `${process.env.NEXT_PUBLIC_API_URL}`;
  private readonly userId: string;

  constructor({ userId }: { userId: string }) {
    this.userId = userId;
  }

  async fetch(url: string) {
    const headers = new Headers({ "x-user-id": this.userId });
    const res = await fetch(url, { headers });
    return res.json();
  }

  async getSummary() {
    const url = `${this.baseUrlApi}/api/summary`;
    return this.fetch(url);
  }

  async getCategories(date: string) {
    const params = new URLSearchParams({ date });
    const url = `${process.env.NEXT_PUBLIC_API_URL}/api/categories?${params.toString()}`;
    return this.fetch(url);
  }
}

import { Notification } from "@/app/app-types";

type NotificationsResponse = {
  items: Notification[];
  unreadCount: number;
};

export class NotificationsApiClient {
  private readonly baseUrlApi = `${process.env.NEXT_PUBLIC_API_URL}`;
  private readonly userId: string;

  constructor({ userId }: { userId: string }) {
    this.userId = userId;
  }

  async getNotifications(): Promise<NotificationsResponse> {
    const url = `${this.baseUrlApi}/api/notifications`;
    const headers = new Headers({ "x-user-id": this.userId });
    const res = await fetch(url, { headers });

    return {
      items: (await res.json()) as Notification[],
      unreadCount: Number(res.headers.get("X-Unread-Count")),
    };
  }
}

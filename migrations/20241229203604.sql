-- Modify "transactions" table
ALTER TABLE "public"."transactions" ADD COLUMN "user_id" uuid NOT NULL, ADD
 CONSTRAINT "fk_transactions_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

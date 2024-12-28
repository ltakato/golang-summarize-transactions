-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_categories_name" UNIQUE ("name")
);
-- Create index "idx_categories_deleted_at" to table: "categories"
CREATE INDEX "idx_categories_deleted_at" ON "public"."categories" ("deleted_at");
-- Create "category_terms" table
CREATE TABLE "public"."category_terms" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "expression" text NOT NULL,
  "category_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_categories_category_terms" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_category_terms_deleted_at" to table: "category_terms"
CREATE INDEX "idx_category_terms_deleted_at" ON "public"."category_terms" ("deleted_at");
-- Create "transactions" table
CREATE TABLE "public"."transactions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NOT NULL,
  "amount" bigint NOT NULL,
  "date" date NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_transactions_deleted_at" to table: "transactions"
CREATE INDEX "idx_transactions_deleted_at" ON "public"."transactions" ("deleted_at");
-- Create "transaction_categories" table
CREATE TABLE "public"."transaction_categories" (
  "transaction_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "category_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  PRIMARY KEY ("transaction_id", "category_id"),
  CONSTRAINT "fk_transaction_categories_category" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transaction_categories_transaction" FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create "tags" table
CREATE TABLE "public"."tags" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_tags_name" UNIQUE ("name")
);
-- Create index "idx_tags_deleted_at" to table: "tags"
CREATE INDEX "idx_tags_deleted_at" ON "public"."tags" ("deleted_at");
-- Create "tag_terms" table
CREATE TABLE "public"."tag_terms" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "expression" text NOT NULL,
  "tag_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_tags_tag_terms" FOREIGN KEY ("tag_id") REFERENCES "public"."tags" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_tag_terms_deleted_at" to table: "tag_terms"
CREATE INDEX "idx_tag_terms_deleted_at" ON "public"."tag_terms" ("deleted_at");
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
-- Create "transaction_tags" table
CREATE TABLE "public"."transaction_tags" (
  "transaction_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tag_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  PRIMARY KEY ("transaction_id", "tag_id"),
  CONSTRAINT "fk_transaction_tags_tag" FOREIGN KEY ("tag_id") REFERENCES "public"."tags" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transaction_tags_transaction" FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

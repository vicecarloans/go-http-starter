-- Create "books" table
CREATE TABLE "books" (
  "id" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NULL,
  "author" text NULL,
  "price" numeric NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_books_deleted_at" to table: "books"
CREATE INDEX "idx_books_deleted_at" ON "books" ("deleted_at");

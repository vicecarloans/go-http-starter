-- Create "authors" table
CREATE TABLE "authors" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "first_name" text NULL,
  "last_name" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_authors_deleted_at" to table: "authors"
CREATE INDEX "idx_authors_deleted_at" ON "authors" ("deleted_at");
-- Modify "books" table
ALTER TABLE "books" DROP COLUMN "author", ADD COLUMN "author_id" bigint NULL, ADD CONSTRAINT "fk_authors_books" FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

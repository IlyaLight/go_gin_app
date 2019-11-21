BEGIN;
CREATE TABLE IF NOT EXISTS "users" (
                                       "id" serial PRIMARY KEY,
                                       "username" varchar(80) UNIQUE,
                                       "password" char(60),
                                       "email" varchar(80) UNIQUE
);
CREATE TABLE IF NOT EXISTS "tags" (
                                      "id" serial PRIMARY KEY,
                                      "name" varchar(80),
                                      "created_by" varchar(80),
                                      "modified_by" varchar(80),
                                      "state" integer
);
CREATE TABLE IF NOT EXISTS "articles" (
                                          "id" serial PRIMARY KEY,
                                          "user_id" integer REFERENCES users(id),
                                          "tag_id" integer REFERENCES tags(id),
                                          "title" varchar(80),
                                          "desc" varchar(80),
                                          "content" text,
                                          "state" integer
);
COMMIT;
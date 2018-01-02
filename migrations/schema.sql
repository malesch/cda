CREATE TABLE "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"password_hash" TEXT NOT NULL,
"locale_id" integer NOT NULL,
"role_id" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "media" (
"id" TEXT PRIMARY KEY,
"name" text,
"type" TEXT NOT NULL,
"size" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "scenes" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"mediumID" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "devices" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"type" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "events" (
"id" TEXT PRIMARY KEY,
"sceneID" char(36) NOT NULL,
"deviceID" char(36) NOT NULL,
"start" integer NOT NULL,
"end" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "props" (
"id" TEXT PRIMARY KEY,
"eventID" char(36) NOT NULL,
"name" TEXT NOT NULL,
"value" TEXT NOT NULL,
"type" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);

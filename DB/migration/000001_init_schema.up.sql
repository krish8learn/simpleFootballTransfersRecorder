CREATE TABLE "footballclubs" (
  "fc_id" SERIAL PRIMARY KEY,
  "club_name" varchar NOT NULL,
  "country_fc" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "players" (
  "p_id" SERIAL PRIMARY KEY,
  "player_name" varchar UNIQUE NOT NULL,
  "position" varchar NOT NULL,
  "country_pl" varchar NOT NULL,
  "value" bigint NOT NULL,
  "footballclub_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "t_id" SERIAL PRIMARY KEY,
  "season" bigint NOT NULL,
  "player_id" int NOT NULL,
  "source_club" int NOT NULL,
  "destination_club" int NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "players" ADD FOREIGN KEY ("footballclub_id") REFERENCES "footballclubs" ("fc_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("p_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("source_club") REFERENCES "footballclubs" ("fc_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("destination_club") REFERENCES "footballclubs" ("fc_id");

CREATE INDEX ON "footballclubs" ("club_name");

CREATE INDEX ON "players" ("player_name");

CREATE INDEX ON "players" ("player_name", "footballclub_id");

CREATE INDEX ON "transfers" ("player_id");

CREATE INDEX ON "transfers" ("source_club");

CREATE INDEX ON "transfers" ("destination_club");

CREATE INDEX ON "transfers" ("source_club", "destination_club");

COMMENT ON COLUMN "footballclubs"."balance" IS 'can be positive or negative';

COMMENT ON COLUMN "players"."value" IS 'must be positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

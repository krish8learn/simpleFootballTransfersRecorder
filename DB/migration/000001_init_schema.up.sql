CREATE TABLE "footballclub" (
  "fc_id" SERIAL PRIMARY KEY,
  "club_name" varchar NOT NULL,
  "country_fc" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "player" (
  "p_id" SERIAL PRIMARY KEY,
  "player_name" varchar UNIQUE NOT NULL,
  "position" varchar NOT NULL,
  "country_pl" varchar NOT NULL,
  "value" bigint NOT NULL,
  "footballclub_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "t_id" SERIAL PRIMARY KEY,
  "season" bigint NOT NULL,
  "player_id" int NOT NULL,
  "source_club" int NOT NULL,
  "destination_club" int NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "player" ADD FOREIGN KEY ("footballclub_id") REFERENCES "footballclub" ("fc_id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("player_id") REFERENCES "player" ("p_id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("source_club") REFERENCES "footballclub" ("fc_id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("destination_club") REFERENCES "footballclub" ("fc_id");

CREATE INDEX ON "footballclub" ("club_name");

CREATE INDEX ON "player" ("player_name");

CREATE INDEX ON "player" ("player_name", "footballclub_id");

CREATE INDEX ON "transfer" ("player_id");

CREATE INDEX ON "transfer" ("source_club");

CREATE INDEX ON "transfer" ("destination_club");

CREATE INDEX ON "transfer" ("source_club", "destination_club");

COMMENT ON COLUMN "footballclub"."balance" IS 'can be positive or negative';

COMMENT ON COLUMN "player"."value" IS 'must be positive';

COMMENT ON COLUMN "transfer"."amount" IS 'must be positive';

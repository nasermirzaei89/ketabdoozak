CREATE TABLE listing_items
(
    id             VARCHAR     NOT NULL PRIMARY KEY,
    title          VARCHAR     NOT NULL,
    owner_id       VARCHAR     NOT NULL,
    owner_name     VARCHAR     NOT NULL,
    location_id    VARCHAR     NOT NULL,
    location_title VARCHAR     NOT NULL,
    types          JSONB       NOT NULL DEFAULT '[]'::jsonb,
    contact_info   JSONB       NOT NULL DEFAULT '[]'::jsonb,
    description    TEXT        NOT NULL,
    status         VARCHAR     NOT NULL,
    lent           BOOLEAN,
    thumbnail_url  VARCHAR     NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    published_at   TIMESTAMPTZ          DEFAULT now()
);

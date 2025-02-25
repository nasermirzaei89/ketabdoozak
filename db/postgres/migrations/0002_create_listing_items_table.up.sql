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

INSERT INTO listing_items (id, title, owner_id, owner_name, location_id, location_title, types, contact_info, description, status, lent, thumbnail_url, created_at, updated_at, published_at) VALUES
                          ('item-1', 'آموزش برنامه نویسی C++ به زبان ساده', 'nasermirzaei89', 'ناصر میرزائی', 'tehran-saadat-abad', 'تهران، سعادت آباد', '["donate", "exchange"]'::jsonb, '[]'::jsonb, '<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>','published',false,'https://placehold.co/300x300', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                          ('item-2', 'آموزش گام به گام برنامه نویسی Go', 'nasermirzaei89', 'ناصر میرزائی', 'tehran-saadat-abad', 'تهران، سعادت آباد', '["sell", "lend"]'::jsonb, '[]'::jsonb, '<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>','published',false,'https://placehold.co/300x300', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

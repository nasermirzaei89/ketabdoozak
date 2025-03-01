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
                          ('item-1', 'آموزش برنامه نویسی C++ به زبان ساده', 'nasermirzaei89', 'ناصر میرزائی', 'tehran-saadat-abad', 'تهران، سعادت آباد', '["donate", "exchange"]'::jsonb, '[{"type":"phoneNumber","value":"+123456789"},{"type":"sms","value":"+123456789"},{"type":"telegram","value":"user1"},{"type":"whatsapp","value":"123456789"}]'::jsonb, '<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>','published',false,'https://picsum.photos/id/367/300', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                          ('item-2', 'آموزش گام به گام برنامه نویسی Go', 'nasermirzaei89', 'ناصر میرزائی', 'tehran-saadat-abad', 'تهران، سعادت آباد', '["sell", "lend"]'::jsonb, '[{"type":"phoneNumber","value":"+989876543210"}]'::jsonb, '<p>
کتاب آموزش برنامه نویسی سی پلاس، تقریبا نو. رایگان برای هرکسی که علاقه به خوندنش داره.
<br>
لطفا برای دریافت آن تماس بگیرید.
</p>','published',false,'https://picsum.photos/id/24/300', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                          ('item-3', 'چطور دیزاینر شویم', 'nasermirzaei89', 'ناصر میرزائی', 'tehran-saadat-abad', 'تهران، سعادت آباد', '["exchange", "lend"]'::jsonb, '[{"type":"phoneNumber","value":"+989876543210"}]'::jsonb, '<p>
            اگر به دنیای طراحی علاقه‌مند هستید و به دنبال یادگیری مسیر تبدیل شدن به یک دیزاینر حرفه‌ای هستید، کتاب
            <strong>"چطور دیزاینر شویم"</strong> می‌تواند منبعی عالی برای شما باشد. این کتاب را برای مطالعه در اختیار دوستان علاقه‌مند قرار می‌دهم
            و همچنین در صورت داشتن کتابی مرتبط، امکان معاوضه نیز وجود دارد.
        </p>
        <p>
            اگر تمایل به امانت یا معاوضه دارید، لطفاً از طریق پیام با من در ارتباط باشید. پیشنهادات و نظرات شما را با کمال میل بررسی خواهم کرد.
        </p>','published',false,'https://picsum.photos/id/464/300', CURRENT_TIMESTAMP,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

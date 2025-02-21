CREATE TABLE file_manager_files
(
    filename     VARCHAR     NOT NULL PRIMARY KEY,
    size         INTEGER     NOT NULL,
    content_type VARCHAR     NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

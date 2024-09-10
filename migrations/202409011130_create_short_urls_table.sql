-- migrate:up
create table short_urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    usage_count INTEGER DEFAULT 0
);

CREATE INDEX idx_short_code ON short_urls(short_code);
CREATE INDEX idx_original_url ON short_urls(original_url);

-- migrate:down
drop table urls;

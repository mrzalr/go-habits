CREATE TABLE IF NOT EXISTS habits(
    id varchar(36) PRIMARY KEY,
    activity varchar(255) NOT NULL,
    description text,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP
)
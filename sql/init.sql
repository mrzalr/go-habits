CREATE TABLE IF NOT EXISTS habit_categories(
    id varchar(36) PRIMARY KEY,
    name varchar(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS habits(
    id varchar(36) PRIMARY KEY,
    category_id varchar(36) NOT NULL,
    activity varchar(255) NOT NULL,
    description text NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS habit_details(
    id varchar(36) PRIMARY KEY,
    habit_id varchar(36) NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    remark text NOT NULL,
    valid int(1) NOT NULL DEFAULT 0
);

ALTER TABLE habits 
ADD CONSTRAINT fk_habitscategory FOREIGN KEY (category_id) REFERENCES habit_categories(id);

ALTER TABLE habit_details
ADD CONSTRAINT fk_detailshabit FOREIGN KEY (habit_id) REFERENCES habits(id);
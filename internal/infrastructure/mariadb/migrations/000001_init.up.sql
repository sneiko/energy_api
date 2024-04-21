CREATE TABLE users
(
    id            INT AUTO_INCREMENT,
    token         VARCHAR(13) NOT NULL,
    last_activity TIMESTAMP            DEFAULT current_timestamp,
    created_at    TIMESTAMP   NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),
    UNIQUE INDEX users_token_uk (token)
);

CREATE TABLE invoices
(
    id                           INT AUTO_INCREMENT,
    invoice_number               VARCHAR(36),
    user_id                      INT,
    from_city                    VARCHAR(100),
    to_city                      VARCHAR(100),
    places                       INT,
    weight                       INT,
    volume                       FLOAT,
    sender_is_paid               BOOLEAN,
    recipient_is_paid            BOOLEAN,
    delivery_date_from           INT,
    delivery_date_from_formatted VARCHAR(48),
    delivery_date_to             INT,
    delivery_date_to_formatted   VARCHAR(48),
    sender_total_price           INT,
    recipient_total_price        INT,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE invoice_states
(
    id                    INT AUTO_INCREMENT,
    site_id               INT,
    invoice_id            INT,
    title                 VARCHAR(255),
    moving_date           INT,
    moving_date_formatted VARCHAR(48),
    moving_from_city      VARCHAR(100),
    moving_to_city        VARCHAR(100),

    PRIMARY KEY (id),
    FOREIGN KEY (invoice_id) REFERENCES invoices (id) ON DELETE CASCADE
);
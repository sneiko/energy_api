CREATE TABLE users
(
    id         INT AUTO_INCREMENT,
    phone      VARCHAR(13)  NOT NULL,
    name       VARCHAR(120) NOT NULL,
    is_active  BOOLEAN      NOT NULL DEFAULT FALSE,
    last_login DATE,
    created_at DATE         NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),

    UNIQUE INDEX users_phone_uk (phone)
);

CREATE TABLE parts
(
    id          INT AUTO_INCREMENT,
    part_number VARCHAR(60)  NOT NULL,
    type        VARCHAR(120) NOT NULL,
    name        VARCHAR(120) NOT NULL,
    description VARCHAR(120) NOT NULL,
    created_at  DATE         NOT NULL DEFAULT NOW(),
    updated_at  DATE         NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),

    INDEX parts_part_number_idx (part_number),
    INDEX parts_type_idx (type)
);

CREATE TABLE parts_images
(
    id        INT AUTO_INCREMENT,
    part_id   INT,
    image_url VARCHAR(120) NOT NULL,

    PRIMARY KEY (id),

    INDEX parts_images_part_id_idx (part_id),

    FOREIGN KEY (part_id) REFERENCES parts (id) ON DELETE CASCADE
);

CREATE TABLE cars
(
    id          INT AUTO_INCREMENT,
    brand       VARCHAR(120) NOT NULL,
    model       VARCHAR(120) NOT NULL,
    generation  INTEGER,
    image_url   VARCHAR(120),
    produced_at DATE         NOT NULL,
    produced_to DATE,

    PRIMARY KEY (id),

    INDEX cars_brand_model_idx (brand, model)
);

CREATE TABLE car_parts
(
    car_id  INT NOT NULL,
    part_id INT NOT NULL,

    UNIQUE INDEX car_parts_car_id_idx (car_id, part_id),

    FOREIGN KEY (car_id) REFERENCES cars (id) ON DELETE CASCADE,
    FOREIGN KEY (part_id) REFERENCES parts (id) ON DELETE CASCADE

);

CREATE TABLE advertisements
(
    id         INT AUTO_INCREMENT,
    user_id    INT          NOT NULL,
    part_id    INT          NOT NULL,
    prev_price REAL,
    price      REAL         NOT NULL,
    location   VARCHAR(200) NOT NULL,
    created_at DATE         NOT NULL DEFAULT NOW(),
    updated_at DATE         NOT NULL DEFAULT NOW(),

    PRIMARY KEY (id),

    INDEX advertisements_user_id_idx (user_id),
    INDEX advertisements_part_id_idx (part_id),

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (part_id) REFERENCES parts (id)
);

CREATE TABLE advertisements_images
(
    id               INT AUTO_INCREMENT,
    advertisement_id INT          NOT NULL,
    image_url        VARCHAR(120) NOT NULL,

    PRIMARY KEY (id),

    INDEX advertisements_images_advertisement_id_idx (advertisement_id),

    FOREIGN KEY (advertisement_id)
        REFERENCES advertisements (id)
        ON DELETE CASCADE
);
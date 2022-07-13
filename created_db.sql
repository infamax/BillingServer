CREATE TABLE user_balance
(
    user_id integer NOT NULL,
    balance numeric(6, 6) NOT NULL,
    PRIMARY KEY(user_id)
);

CREATE TABLE payment
(
    payment_id serial NOT NULL,
    user_id integer NOT NULL,
    date DATE NOT NULL,
    payment numeric(6, 6) NOT NULL,
    reason text,
    PRIMARY KEY(payment_id),
    FOREIGN KEY(user_id)
    REFERENCES user_balance(user_id)
);

CREATE TABLE accrual
(
    accrual_id serial NOT NULL,
    user_id integer NOT NULL,
    date DATE NOT NULL,
    accrual numeric(6, 6) NOT NULL,
    reason text,
    PRIMARY KEY(accrual_id),
    FOREIGN KEY(user_id)
    REFERENCES user_balance(user_id)
);

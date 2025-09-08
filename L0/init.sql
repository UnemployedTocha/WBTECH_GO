create table orders
(
    order_uid          varchar primary key not null,
    track_number       varchar,
    entry              varchar,
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar,
    delivery_service   varchar,
    shardkey           varchar,
    sm_id              bigint,
    date_created       timestamp,
    oof_shard          varchar
);

create table delivery
(
    id   serial primary key,
    order_uid varchar references orders (order_uid) on delete cascade,
    name      varchar,
    phone     varchar,
    zip       varchar,
    city      varchar,
    address   varchar,
    region    varchar,
    email     varchar
);

create table payment
(
    id            serial primary key,
    order_uid     varchar references orders (order_uid) on delete cascade,
    transaction   varchar,
    request_id    varchar,
    currency      varchar,
    provider      varchar,
    amount        int,
    payment_dt    bigint,
    bank          varchar,
    delivery_cost int,
    goods_total   int,
    custom_fee    int
);

create table items
(
    id           serial primary key,
    order_uid    varchar references orders (order_uid) on delete cascade,
    chrt_id      bigint,
    track_number varchar,
    price        int,
    rid          varchar,
    name         varchar,
    sale         int,
    size         varchar,
    total_price  int,
    nm_id        bigint,
    brand        varchar,
    status       int
);

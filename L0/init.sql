create table orders
(
    order_uid          text primary key,
    track_number       text unique not null,
    entry              text not null,
    locale             varchar(8) not null,
    internal_signature text,
    customer_id        text not null,
    delivery_service   text not null,
    shardkey           text not null,
    sm_id              bigint not null,
    date_created       timestamp default now() not null,
    oof_shard          text not null
);

create table delivery
(
    order_uid text primary key references orders (order_uid) on delete cascade,
    name      text not null,
    phone     varchar(20) not null,
    zip       text not null,
    city      text not null,
    address   text not null,
    region    text not null,
    email     varchar(25) not null
);

create table payment
(
    transaction   text primary key references orders (order_uid) on delete cascade,
    request_id    text,
    currency      varchar(8) not null,
    provider      text not null,
    amount        double precision not null,
    payment_dt    bigint not null,
    bank          text not null,
    delivery_cost double precision not null,
    goods_total   int not null,
    custom_fee    int not null
);

create table items
(
    rid          text default gen_random_uuid() primary key,
    order_uid    text references orders (order_uid) on delete cascade,
    chrt_id      bigint not null,
    track_number text not null,
    price        double precision not null,
    name         text not null,
    sale         double precision not null,
    size         text not null,
    total_price  double precision not null,
    nm_id        bigint not null,
    brand        text not null,
    status       int not null
);

CREATE TABLE user_groups
(
    user_id     serial          not null,
    group_name  varchar(255)    not null
);

CREATE TABLE operations
(
    user_id     serial          not null,
    group_name  varchar(255)    not null,
    operation   varchar(3)      not null,
    done_at     timestamp       not null
);
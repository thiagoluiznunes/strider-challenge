INSERT INTO
    `strider`.`posts` (
        `id`,
        `uuid`,
        `type`,
        `text`,
        `user_id`
    )
VALUES
    (
        1,
        '9d2b1a55-8aae-4f52-85b4-9e5b67ddf66f',
        'original',
        'text',
        1
    ),
    (
        2,
        'f5cdc783-acec-43be-ba6f-53608617a379',
        'original',
        'text',
        2
    ),
    (
        3,
        '3ef9d583-fa5c-45e4-8083-4a12d1594c3d',
        'original',
        'text',
        3
    );

INSERT INTO
    `strider`.`posts` (
        `uuid`,
        `type`,
        `text`,
        `user_id`,
        `post_id`
    )
VALUES
    (
        '9d2b1a55-8aae-4f52-85b4-9e5b67ddf66f',
        'repost',
        'text',
        3,
        1
    ),
    (
        'f5cdc783-acec-43be-ba6f-53608617a379',
        'quote',
        'text',
        2,
        1
    ),
    (
        '3ef9d583-fa5c-45e4-8083-4a12d1594c3d',
        'quote',
        'text',
        1,
        1
    );
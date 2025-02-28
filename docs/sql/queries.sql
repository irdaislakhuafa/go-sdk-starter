-- name: ListTodo :many
SELECT *
FROM `todo`
WHERE
 `is_deleted` = ?
 AND (
  `title` LIKE CONCAT('%', ? , '%') OR `description` LIKE CONCAT('%', ?, '%')
 )
ORDER BY (
  `title` LIKE CONCAT('%', ? , '%') OR `description` LIKE CONCAT('%', ?, '%')
) DESC
LIMIT ?
OFFSET ?;

-- name: CountTodo :one
SELECT
 COUNT(`id`) AS total
FROM
 `todo`
WHERE
 `is_deleted` = ?
 AND (
  `title` LIKE CONCAT('%', ? , '%') OR `description` LIKE CONCAT('%', ?, '%')
 );

-- name: CreateTodo :execresult
INSERT INTO `todo` (`title`, `description`, `created_at`, `created_by`)
VALUES (?, ?, ?, ?);

-- name: GetTodo :one
SELECT * FROM `todo` WHERE `id` = ?;

-- name: UpdateTodo :execresult
UPDATE `todo` SET
 `title` = ?,
 `description` = ?,
 `updated_at` = ?,
 `updated_by` = ?,
 `deleted_at` = ?,
 `deleted_by` = ?,
 `is_deleted` = ?
WHERE `id` = ?;

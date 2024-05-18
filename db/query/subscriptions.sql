-- name: CreateSubscription :one
INSERT INTO subscriptions (
    email
) VALUES (
    $1
) RETURNING *;

-- name: GetSubscriptionByEmail :one
SELECT * FROM subscriptions
WHERE email = $1;

-- name: GetSubscriptions :many
SELECT * FROM subscriptions
ORDER BY id;

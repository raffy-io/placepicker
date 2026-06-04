-- name: ListAvailableLocations :many
SELECT * FROM locations
WHERE is_dream_location = 0
ORDER BY title;

-- name: ListDreamLocations :many
SELECT * FROM locations
WHERE is_dream_location = 1
ORDER BY title;

-- name: SetDreamLocationStatus :exec
UPDATE locations
SET is_dream_location = ?
WHERE id = ?;

-- name: PollSuggestions :many
SELECT * FROM locations
WHERE is_dream_location = 0
ORDER BY RANDOM()
LIMIT 3;
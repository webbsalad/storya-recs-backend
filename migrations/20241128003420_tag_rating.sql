-- +goose Up
-- +goose StatementBegin
CREATE TABLE tag_rating (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    rating INT,
    PRIMARY KEY (user_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tag_rating;
-- +goose StatementEnd

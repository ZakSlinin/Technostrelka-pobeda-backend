CREATE TABLE subscriptions (
    subscription_id UUID NOT NULL,
    user_id UUID NOT NULL,

    name TEXT NOT NULL,
    cost NUMERIC(10,2) NOT NULL CHECK (cost >= 0),

    next_billing TIMESTAMP NOT NULL,

    status BOOLEAN,

    subscription_avatar_url TEXT,
    category TEXT,
    url_service TEXT,
    use_in_this_month BOOLEAN,
    cancellation_link TEXT,

    FOREIGN KEY (user_id)
       REFERENCES users(id)
       ON DELETE CASCADE
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);
CREATE INDEX idx_subscriptions_next_billing ON subscriptions(next_billing);
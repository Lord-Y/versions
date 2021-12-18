CREATE TABLE versions (
  versions_id SERIAL PRIMARY KEY,
  workload varchar(100) NOT NULL,
  platform varchar(100) NOT NULL,
  environment varchar(100) NOT NULL,
  version varchar(100) NOT NULL,
  changelog_url text DEFAULT NULL,
  raw text DEFAULT NULL,
  status text NOT NULL,
  date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_workload ON versions (workload);
CREATE INDEX idx_environment ON versions (environment);
CREATE INDEX idx_version ON versions (version);

-- cannot do this as lib pg don't like it in transactions
-- CREATE INDEX CONCURRENTLY idx_workload ON versions (workload);
-- CREATE INDEX CONCURRENTLY idx_environment ON versions (environment);
-- CREATE INDEX CONCURRENTLY idx_version ON versions (version);
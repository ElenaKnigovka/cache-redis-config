/**
 * Types for cache-redis-config package.
 */
export type TConfig = {
  host: string;
  port: number;
  password: string;
  db: number;
  prefix: string;
  expire: number;
  maxRetries: number;
  retryDelay: number;
};

export type TRedisConfig = {
  host: string;
  port: number;
  password: string;
  db: number;
  prefix: string;
  maxRetries: number;
  retryDelay: number;
};

export type TClientOptions = {
  retryStrategy: 'exponential' | 'linear' | 'no-retries';
  retryMaxAttempts: number;
  retryTimeout: number;
};

export type TClient = {
  config: TRedisConfig;
  options: TClientOptions;
};
// types.ts
export interface RedisConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
}

export interface CacheConfig {
  redis: RedisConfig;
  ttl?: number;
  maxItems?: number;
}

export type CacheStore = 'redis' | 'memory';

export interface CacheOptions {
  store: CacheStore;
  config: CacheConfig;
}

export interface CacheItem<T> {
  key: string;
  value: T;
  createdAt: Date;
  ttl: number;
}
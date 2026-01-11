// types.ts
export interface CacheConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
}

export interface RedisCacheConfig extends CacheConfig {
  tls?: boolean;
  username?: string;
}

export type CacheStore = 'redis' | 'memory';

export interface CacheOptions {
  store: CacheStore;
  ttl?: number;
  max?: number;
}

export interface CacheRedisOptions extends CacheOptions {
  config: RedisCacheConfig;
}
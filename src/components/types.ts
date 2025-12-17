// types.ts
export interface CacheConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
}

export interface RedisClientOptions {
  socket_keepalive: boolean;
  enable_ready_check: boolean;
}

export type CacheStore = 'redis' | 'memory';

export interface CacheOptions {
  store: CacheStore;
  ttl: number;
  max: number;
}

export interface CacheRedisOptions extends CacheOptions {
  config: CacheConfig;
  clientOptions?: RedisClientOptions;
}

export interface CacheMemoryOptions extends CacheOptions {
  // no additional options for memory store
}
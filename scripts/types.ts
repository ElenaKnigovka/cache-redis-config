// types.ts
export interface CacheConfig {
  host: string;
  port: number;
  password?: string;
  db?: number;
}

export interface RedisConfig extends CacheConfig {
  tls?: boolean;
  username?: string;
}

export interface CacheOptions {
  cacheConfig: CacheConfig | RedisConfig;
  ttl: number;
  maxMemory: number;
}

export type CacheType = 'memory' | 'redis'; 

export interface CacheRedisConfig {
  type: CacheType;
  options: CacheOptions;
}
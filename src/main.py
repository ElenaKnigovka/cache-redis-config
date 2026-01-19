import os
import redis
from dotenv import load_dotenv

load_dotenv()

class RedisConfig:
    def __init__(self):
        self.host = os.getenv('REDIS_HOST')
        self.port = int(os.getenv('REDIS_PORT'))
        self.db = int(os.getenv('REDIS_DB'))
        self.password = os.getenv('REDIS_PASSWORD')

    def connect(self):
        try:
            self.redis_client = redis.Redis(host=self.host, port=self.port, db=self.db, password=self.password)
            return self.redis_client
        except redis.exceptions.ConnectionError as e:
            print(f"Connection to Redis failed: {e}")
            return None

def main():
    config = RedisConfig()
    redis_client = config.connect()
    if redis_client:
        print("Connected to Redis successfully")
        # Test the connection by setting and getting a value
        redis_client.set('test_key', 'test_value')
        print(redis_client.get('test_key'))

if __name__ == "__main__":
    main()
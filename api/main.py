# main.py
import json
import redis
from logging.config import dictConfig
from logging.handlers import RotatingFileHandler
from cachetools import TTLCache

# Set up logging configuration
dictConfig({
    'version': 1,
    'formatters': {
        'default': {
            'format': '[%(asctime)s] %(levelname)s in %(module)s: %(message)s'
        }
    },
    'handlers': {
        'wsgi': {
            'class': 'logging.StreamHandler',
            'stream': 'ext://flask.logging.wsgi_errors_stream',
            'level': 'ERROR'
        },
        'file': {
            'class': 'logging.handlers.RotatingFileHandler',
            'filename': 'cache_redis.log',
            'maxBytes': 1000000,
            'backupCount': 3,
            'level': 'INFO'
        }
    },
    'root': {
        'level': 'INFO',
        'handlers': ['file']
    }
})

# Set up Redis connection
redis_client = redis.Redis(host='localhost', port=6379)

# Set up cache
cache = TTLCache(maxsize=1000, ttl=60)

# Function to store data in Redis
def store_in_redis(key, value, expires):
    redis_client.setex(key, value, expires)

# Function to retrieve data from Redis
def get_from_redis(key):
    return redis_client.get(key)

# Function to remove data from Redis
def remove_from_redis(key):
    redis_client.delete(key)
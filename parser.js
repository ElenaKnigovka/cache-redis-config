const redis = require('redis');
const fs = require('fs');
const path = require('path');
const config = require('./config');

class Parser {
  constructor() {
    this.client = redis.createClient({
      host: config.host,
      port: config.port,
      password: config.password,
    });
  }

  parseFile(filePath) {
    try {
      const data = fs.readFileSync(filePath, 'utf8');
      const lines = data.split('\n');
      const parsedData = [];

      lines.forEach((line) => {
        if (line.trim() !== '') {
          const match = line.match(/^(\w+):\s*(.*)$/);
          if (match) {
            parsedData.push({
              key: match[1],
              value: match[2],
            });
          }
        }
      });

      return parsedData;
    } catch (error) {
      if (error.code === 'ENOENT') {
        console.error('Error: The file', filePath, 'does not exist.');
      } else if (error.code === 'EACCES') {
        console.error('Error: You do not have permission to read the file', filePath);
      } else {
        console.error('Error parsing file:', error);
      }
      return [];
    }
  }

  saveFile(filePath, data) {
    try {
      fs.writeFileSync(filePath, data.join('\n'));
      console.log('File saved successfully to', filePath);
    } catch (error) {
      console.error('Error saving file:', error);
    }
  }
}

module.exports = Parser;
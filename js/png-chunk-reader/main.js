const fs = require('fs').promises;
const { getChunkData } = require('./chunkReader');

async function main() {
  const filenames = [
    "/app/test/logo.png"
  ];
  const keyword = "workflow"; // 可以根据需要更改关键字

  for (const filename of filenames) {
    try {
      const buffer = await fs.readFile(filename);
      const chunkData = await getChunkData(buffer, keyword);

      if (chunkData) {
        console.log(`Data for keyword '${keyword}' in ${filename}:`, chunkData);
      } else {
        console.log(`No data found for keyword '${keyword}' in ${filename}.`);
      }
    } catch (error) {
      console.error(`Error processing file ${filename}:`, error);
    }
  }
}

main();

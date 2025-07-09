const zlib = require('zlib');
const { promisify } = require('util');

const inflateAsync = promisify(zlib.inflate);

async function getChunkData(buffer, keyword) {
  let offset = 8; // Skip PNG signature

  while (offset < buffer.length) {
    const length = buffer.readUInt32BE(offset);
    const type = buffer.toString('ascii', offset + 4, offset + 8);
    const data = buffer.slice(offset + 8, offset + 8 + length);

    if (['tEXt', 'iTXt', 'zTXt'].includes(type)) {
      const nullSeparatorIndex = data.indexOf(0);
      const chunkKeyword = data.slice(0, nullSeparatorIndex).toString('utf8');

      if (chunkKeyword === keyword) {
        let text;
        switch (type) {
          case 'tEXt':
            text = data.slice(nullSeparatorIndex + 1);
            break;
          case 'zTXt':
            {
              const compressionMethod = data[nullSeparatorIndex + 1];
              if (compressionMethod !== 0) {
                throw new Error(`Unsupported compression method for zTXt: ${compressionMethod}`);
              }
              const compressedText = data.slice(nullSeparatorIndex + 2);
              text = await inflateAsync(compressedText);
            }
            break;
          case 'iTXt':
            let currentIndex = nullSeparatorIndex + 1;
            const compressionFlag = data[currentIndex++];
            const compressionMethod = data[currentIndex++];

            // Skip language tag
            currentIndex = data.indexOf(0, currentIndex) + 1;
            // Skip translated keyword
            currentIndex = data.indexOf(0, currentIndex) + 1;

            const rawText = data.slice(currentIndex);

            if (compressionFlag === 1) {
              if (compressionMethod !== 0) {
                throw new Error(`Unsupported compression method for iTXt: ${compressionMethod}`);
              }
              text = await inflateAsync(rawText);
            } else {
              text = rawText;
            }
            break;
        }
        return text.toString('utf8');
      }
    }

    offset += 12 + length; // Move to next chunk
  }

  return null;
}

module.exports = { getChunkData };
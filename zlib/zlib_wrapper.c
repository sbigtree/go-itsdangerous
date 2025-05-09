#include <zlib.h>
#include <stdlib.h>

unsigned char* zlib_deflate(const char* input, int inputLen, int* outLen) {
    uLongf compressedLen = compressBound(inputLen);
    unsigned char* compressedData = (unsigned char*)malloc(compressedLen);

    int res = compress(compressedData, &compressedLen, (const unsigned char*)input, inputLen);
    if (res != Z_OK) {
        return NULL;  // 压缩失败
    }

    *outLen = (int)compressedLen;
    return compressedData;
}

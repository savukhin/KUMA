#include <FFat.h>
// #include <SPIFFS.h>

// #define FILESYSTEM SPIFFS
#define FILESYSTEM FFat

void writeFile(const char* filename, String data) {
    File file = FILESYSTEM.open(filename, "w", true);
    if (!file) {
        throw new std::runtime_error("Cannot open write file " + std::string(filename));
    }

    for (size_t i = 0; i < data.length(); i++) {
        file.write(data[i]);
    }
}
void writeFile(const char* filename, std::string data) {
    File file = FILESYSTEM.open(filename, "w", true);
    if (!file) {
        throw new std::runtime_error("Cannot open write file " + std::string(filename));
    }

    for (size_t i = 0; i < data.length(); i++) {
        file.write(data[i]);
    }
}

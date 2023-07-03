#pragma once

#include <FFat.h>
// #include <SPIFFS.h>

// #define FILESYSTEM SPIFFS
#define FILESYSTEM FFat

String readFile(File &file) {
    String result;

    while (file.available()){
        result += file.read();
    }

    return result;
}

String readFile(const char* filename) {
    File file = FILESYSTEM.open(filename);
    if (!file) {
        throw new std::runtime_error("Cannot open read file " + std::string(filename));
    }

    return readFile(file);
}

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

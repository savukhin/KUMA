#pragma once

#include "utils.hpp"

const char* titleFileName = "title.txt";
const char* usernameFileName = "username.txt";
const char* passwordFileName = "password.txt";

void setTitle(String newTitle) {
    return writeFile(titleFileName, newTitle);
}
void setUsername(String newUsername) {
    return writeFile(usernameFileName, newUsername);
}
void setPassword(String newPassword) {
    return writeFile(passwordFileName, newPassword);
}
void setTitle(std::string newTitle) {
    return writeFile(titleFileName, newTitle);
}
void setUsername(std::string newUsername) {
    return writeFile(usernameFileName, newUsername);
}
void setPassword(std::string newPassword) {
    return writeFile(passwordFileName, newPassword);
}

String getTitle() {
    return readFile(titleFileName);
}
String getUsername() {
    return readFile(usernameFileName);
}
String getPassword() {
    return readFile(passwordFileName);
}

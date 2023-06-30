#include "utils.hpp"

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

std::string getTitle() {
    return readFile(titleFileName);
}
std::string getUsername() {
    return readFile(usernameFileName);
}
std::string getPassword() {
    return readFile(passwordFileName);
}

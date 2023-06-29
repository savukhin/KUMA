#include <WiFi.h>
#include <WiFiClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>
#include <Update.h>
#include <string.h>
#include <FFat.h>
// #include <SPIFFS.h>

// #define FILESYSTEM SPIFFS
#define FILESYSTEM FFat

const char* host = "esp32";
const char* ssid = "xxx";
const char* password = "xxxx";

const char* titleArg = "title";
const char* usernameArg = "username";
const char* passwordArg = "password";

const char* titleFileName = "title.txt";
const char* usernameFileName = "username.txt";
const char* passwordFileName = "password.txt";

WebServer server(80);

File *titleFile;
File *usernameFile;
File *passwordFile;

const char* indexPage;
const char* indexPageFallback = "Error loading page: check if SPIFFS works";

std::string readFile(File &file) {
    std::string result;

    while (file.available()){
        result.push_back(file.read());
    }

    return result;
}

std::string readFile(const char* filename) {
    File file = FILESYSTEM.open(filename);
    if (!file) {
        throw new std::runtime_error("Cannot open read file " + std::string(filename));
    }

    return readFile(file);
}

void setupIndexPage() {
    File indexPageFile = FILESYSTEM.open("/text.txt");

    if(!indexPageFile) {
        indexPage = indexPageFallback;
    } else {
        indexPage = readFile(indexPageFile).c_str();
    }

    indexPageFile.close();
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

/*
 * setup function
*/
void setup(void) {
    setupIndexPage();

    Serial.begin(115200);

    // Connect to WiFi network
    WiFi.begin(ssid, password);
    Serial.println("");

    // Wait for connection
    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");
    }
    Serial.println("");
    Serial.print("Connected to ");
    Serial.println(ssid);
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());

    /*use mdns for host name resolution*/
    if (!MDNS.begin(host)) { //http://esp32.local
        Serial.println("Error setting up MDNS responder!");
        while (1) {
        delay(1000);
        }
    }
    Serial.println("mDNS responder started");
    /*return index page which is stored in serverIndex */
    server.on("/", HTTP_GET, []() {
        server.sendHeader("Connection", "close");
        server.send(200, "text/html", indexPage);
    });

    /*handling update data: title, username and password */
    server.on("/update-data", HTTP_POST, []() {
        bool isCorrectQuery = (server.hasArg(titleArg) && server.hasArg(usernameArg) && server.hasArg(passwordArg));
        if (!isCorrectQuery) {
            server.send(400, "text/plain", "Query is incorrect");
            return;
        }

        setTitle(server.arg(titleArg));
        setUsername(server.arg(usernameArg));
        setPassword(server.arg(passwordArg));
    });

    server.on("/get-data", HTTP_GET, []() {
        std::string result = "{"
"title: " + getTitle() + ",";
"username: " + getUsername() + ",";
"password: " + getPassword() + ",";
"}";

        server.sendHeader("Content-Type", "applcation/json");
        server.send(200, "application/json", result.c_str());
    });

    /*handling uploading firmware file */
    server.on("/update-firmware", HTTP_POST, []() {
        server.sendHeader("Connection", "close");
        server.send(200, "text/plain", (Update.hasError()) ? "FAIL" : "OK");
        ESP.restart();
    }, []() {
        HTTPUpload& upload = server.upload();
        if (upload.status == UPLOAD_FILE_START) {
        Serial.printf("Update: %s\n", upload.filename.c_str());
        if (!Update.begin(UPDATE_SIZE_UNKNOWN)) { //start with max available size
            Update.printError(Serial);
        }
        } else if (upload.status == UPLOAD_FILE_WRITE) {
        /* flashing firmware to ESP*/
        if (Update.write(upload.buf, upload.currentSize) != upload.currentSize) {
            Update.printError(Serial);
        }
        } else if (upload.status == UPLOAD_FILE_END) {
        if (Update.end(true)) { //true to set the size to the current progress
            Serial.printf("Update Success: %u\nRebooting...\n", upload.totalSize);
        } else {
            Update.printError(Serial);
        }
        }
    });
    server.begin();
}

void loop(void) {
    server.handleClient();
    delay(1);
}
#include <WiFi.h>
#include <WiFiClient.h>
#include <HTTPClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>
#include <Update.h>
#include <string.h>

#include "cnc.hpp"
#include "utils.hpp"
#include "credentials.hpp"
#include "requests.hpp"

const char* host = "esp32";
const char* ssid = "xxx";
const char* password = "xxxx";

const char* titleArg = "title";
const char* usernameArg = "username";
const char* passwordArg = "password";

const uint8_t PIN_STOPPED = 1;
const uint8_t PIN_WORKING = 2;
const uint8_t PIN_BROKEN = 3;
Cnc cnc;

WebServer server(80);
Tokens tokens;

File *titleFile;
File *usernameFile;
File *passwordFile;

const char* indexPage;
const char* indexPageFallback = "Error loading page: check if SPIFFS works";

String serverName = "kuma.slyfoxprod.ru";

void setupIndexPage() {
    File indexPageFile = FILESYSTEM.open("/text.txt");

    if(!indexPageFile) {
        indexPage = indexPageFallback;
    } else {
        indexPage = readFile(indexPageFile).c_str();
    }

    indexPageFile.close();
}


/*
 * setup function
*/
void setup(void) {
    tokens = getTokens(serverName);

    cnc = Cnc(PIN_STOPPED, PIN_WORKING, PIN_BROKEN);
    void (*onChange)(CncStatus&) = [](CncStatus &new_status) {
    // std::function<void (CncStatus&)> onChange = [&](CncStatus &new_status) {
        tokens = sendStatus(serverName, tokens, new_status);
    };

    cnc.addOnChange(onChange);

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
    server.on("/update-credentials", HTTP_POST, []() {
        bool isCorrectQuery = (server.hasArg(titleArg) && server.hasArg(usernameArg) && server.hasArg(passwordArg));
        if (!isCorrectQuery) {
            server.send(400, "text/plain", "Query is incorrect");
            return;
        }

        setTitle(server.arg(titleArg));
        setUsername(server.arg(usernameArg));
        setPassword(server.arg(passwordArg));
    });

    server.on("/get-credentials", HTTP_GET, []() {
        String result = "{"
"title: " + getTitle() + ","
"username: " + getUsername() + ","
"password: " + getPassword() +
"}";

        server.sendHeader("Content-Type", "applcation/json");
        server.send(200, "application/json", result.c_str());
    });

    server.on("/get-status", HTTP_GET, []() {
        std::string result = "{"
"title: " + statusToInt(cnc.getStatus());
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
    cnc.loop();
    delay(1);
}
#pragma once
#include <WiFi.h>
#include <WiFiClient.h>
#include <HTTPClient.h>
#include <ArduinoJson.h>
#include <exception>

#include "credentials.hpp"
#include "cnc.hpp"

struct Tokens {
  String access_token;
  String refresh_token;
};

Tokens getTokens(String serverName) {
    if (WiFi.status() != WL_CONNECTED) {
      Serial.println("Wifi not connected");
      return Tokens{"-1", "-1"};
    }

    String authPath = serverName + "/api/v1/auth/login";

    HTTPClient http;
    http.begin(authPath.c_str());
    String payload = "{"
"username : " + getUsername() + ","
"password : " + getPassword() +
"}";


    // const char* method = "GET";
    int httpResponseCode = http.sendRequest("GET", payload);

    if (httpResponseCode != 200) {
      http.end();
      Serial.print("Error code: ");
      Serial.println(httpResponseCode);
      throw new std::runtime_error("error auth");
    }
    String response = http.getString();
    http.end();

    StaticJsonDocument<200> doc;
    DeserializationError error = deserializeJson(doc, response);

    // Test if parsing succeeds.
    if (error) {
      Serial.print(F("Deserialization of json failed: "));
      Serial.println(error.f_str());
      throw new std::runtime_error("cannot deserialize");
    }

    String access_token = doc["access-token"];
    String refresh_token = doc["refresh-token"] ;

    return Tokens{access_token, refresh_token};
}


Tokens sendStatus(String serverName, Tokens tokens, CncStatus status) {
    if (WiFi.status() != WL_CONNECTED) {
        Serial.println("Wifi not connected");
        return tokens;
    }

    Tokens new_tokens = tokens;
    HTTPClient http;

    String serverPath = serverName + "/api/v1/update-status/cnc";
    
    // Your Domain name with URL path or IP address with path
    http.begin(serverPath.c_str());

    // const char* headerNames[] = { "access-token", "refresh-token" };
    // http.collectHeaders(headerNames, sizeof(headerNames)/sizeof(headerNames[0]));
    
    http.addHeader("refresh-token", tokens.refresh_token);
    http.addHeader("Authorization", tokens.access_token);

    String payload = "{"
"status : " + String(statusToInt(status)) + ","
"}";
    
    // Send HTTP GET request
    int httpResponseCode = http.POST(payload);
    
    if (httpResponseCode != 200) {
      http.end();
      Serial.print("Error code: ");
      Serial.println(httpResponseCode);
      throw new std::runtime_error("error send request");
    }

    if (http.hasHeader("access-token") && http.hasHeader("refresh-token")) {
      new_tokens.access_token = http.header("access-token");
      new_tokens.refresh_token = http.header("refresh-token");
    }

    http.end();

    return new_tokens;
}

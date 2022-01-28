#include <WiFiUdp.h>
#include <NTPClient.h>
#include <ESP8266WiFi.h>
#include <Servo.h>
#include <SPI.h>
#include <MFRC522.h>
#include <ArduinoJson.h>
#include <PubSubClient.h>
#include "Hash.h"
#include "constants.h"


WiFiUDP ntpUDP;
NTPClient timeClient(ntpUDP, "pool.ntp.org");

struct _time{
  int hour;
  int minute;  
};
_time lightsOnEntranceTimes[2];

void addOnAndOffLightTimes(_time onTime, _time offTime){
    lightsOnEntranceTimes[0] = onTime;
    lightsOnEntranceTimes[1] = offTime;
}
_time getTime(){
  timeClient.update();
  int currentHour = timeClient.getHours();
  int currentMinute = timeClient.getMinutes();

  _time t; 
  t.hour = currentHour;
  t.minute = currentMinute;

  return t;
}
bool shouldTurnOnLight(_time t){

  if(t.hour > lightsOnEntranceTimes[0].hour || (t.hour == lightsOnEntranceTimes[0].hour && t.minute >= lightsOnEntranceTimes[0].minute)){
    if(t.hour < lightsOnEntranceTimes[1].hour || (t.hour == lightsOnEntranceTimes[1].hour && t.minute <= lightsOnEntranceTimes[1].minute)){
       Serial.print(t.hour);
       Serial.print(":");
       Serial.print(t.minute);

       Serial.print(" <= ");
       Serial.print(lightsOnEntranceTimes[0].hour);
       Serial.print(":");
       Serial.print(lightsOnEntranceTimes[0].minute);

       Serial.print(" <= ");
       Serial.print(lightsOnEntranceTimes[1].hour);
       Serial.print(":");
       Serial.print(lightsOnEntranceTimes[1].minute);
       return true;
    }
  }
  return false;
}
_time parseTime(String time){
    int i0 = time.indexOf(":");
    int hour = time.substring(0, i0).toInt();
    int i1 = time.lastIndexOf(":");
    int min = time.substring(i0, i1).toInt();

    return {hour, min};
}
void scheduleOfficeLights(String schedule){
    StaticJsonDocument<200> doc;
    DeserializationError error = deserializeJson(doc, schedule.c_str());

    // Test if parsing succeeds.
    if (error) {
        Serial.print(F("deserializeJson() failed: "));
        Serial.println(error.f_str());
        return;
    }
    String lightOffTime = (const char*)doc["light_off_time"];
    String lightOnTime= (const char*)doc["light_on_time"];

    _time onTime = parseTime(lightOnTime);
    _time offTime = parseTime(lightOffTime);
    addOnAndOffLightTimes(onTime, offTime);    
}

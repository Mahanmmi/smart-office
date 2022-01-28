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
       Serial.print("now: ");
       Serial.print(t.hour);
       Serial.print(":");
       Serial.println(t.minute);

       Serial.print("Start: ");
       Serial.print(lightsOnEntranceTimes[0].hour);
       Serial.print(":");
       Serial.println(lightsOnEntranceTimes[0].minute);

       Serial.print("End: ");
       Serial.print(lightsOnEntranceTimes[1].hour);
       Serial.print(":");
       Serial.println(lightsOnEntranceTimes[1].minute);
       return true;
    }
  }
  return false;
}
#include "timehandler.h"
_time parseTime(String time){
    int i0 = time.indexOf(":");
    int hour = time.substring(0, i0).toInt();
    int i1 = time.lastIndexOf(":");
    int min = time.substring(i0, i1).toInt();

    Serial.println(hour);
    Serial.println(min);
    
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
    Serial.println((const char*)doc["light_off_time"]);
    String lightOffTime = (const char*)doc["light_off_time"];
    String lightOnTime= (const char*)doc["light_on_time"];

    _time onTime = parseTime(lightOnTime);
    _time offTime = parseTime(lightOffTime);
    addOnAndOffLightTimes(onTime, offTime);    
}
void mqttMessageHandler(char *topic, byte *payload, unsigned int length) {
	Serial.print("Message arrived in topic: ");
	Serial.println(topic);
	String message = "";
	for (int i = 0; i < length; i++)
		message += (char)payload[i];
	
	Serial.print("Message: "+message);
	Serial.println("-----------------------");

	if(String(topic) == String(LIGHT_SCHEDULE_TOPIC)){
        scheduleOfficeLights(message);
	}else{

    }
}

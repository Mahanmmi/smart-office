// Wifi
#ifndef APSSID
#define APSSID "wood1"
#define APPSK  "f!rm!n0w00d"
#endif

// MQTT
const char *MQTT_BROKER = "192.168.1.106";
// const char * = "test1";
const char *MQTT_USERNAME = "emqx";
const char *MQTT_PASSWORD = "public";
const int MQTT_PORT = 9100;

// Ultrusonic
#define echoPin D2
#define trigPin D1
const int OBJECT_DETECTION_DISTANCE = 5;

//Servo
#define SERVO          D8
const int DOOR_CLOSED_DEGREE = 0;
const int DOOR_OPENED_DEGREE = 90;

//RFID
#define RST_PIN         D3          
#define SS_PIN          D4         

// Office LED
#define OFFICE_LIGHT D0
#define ON LOW
#define OFF HIGH

// free RX, TX
#define RX 3
#define TX 1

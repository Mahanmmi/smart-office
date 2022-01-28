#include "mqtthandlers.h"


const char *ssid = APSSID;
const char *password = APPSK;

double duration; // variable for the duration of sound wave travel
double distance; // variable for the distance measurement

const String mqtt_client_id = "ESP8266-" + sha1(WiFi.macAddress());

WiFiClient espClient;
PubSubClient client(espClient);

Servo myservo;  

MFRC522 mfrc522(SS_PIN, RST_PIN);  // Create MFRC522 instance

struct KeyValue{
	String key;
	bool value;
};
KeyValue cardInRoom[2] = {{CARD_ID_1, false},{CARD_ID_2, false}};


double calculateDistance(){
	// Clears the trigPin condition
	digitalWrite(trigPin, LOW);
	delayMicroseconds(2);
	// Sets the trigPin HIGH (ACTIVE) for 10 microseconds
	digitalWrite(trigPin, HIGH);
	delayMicroseconds(10);
	digitalWrite(trigPin, LOW);
	// Reads the echoPin, returns the sound wave travel time in microseconds
	duration = pulseIn(echoPin, HIGH);
	distance = duration * 0.034 / 2; // Speed of sound wave divided by 2 (go and back)
	return distance;
}
void turnOfficeLight(int state){//1:on 0:off
	if(state==1)
		digitalWrite(OFFICE_LIGHT, ON);
	else
		digitalWrite(OFFICE_LIGHT, OFF);
}
void checkOfficeDoor(){
	double distance = calculateDistance();
	if (distance <= OBJECT_DETECTION_DISTANCE){
		Serial.println("Object detected");
		_time currentTime = getTime();
		if(shouldTurnOnLight(currentTime)){
			turnOfficeLight(1);
		}else{
			turnOfficeLight(0);
		}
	}
}
void moveDoor(int degree){
	myservo.write(degree);
}
void changeCardState(String content){
	for (int i=0; i<sizeof(cardInRoom); i++){
		if(cardInRoom[i].key == content){
			if(!cardInRoom[i].value){
				client.publish("checkin", content.c_str());	
			}else{
				client.publish("checkout", content.c_str());
			}
			cardInRoom[i].value = !cardInRoom[i].value;
			break;
		}
	}
}
String checkRFID(){
	static int lastCardCheck; 
	if(millis() - lastCardCheck < REST_AFTER_CHECK)
		return "";
	if(mfrc522.PICC_IsNewCardPresent()) {
		if (mfrc522.PICC_ReadCardSerial()) {
			String content= "";
			byte letter;
			for (byte i = 0; i < mfrc522.uid.size; i++) 
			{
				content.concat(String(mfrc522.uid.uidByte[i] < 0x10 ? " 0" : " "));
				content.concat(String(mfrc522.uid.uidByte[i], HEX));
			}
			content.trim();
			Serial.println(content);
			
			changeCardState(content);
			lastCardCheck = millis();
			return content;
		}
	}
	return "";
}
double readLDR(){
	int adc_value = analogRead(A0);
	return (double)adc_value/1024*100;
}
void connectWifi(){
	WiFi.begin(ssid, password);
	while (WiFi.status() != WL_CONNECTED) {
	delay(1000);
	Serial.println("Connecting to WiFi..");
	}

	Serial.println(WiFi.localIP());
}
void connectMQTT(){
	client.setServer(MQTT_BROKER, MQTT_PORT);
	client.setCallback(mqttMessageHandler);
	while (!client.connected()) {
		Serial.printf("The client %s connects to the public mqtt broker\n", mqtt_client_id);
		// if (client.connect(mqtt_client_id.c_str(), mqtt_username, mqtt_password)) {
		if (client.connect(mqtt_client_id.c_str())) {
			Serial.println("connected to mqtt broker");
		} else {
			Serial.print("failed with state ");
			Serial.print(client.state());
			delay(2000);
		}
	}
}
void setup() {
	//Serial
	Serial.begin(9600);

	//RFID
	SPI.begin();			
	mfrc522.PCD_Init();		
	delay(4);				// Optional delay. Some board do need more time after init to be ready, see Readme
	mfrc522.PCD_DumpVersionToSerial();	// Show details of PCD - MFRC522 Card Reader details

	//Servo
	myservo.attach(SERVO);  

	//Ultrasonic
	pinMode(trigPin, OUTPUT); 
	pinMode(echoPin, INPUT); 

	//Office LED
	pinMode(OFFICE_LIGHT, OUTPUT);
	digitalWrite(OFFICE_LIGHT, OFF);

	//LDR
	// pinMode(RX, FUNCTION_3);
	// pinMode(TX, FUNCTION_3);

	connectWifi();
	connectMQTT();

	client.subscribe(LIGHT_SCHEDULE_TOPIC);
	client.publish(OFFICE_CONNECT_TOPIC, mqtt_client_id.c_str());

	timeClient.begin();
	timeClient.setTimeOffset(12600);

}
void loop() {
	//OFFICE FOOR
	checkOfficeDoor();
	
	//RFID ROOM
	checkRFID();
	
	//MQTT
	if(!client.connected()){
		connectMQTT();
	}
	client.loop();
}
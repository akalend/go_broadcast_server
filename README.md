# CLIENT and broadcast SERVER

This project is a messenger prototype.

## Installation
Start the script `build.sh`

    bash build.sh

## RUN
Start on the console the `server` and start some clients programm `client`. After the start, each client is assigned a number.   

	 ./client 
	Connect from server: client #0
	>


## Messaging

Send the message all clients:

		>tis message send all clients

Sending the message to a specific client, number 3 (#3) for example

	 ./client 
	Connect from server: client #0
	>#3 Message for client number 3
	>[0 30 35 48 32 77 101 115 115 97 103 101 32 102 111 114 32 99 108 105 101 110 116 32 110 117 109 98 101 114 32 51]
	Send bytes: 32

## Binary protocol

	 ./client 
	Connect from server: client #1
	ab
	>[0 2 97 98]
	    | text of message: [ab]
	    |
	the first two bytes is length of message

## Details
The max length of message is 252 bytes. The max connections is 256. The legth package of message is length of text and one byte. 




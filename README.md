# Botzilla
Botzilla serves as the core protocol facilitating communication between the various components of my robot. It acts as the hub, overseeing all the messages exchanged within the system. While this approach introduces a small amount of delay in the messaging process, it's easier to implement.

Command Types

For now, I've identified two main types of communication:

- Continuous Data Flow:
This includes scenarios where data needs to be sent and received continuously, such as with cameras, sensors, or any variable that requires a constant stream of information.

- Request-Response (End-to-End or Broadcast):
This is the classic request-response model, where a request is sent and a corresponding response is received, commonly used for communication between a sender and a receiver.

# Data Structure

## Continuous Data
Currently, continuous data is transmitted using the WebSocket protocol. Every component that generates continuous data will automatically begin sending it to Bozilla upon establishing an initial connection. Other components can access this data in real time by opening their own WebSocket connection to Bozilla, or they can retrieve the latest data snapshot via a TCP connection (though this snapshot may have some delay).

Examples: Sensors, Camera Modules

## Discrete Data
All discrete data is sent via TCP connections. Each TCP connection includes a header specifying the intended target(s) for the data, which can be either a single component or an array of components. Responses include the IDs of the devices that acknowledged the data along with their respective response codes.

# Command Handler
Botzilla itself is a component that can be connected to. The ID of Botzilla is "0000" and it doesn't belong to any group. Here are the list of commands:
- "0000" : Connect
- "0001" : Get connected components
- "0002" : Create A new Group
- "0003" | "Group Id" : Assign the component to Group
- "0004" | "Group Id" : Remove the component from Group

# Registery
Upon activation, each component will send a packet to Botzilla to register itself. From then on, Botzilla will periodically send packets to verify if the components are still active.

# Serialization
Currently, all data is transmitted in JSON format, but this will be changed in the future to use byte codes for improved efficiency.

# Logging / Debugging
All the data is sent to a simple application for debugging purposes, which can be found here.
<a href="https://github.com/rima1881/Eye-of-Madness">Eye of Madness</a>
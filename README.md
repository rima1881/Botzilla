# Botzilla
Botzilla serves as the core protocol facilitating communication between the various components of my robot. It acts as the hub, overseeing all the messages exchanged within the system. While this approach introduces a small amount of delay in the messaging process, it's easier to implement.

Here’s a rephrased version:
Command Types

For now, I've identified three main types of communication:

- Continuous Data Flow:
This includes scenarios where data needs to be sent and received continuously, such as with cameras, sensors, or any variable that requires a constant stream of information.

- Request-Response (End-to-End):
This is the classic request-response model, where a request is sent and a corresponding response is received, commonly used for communication between a sender and a receiver.

- Broadcast Commands:
This type is used when a command needs to be sent to multiple recipients at once, ensuring that all relevant components receive the message.

# Structure



# Registery
Upon activation, each component will send a packet to Botzilla to register itself. From then on, Botzilla will periodically send packets to verify if the components are still active.

# Logging / Debugging
All the data is sent to a simple application for debugging purposes, which can be found here.
<a href="https://github.com/rima1881/Eye-of-Madness">Eye of Madness</a>
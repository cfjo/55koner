a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
- Packages in our implementation are ints. The datastructure we use to transmit data and meta-data is channels, which is a datastructure that works like a circular queue.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
- Our implementation uses threads to simulate how the protocol would work. It is not realistic to actually use threads since the protocol should run across a network.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
- The sender will use sequence numbers to indicate the order of the packages being send. Then it's the receivers responsibility to re-order the packages by the sequence numbers that came with the packages. 

d) In case messages can be delayed or lost, how does your implementation handle message loss?
- If the numbers don't add up, the connection will fail. If the messages are delayed it won't have an effect on our code, as both parties will wait for a message from the other, before continuing.

e) Why is the 3-way handshake important?
- It is important because it establishes a reliable connection between a client and a server over a network.

#ifndef UDP_CLIENT_H
#define UDP_CLIENT_H

class UDPResponse {
public:
  UDPResponse(int error, String body);

  int error;
  String body;
};

class UDPClient {
public:
  UDPClient(int localPort, IPAddress remoteIP, int remotePort);
  void begin();

  UDPResponse request(String route, String body);
  UDPResponse requestWait(String route, String body);

private:
  UDP Udp;

  int localPort;
  IPAddress remoteIP;
  int remotePort;

  int timeoutPeriod;
};

#endif

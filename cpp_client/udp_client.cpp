#include "application.h"
#include "udp_client.h"

UDPResponse::UDPResponse(int err, String bd) {
  error = err;
  body = bd;
}

UDPClient::UDPClient(int lp, IPAddress rip, int rp) {
  localPort = lp;
  remoteIP = rip;
  remotePort = rp;

  timeoutPeriod = 1000;
}

void UDPClient::begin() {
  Udp.begin(localPort);
}

UDPResponse UDPClient::request(String route, String body) {
  // Prepare request body
  String requestBuffer = route + "|" + body;

  // Send UDP packet
  if (Udp.sendPacket(requestBuffer, requestBuffer.length(), remoteIP, remotePort) < 0) {
    // Failed
    return UDPResponse(1, "");
  }

  // Return empty response
  return UDPResponse(0, "");
}

UDPResponse UDPClient::requestWait(String route, String body) {
  // Prepare request body
  String requestBuffer = route + "|" + body;

  // Send UDP packet
  if (Udp.sendPacket(requestBuffer, requestBuffer.length(), remoteIP, remotePort) < 0) {
    // Failed
    return UDPResponse(1, "");
  }

  // Wait for incoming UDP packet
  int resBufSize;
  long startTime = millis();
  while (true) {
    // timeout after 1s
    if (millis() - startTime > timeoutPeriod) {
      return UDPResponse(2, "");
    }
    resBufSize = Udp.parsePacket();
    if (resBufSize != 0) {
      break;
    }
  }

  // Try reading incoming UDP packet
  char resBuf[resBufSize];
  Udp.read(resBuf, resBufSize);
  return UDPResponse(0, resBuf);
}

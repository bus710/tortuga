import 'dart:async'; // for streams
// import 'package:http/http.dart' as http;
import 'dart:convert'; // for JSON/struct conversion
import 'dart:html' as html;

/* Project packages */
import 'package:tortuga_frontend/AppEvent.dart';

class AppBLoC {
  /* Downstream: the front > the event > the native */
  final _backwardController = StreamController<AppEvent>();
  Sink<AppEvent> get backwardSink =>
      _backwardController.sink; // to send events from the page
  Stream<AppEvent> get backwardStream =>
      _backwardController.stream; // to receive events here

  /* Upstream: the native > the handler > the front */
  final _forwardController = StreamController<String>.broadcast();
  StreamSink<String> get forwardSink =>
      _forwardController.sink; // to send event here
  Stream<String> get forwardStream =>
      _forwardController.stream; // to receive events from the page

  /* To prevent the meesage being sent too fast */
  Timer timer;

  /* For the websocket comm */
  html.WebSocket socket;
  String socketState = "inactive";

  /* Buffer 
  1. to keep the data from stream 
  2. to use it in the timer handler */
  String buttonName = "none/none";

  // constructor
  AppBLoC() {
    // Connect the event and the handler.
    backwardStream.listen(_frontendHandler);

    // upStream.setMethodCallHandler(this._upStreamHandler);
    timer = Timer.periodic(Duration(milliseconds: 1000), callback);

    // Init the socket
    socketInit();
  }

  dispose() {
    // _appStreamController.close();
    _backwardController.close();
    _forwardController.close();
  }

  /* This handler connects these: 
  - The event of gesture control
  - The websocket.
  - So this is the place to make some action (i.e. marshaling). */
  _frontendHandler(AppEvent event) async {
    if (event.runtimeType.toString() == "ButtonEvent") {
      if (socket != null && socket.readyState == html.WebSocket.OPEN) {
        socket.send(json.encode({
          "ButtonName": event.buttonName,
        }));
      } else {
        // print('WebSocket not connected, message data not sent');
      }
      buttonName = event.buttonName;
    }
  }

  /* This timer handler prints and sends the last message */
  void callback(Timer timer) async {
    if (socket != null && socket.readyState == html.WebSocket.OPEN) {
      socket.send(json.encode({
        "ButtonName": buttonName,
      }));
    }

    print(buttonName);
  }

  socketInit() {
    socket = html.WebSocket(
        'ws://' + html.window.location.hostname + ':3000/message');

    socket.onOpen.listen((e) {
      print("websocket: opened");
      forwardSink.add("active");
      socketState = "active";
    });

    socket.onClose.listen((e) {
      print("websocket: closed");
      forwardSink.add("inactive");
      socketState = "inactive";
    });

    socket.onMessage.listen((e) {
      // Do nothing
    });
  }
}

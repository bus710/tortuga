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

  /* Buffer */
  int OriginalX, OriginalY;
  int DraggedX, DraggedY;

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
    if (event.runtimeType.toString() == "GestureEvent") {
      // If the gesture is done, send a 0,0,0,0 to stop the robot
      if (event.OriginalX == 0 &&
          event.OriginalY == 0 &&
          event.DraggedX == 0 &&
          event.DraggedY == 0) {
        if (socket != null && socket.readyState == html.WebSocket.OPEN) {
          socket.send(json.encode({
            "OriginalX": event.OriginalX,
            "OriginalY": event.OriginalY,
            "DraggedX": event.DraggedX,
            "DraggedY": event.DraggedY,
          }));
        } else {
          // print('WebSocket not connected, message data not sent');
        }
      }

      // Store the date from this event for later use
      OriginalX = event.OriginalX;
      OriginalY = event.OriginalY;
      DraggedX = event.DraggedX;
      DraggedY = event.DraggedY;
    }
  }

  void callback(Timer timer) async {
    // expired = true;
    if (socket != null && socket.readyState == html.WebSocket.OPEN) {
      socket.send(json.encode({
        "OriginalX": OriginalX,
        "OriginalY": OriginalY,
        "DraggedX": DraggedX,
        "DraggedY": DraggedY,
      }));
    }
  }

  socketInit() {
    socket = html.WebSocket(
        'ws://' + html.window.location.hostname + ':3000/message');

    socket.onOpen.listen((e) {
      print("websocket: opened");
    });

    socket.onClose.listen((e) {
      print("websocket: closed");
    });

    socket.onMessage.listen((e) {
      // Do nothing
    });
  }
}

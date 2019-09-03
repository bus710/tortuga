import 'dart:async'; // for streams
// import 'package:http/http.dart' as http;
import 'dart:convert'; // for JSON/struct conversion
import 'dart:html' as html;

/* Project packages */
import 'package:tortuga_frontend/AppEvent.dart';

class AppBLoC {
  /* Downstream: the front > the event > the native */
  final _backwardController = StreamController<AppEvent>();
  Sink<AppEvent> get backwardSink => _backwardController.sink; // to send events from the page
  Stream<AppEvent> get backwardStream => _backwardController.stream; // to receive events here

  /* Upstream: the native > the handler > the front */
  final _forwardController = StreamController<String>.broadcast();
  StreamSink<String> get forwardSink => _forwardController.sink; // to send event here
  Stream<String> get forwardStream => _forwardController.stream; // to receive events from the page

  /* Timer related */
  Timer timer;

  html.WebSocket socket;

  // constructor
  AppBLoC() {
    // Connect the event and the handler.
    backwardStream.listen(_frontendHandler);

    // upStream.setMethodCallHandler(this._upStreamHandler);
    timer = Timer.periodic(Duration(milliseconds: 100), callback);
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
      print(event.Speed);
      // Map map = new Map<String, dynamic>();
      // map[event.type] = event.data;
      // String m = await downStream.invokeMethod("setConfiguration", map);
      // print("${event.runtimeType.toString()} / " + m);
    }
  }

  void callback(Timer timer) async {}

  socketInit() {
    socket = html.WebSocket(
        'ws://' + html.window.location.hostname + ':3000/message');
    
    socket.onOpen.listen((e){
      print("websocket: opened");
    });

    socket.onClose.listen((e){
      print("websocket: closed");
    });

    socket.onMessage.listen((e){
      // Do nothing
    });
    
  }
}

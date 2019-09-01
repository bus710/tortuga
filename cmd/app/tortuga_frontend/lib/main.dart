import 'package:flutter_web/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Tortuga Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        fontFamily: 'OpenSans',
      ),
      home: MyHomePage(title: 'Tortuga'),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatelessWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
        centerTitle: true,
      ),
      body: AppPage(title: this.title),
    );
  }
}

class AppPage extends StatefulWidget {
  AppPage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() {
    return _AppState();
  }
}

class _AppState extends State<AppPage> {
  // max screen size to calcurate the gesture area
  double maxW;
  double maxH;

  // actual width of the gesture area
  double padW = 0;

  // actual size of the label
  double widX = 0;
  double widY = 0;

  // delta x, y to get the update of dragging
  double x = 0;
  double y = 0;

  // original x, y to store the spot when locked on
  double originalX = 0;
  double originalY = 0;

  // dragged x, y to show/send the value
  double draggedX = 0;
  double draggedY = 0;

  // message to show
  String message = "Released";

  @override
  Widget build(BuildContext context) {
    return Center(
      child: LayoutBuilder(builder: (context, constraint) {
        maxW = constraint.maxWidth;
        maxH = constraint.maxHeight;
        // print(maxW.toString() + " / " + maxH.toString());
        return _getColumn();
      }),
    );
  }

  Widget _getColumn() {
    if (maxW > 512) {
      padW = 512;
    } else {
      padW = maxW;
    }

    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: <Widget>[
        Container(
          width: maxW - 10,
          height: maxH / 3,
          child: Center(
            child: _get_dragged_string(),
          ),
        ),
        Container(
          width: padW - 30,
          height: padW - 30,
          decoration: BoxDecoration(
            color: Colors.white,
            border: Border.all(color: Colors.black, width: 10),
            borderRadius: BorderRadius.all(Radius.circular(20)),
          ),
          child: Stack(
            children: <Widget>[
              Positioned(
                right: widX,
                top: widY,
                child: Text(message, style: TextStyle(fontSize: 18)),
              ),
              GestureDetector(
                onPanStart: (d) => {
                  print("location: " + widX.toString() + "/" + widY.toString()),
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  widX = (x * -1) + padW - 10,
                  widY = y - 100,
                  print("start: " + x.toString() + "/" + y.toString()),
                  message = "Lock on",
                  originalX = x,
                  originalY = y,
                  draggedX = 0,
                  draggedY = 0,
                  setState(() {}),
                },
                onPanUpdate: (d) => {
                  print("location: " + widX.toString() + "/" + widY.toString()),
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  widX = (x * -1) + padW - 10,
                  widY = y - 100,
                  print("update: " + x.toString() + "/" + y.toString()),
                  draggedX = x - originalX,
                  draggedY = (y - originalY) * -1,
                  setState(() {}),
                },
                onPanEnd: (d) => {
                  message = "Released",
                  originalX = 0,
                  originalY = 0,
                  draggedX = 0,
                  draggedY = 0,
                  setState(() {}),
                },
                // child:
              ),
            ],
          ),
        ),
      ],
    );
  }

  Widget _get_dragged_string() {
    if (message == "Lock on") {
      return Text(
          "Dragged: " +
              draggedX.toInt().toString() +
              " / " +
              draggedY.toInt().toString(),
          style: TextStyle(fontSize: 24));
    } else {
      return Text("PRESS and DRAG", style: TextStyle(fontSize: 24));
    }
  }
}

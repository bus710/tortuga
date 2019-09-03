import 'package:flutter_web/gestures.dart';
import 'package:flutter_web/material.dart';
import 'package:tortuga_frontend/AppBLoC.dart';
import 'package:tortuga_frontend/AppEvent.dart';

class AppPage extends StatefulWidget {
  AppPage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() {
    return _AppState();
  }
}

class _AppState extends State<AppPage> {
  final _bloc = AppBLoC();

  // max screen size to calcurate the gesture area
  double maxW;
  double maxH;

  // actual width of the gesture area
  double padW = 0;

  // actual location of the label
  double labelX = 0;
  double labelY = 0;

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
              GestureDetector(
                dragStartBehavior: DragStartBehavior.start,
                // child: Container(
                //   width: 100,
                //   height: 100,
                // ),
                onPanStart: (d) => {
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  labelX = (x * -1) + padW - 50,
                  labelY = y - 50,
                  message = "Locked on",
                  originalX = x,
                  originalY = y,
                  draggedX = 0,
                  draggedY = 0,
                  // print("location: " +
                  //     labelX.toString() +
                  //     "/" +
                  //     labelY.toString()),
                  // print("start: " + x.toString() + "/" + y.toString()),
                  setState(() {}),
                },
                onPanUpdate: (d) => {
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  labelX = (x * -1) + padW - 50,
                  labelY = y - 50,
                  draggedX = x - originalX,
                  draggedY = (y - originalY) * -1,
                  // print("location: " +
                  //     labelX.toString() +
                  //     "/" +
                  //     labelY.toString()),
                  // print("update: " + x.toString() + "/" + y.toString()),
                  setState(() {}),
                },
                onPanEnd: (d) => {
                  message = "Released",
                  labelX = padW / 2,
                  labelY = padW / 2,
                  originalX = 0,
                  originalY = 0,
                  draggedX = 0,
                  draggedY = 0,
                  _bloc.backwardSink.add(GestureEvent(0, 0, 0, 0)),
                  setState(() {}),
                },
                // child:
              ),
              Positioned(
                right: labelX,
                top: labelY,
                child: Text(message, style: TextStyle(fontSize: 18)),
              ),
              _get_original_location(),
            ],
          ),
        ),
      ],
    );
  }

  Widget _get_original_location() {
    if (message == "Locked on") {
      return Positioned(
          right: (originalX * -1) + padW - 50,
          top: originalY,
          child: Icon(Icons.open_with));
    } else {
      return Positioned(right: -1000, top: -1000, child: Icon(Icons.open_with));
    }
  }

  Widget _get_dragged_string() {
    if (message == "Locked on") {
      _bloc.backwardSink.add(GestureEvent(originalX.toInt(), originalY.toInt(),
          draggedX.toInt(), draggedY.toInt()));

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
